package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

var AllCommands []base.Command = make([]base.Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
	})
	registerCommandBase(&server.DeleteReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "delete",
	})
	registerCommandBase(&server.UpdateReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "update",
	})
	registerCommandBase(&server.GetReq{}, &server.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "get",
	})
	registerCommandBase(&server.GetCredentialsReq{}, &server.GetCredentialsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/credentials",
		Resource: "server",
		Command:  "get-credentials",
	})
	registerCommandBase(&server.GetImportsReq{}, &server.GetImportsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}/{LocationId}/available",
		Resource: "server",
		Command:  "get-imports",
	})
	registerCommandBase(&server.GetIPAddressReq{}, &server.GetIPAddressRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "get-ip-address",
	})
	registerCommandBase(&server.AddIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses",
		Resource: "server",
		Command:  "add-ip-address",
	})

	registerCommandBase(&datacenter.ListReq{}, &[]datacenter.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}",
		Resource: "data-center",
		Command:  "list",
	})
	registerCommandBase(&datacenter.GetReq{}, &datacenter.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}/{DataCenter}?groupLinks={GroupLinks}",
		Resource: "data-center",
		Command:  "get",
	})
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info commands.CommandExcInfo) {
	cmd := &commands.CommandBase{
		Input:   inputModel,
		Output:  outputModel,
		ExcInfo: info,
	}
	AllCommands = append(AllCommands, cmd)
}
