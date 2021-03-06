package model_loader

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/parser"
)

func LoadModel(parsedArgs map[string]interface{}, inputModel interface{}) error {
	metaModel := reflect.ValueOf(inputModel)
	if !metaModel.IsValid() {
		return nil
	}
	if metaModel.Kind() != reflect.Ptr {
		return fmt.Errorf("Input model must be passed by pointer.")
	}
	for k, v := range parsedArgs {
		field, err := getFieldByName(metaModel, k)
		if err != nil {
			return err
		}
		err = loadValue(k, v, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadValue(key string, arg interface{}, field reflect.Value) error {
	switch field.Interface().(type) {
	case int64, *int64:
		var argInt int64
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Int {
			argInt = arg.(int64)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.Float64 {
			if valid.IsWhole(arg.(float64)) {
				argInt = int64(arg.(float64))
				mismatch = false
			}
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if valid.IsInt(arg.(string)) {
				var err error
				argInt, err = valid.ToInt(arg.(string))
				if err != nil {
					if num, ok := err.(*strconv.NumError); ok {
						if num.Err == strconv.ErrRange {
							return fmt.Errorf("Value `%s` is too big.", arg.(string))
						}
					}
					return err
				}
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be integer.", key)
		} else {
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&argInt))
			} else {
				field.SetInt(argInt)
			}
			return nil
		}
	case float64:
		var argFloat64 float64
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Float64 {
			argFloat64 = arg.(float64)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if valid.IsFloat(arg.(string)) {
				var err error
				argFloat64, err = valid.ToFloat(arg.(string))
				if num, ok := err.(*strconv.NumError); ok {
					if num.Err == strconv.ErrRange {
						return fmt.Errorf("Value `%s` is too big.", arg.(string))
					}
				}
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be float.", key)
		} else {
			field.SetFloat(argFloat64)
			return nil
		}
	case time.Time:
		var argTime time.Time
		var err error
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.String {
			if argTime, err = time.Parse(base.TIME_FORMAT, arg.(string)); err == nil {
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be datetime in `%s` format.", key, base.TIME_FORMAT_REPR)
		} else {
			field.Set(reflect.ValueOf(argTime))
			return nil
		}
	case bool:
		var argBool bool
		var mismatch = true
		if reflect.ValueOf(arg).Kind() == reflect.Bool {
			argBool = arg.(bool)
			mismatch = false
		} else if reflect.ValueOf(arg).Kind() == reflect.String {
			if arg == "true" {
				argBool = true
				mismatch = false
			} else if arg == "false" {
				argBool = false
				mismatch = false
			}
		}
		if mismatch {
			return fmt.Errorf("Type mismatch: %s value must be either true or false.", key)
		} else {
			field.SetBool(argBool)
		}
		return nil
	case string:
		if reflect.ValueOf(arg).Kind() != reflect.String {
			return fmt.Errorf("Type mismatch: %s value must be string.", key)
		}
		field.SetString(arg.(string))
		return nil
	case base.NilField:
		if arg == nil {
			field.FieldByName("Set").SetBool(true)
			return nil
		}
		return fmt.Errorf("%s does not accept any value.", key)
	case map[string]string:
		argMapStringString, err := parseMapStringString(arg)
		if err != nil {
			if _, ok := err.(ParseObjWrongTypeError); ok {
				return fmt.Errorf("Type mismatch: %s must be an object", key)
			}
			return err
		}
		field.Set(reflect.ValueOf(argMapStringString))
		return nil
	}
	if isStruct(field) {
		argStruct, err := parseStruct(arg)
		if err != nil {
			if _, ok := err.(ParseObjWrongTypeError); ok {
				return fmt.Errorf("Type mismatch: %s must be an object.", key)
			}
			return err
		}
		for k, v := range argStruct {
			nestedField, err := getFieldByName(field.Addr(), k)
			if err != nil {
				return err
			}
			err = loadValue(k, v, nestedField)
			if err != nil {
				return err
			}
		}
		return nil
	} else if isSlice(field) {
		argSlice, err := parseSlice(arg)
		if err != nil {
			if _, ok := err.(ParseObjWrongTypeError); ok {
				return fmt.Errorf("Type mismatch: %s must be an array.", key)
			}
			return err
		}
		for _, v := range argSlice {
			elementPtr := getEmplySliceType(field)
			err = loadValue(key, v, elementPtr.Elem())
			if err != nil {
				return err
			}
			field.Set(reflect.Append(field, elementPtr.Elem()))
		}
		return nil
	}
	return fmt.Errorf("Unsupported field type %s", field.Kind())
}

func getFieldByName(model reflect.Value, name string) (reflect.Value, error) {
	field := model.Elem().FieldByName(name)
	fieldMeta, exist := model.Elem().Type().FieldByName(name)
	if !exist || fieldMeta.Tag.Get("argument") == "composed" {
		return reflect.ValueOf(nil), fmt.Errorf("Unknown option or argument: `%s`.", name)
	}
	return field, nil
}

func isStruct(model reflect.Value) bool {
	return model.Kind() == reflect.Struct
}

func isSlice(model reflect.Value) bool {
	return model.Kind() == reflect.Slice
}

// Parses an object of type map[string]interface{} either from JSON or from a=b,c=d,.. notation.
// Also, calls NormalizeKeys with the parsed object.
// If arg is already of type map[string]interface{} returns it as is.
func parseStruct(arg interface{}) (map[string]interface{}, error) {
	if argMap, isMap := arg.(map[string]interface{}); isMap {
		parser.NormalizeKeys(argMap)
		return argMap, nil
	}

	argString, isString := arg.(string)
	if !isString {
		return nil, ParseObjWrongTypeError{}
	}

	parsed := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(argString), &parsed); err == nil {
		parser.NormalizeKeys(parsed)
		return parsed, nil
	}
	if parsed, err := parser.ParseObject(argString, true); err == nil {
		return parsed, nil
	}
	return nil, fmt.Errorf("`%s` must be an object specified either in JSON or in key=value,.. format", argString)
}

// Parses an object of type []interface{} either from JSON.
// If arg is already of type []interface{} returns it as is.
func parseSlice(arg interface{}) ([]interface{}, error) {
	if argSlice, isSlice := arg.([]interface{}); isSlice {
		return argSlice, nil
	}

	argString, isString := arg.(string)
	if !isString {
		return nil, ParseObjWrongTypeError{}
	}

	parsed := make([]interface{}, 0)
	if err := json.Unmarshal([]byte(argString), &parsed); err == nil {
		return parsed, nil
	}
	return []interface{}{arg}, nil
}

// Parses an object of type map[string]string either from JSON or from a=b,c=d,.. notation.
func parseMapStringString(arg interface{}) (map[string]string, error) {
	var argMapStringString = make(map[string]string, 0)

	argMap, isMap := arg.(map[string]interface{})
	if !isMap {
		argString, isString := arg.(string)
		if !isString {
			return nil, ParseObjWrongTypeError{}
		}
		argMap = make(map[string]interface{}, 0)
		err := json.Unmarshal([]byte(argString), &argMap)
		if err != nil {
			argMap, err = parser.ParseObject(argString, false)
		}
		if err != nil {
			return nil, ParseObjWrongTypeError{}
		}
	}
	for k, v := range argMap {
		if _, isString := v.(string); !isString {
			return nil, fmt.Errorf("Type mismatch: `%s` must be string", k)
		}
		argMapStringString[k] = v.(string)
	}
	return argMapStringString, nil
}

func getEmplySliceType(slice reflect.Value) reflect.Value {
	return reflect.New(slice.Type().Elem())
}
