package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/affinity"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"github.com/centurylinkcloud/clc-go-cli/models/balancer"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/firewall"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/models/network"
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
		Command:  "get-public-ip-address",
	})
	registerCommandBase(&server.AddIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses",
		Resource: "server",
		Command:  "add-public-ip-address",
	})
	registerCommandBase(&server.RemoveIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "remove-public-ip-address",
	})
	registerCommandBase(&server.UpdateIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "update-public-ip-address",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOn",
		Resource: "server",
		Command:  "power-on",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOff",
		Resource: "server",
		Command:  "power-off",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/pause",
		Resource: "server",
		Command:  "pause",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/reset",
		Resource: "server",
		Command:  "reset",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/shutDown",
		Resource: "server",
		Command:  "shut-down",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/archive",
		Resource: "server",
		Command:  "archive",
	})
	registerCommandBase(&server.RestoreReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/restore",
		Resource: "server",
		Command:  "restore",
	})
	registerCommandBase(&server.CreateSnapshotReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/createSnapshot",
		Resource: "server",
		Command:  "create-snapshot",
	})
	registerCommandBase(&server.RevertToSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}/restore",
		Resource: "server",
		Command:  "revert-to-snapshot",
	})
	registerCommandBase(&server.DeleteSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}",
		Resource: "server",
		Command:  "delete-snapshot",
	})

	registerCommandBase(&group.GetReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "get",
	})
	registerCommandBase(&group.CreateReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}",
		Resource: "group",
		Command:  "create",
	})
	registerCommandBase(&group.DeleteReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "delete",
	})
	registerCommandBase(&group.GetBillingReq{}, &group.GetBillingRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/billing",
		Resource: "group",
		Command:  "get-billing-details",
	})
	registerCommandBase(&group.GetStatsReq{}, &[]group.GetStatsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/statistics?start={Start}&end={End}&sampleInterval={SampleInterval}&type={Type}",
		Resource: "group",
		Command:  "get-monitoring-statistics",
	})
	registerCommandBase(&group.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "update",
	})
	registerCommandBase(&group.GetReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/archive",
		Resource: "group",
		Command:  "archive",
	})
	registerCommandBase(&group.RestoreReq{}, &group.RestoreRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/restore",
		Resource: "group",
		Command:  "restore",
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
	registerCommandBase(&datacenter.GetDCReq{}, &datacenter.GetDCRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}/{DataCenter}/deploymentCapabilities",
		Resource: "data-center",
		Command:  "get-deployment-capabilities",
	})

	registerCommandBase(&network.ListReq{}, &[]network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}",
		Resource: "network",
		Command:  "list",
	})
	registerCommandBase(&network.GetReq{}, &network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}?ipAddresses={IpAddresses}",
		Resource: "network",
		Command:  "get",
	})
	registerCommandBase(&network.ListIpAddresses{}, &[]network.IpAddress{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}/ipAddresses?type={Type}",
		Resource: "network",
		Command:  "list-ip-addresses",
	})
	registerCommandBase(&network.CreateReq{}, &network.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/claim",
		Resource: "network",
		Command:  "create",
	})
	registerCommandBase(&network.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}",
		Resource: "network",
		Command:  "update",
	})
	registerCommandBase(&network.ReleaseReq{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}/release",
		Resource: "network",
		Command:  "release",
	})

	registerCommandBase(&alert.CreateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "create",
	})
	registerCommandBase(nil, &alert.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "list",
	})
	registerCommandBase(&alert.GetReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "get",
	})
	registerCommandBase(&alert.UpdateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "update",
	})
	registerCommandBase(&alert.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "delete",
	})

	registerCommandBase(&affinity.CreateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "create",
	})
	registerCommandBase(nil, &affinity.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "list",
	})
	registerCommandBase(&affinity.GetReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "get",
	})
	registerCommandBase(&affinity.UpdateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "update",
	})
	registerCommandBase(&affinity.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "delete",
	})

	registerCommandBase(&firewall.CreateReq{}, &firewall.CreateRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}",
		Resource: "firewall-policy",
		Command:  "create",
	})
	registerCommandBase(&firewall.ListReq{}, &[]firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}?destinationAccount={DestinationAccountAlias}",
		Resource: "firewall-policy",
		Command:  "list",
	})
	registerCommandBase(&firewall.GetReq{}, &firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "get",
	})
	registerCommandBase(&firewall.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "update",
	})
	registerCommandBase(&firewall.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "delete",
	})

	registerCommandBase(&balancer.CreatePool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "create",
	})
	registerCommandBase(&balancer.Create{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "create",
	})
	registerCommandBase(&balancer.ListPools{}, &[]balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "list",
	})
	registerCommandBase(&balancer.List{}, &[]balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "list",
	})
	registerCommandBase(&balancer.GetPool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "get",
	})

	registerCommandBase(nil, &[]customfields.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/accounts/{accountAlias}/customFields",
		Resource: "custom-fields",
		Command:  "get",
	})

	registerCustomCommand(commands.NewGroupList(commands.CommandExcInfo{
		Resource: "group",
		Command:  "list",
	}))
	registerCustomCommand(commands.NewServerList(commands.CommandExcInfo{
		Resource: "server",
		Command:  "list",
	}))
	registerCustomCommand(commands.NewWait(commands.CommandExcInfo{
		Resource: "wait",
	}))
	registerCustomCommand(commands.NewLogin(commands.CommandExcInfo{
		Resource: "login",
	}))
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info commands.CommandExcInfo) {
	cmd := &commands.CommandBase{
		Input:   inputModel,
		Output:  outputModel,
		ExcInfo: info,
	}
	AllCommands = append(AllCommands, cmd)
}

func registerCustomCommand(cmd base.Command) {
	AllCommands = append(AllCommands, cmd)
}
