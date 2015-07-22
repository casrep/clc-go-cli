package connection_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
)

var serveMux *http.ServeMux
var server *httptest.Server

func initTest() {
	serveMux = http.NewServeMux()
	server = httptest.NewServer(serveMux)
	connection.BaseUrl = server.URL + "/"
}

func finishTest() {
	server.Close()
}

func addHandler(t *testing.T, url string, reqModel, resModel interface{}) {
	serveMux.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		if reqModel != nil {
			reqContent, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Error(err)
			}
			reqModel1 := reflect.New(reflect.ValueOf(reqModel).Elem().Type()).Interface()
			err = json.Unmarshal(reqContent, reqModel1)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(reqModel, reqModel1) {
				t.Errorf("Expected: %#v, obtained: %#v", reqModel, reqModel1)
			}
		}
		js, err := json.Marshal(resModel)
		if err != nil {
			t.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func newConnection(t *testing.T, registerHandler bool) (base.Connection, error) {
	if registerHandler {
		resModel := &authentication.LoginRes{AccountAlias: "ALIAS", BearerToken: "token"}
		reqModel := &authentication.LoginReq{Username: "user", Password: "password"}
		addHandler(t, "/authentication/login", reqModel, resModel)
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return connection.NewConnection("user", "password", logger)
}

type testReqModel struct {
	P1, P2 string
}

type testResModel struct {
	P1, P2 string
}

func TestNewConnection(t *testing.T) {
	initTest()
	defer finishTest()
	cn, err := newConnection(t, true)
	if err != nil {
		t.Error(err)
	}
	//test that bearer token and account alias are attached to subsequent requests
	serveMux.HandleFunc("/some-url/ALIAS", func(w http.ResponseWriter, req *http.Request) {
		if h, ok := req.Header["Authorization"]; !ok || len(h) == 0 || h[0] != "Bearer token" {
			t.Errorf("Incorrect request: bearer token not found, headers: %#v", req.Header)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(""))
	})
	err = cn.ExecuteRequest("GET", connection.BaseUrl+"some-url/{accountAlias}", nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewConnectionError(t *testing.T) {
	initTest()
	defer finishTest()
	_, err := newConnection(t, false)
	if err == nil || err.Error() != "Error occured while sending request to API. Status code: 404." {
		t.Errorf("Unexpected error: %s", err)
	}
}
