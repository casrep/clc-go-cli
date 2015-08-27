package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/help"
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
		Help: help.Command{
			Brief: []string{
				"Creates a new server.",
				"Use this API operation when you want to create a new server from a standard or custom template, or clone an existing server.",
			},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server"},
				},
				{
					"--group-id",
					[]string{"ID of the parent group."},
				},
				{
					"--source-server-id",
					[]string{"ID of the server to use a source."},
				},
				{
					"--is-managed-os",
					[]string{
						"Whether to create the server as managed or not. Default is false.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--is-managed-backup",
					[]string{
						"Whether to add managed backup to the server. Must be a managed OS server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--ip-address",
					[]string{
						"IP address to assign to the server. If not provided, one will be assigned automatically.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--root-password",
					[]string{"Password of administrator or root user on server. If not provided, one will be generated automatically."},
				},
				{
					"--source-server-password",
					[]string{
						"Password of the source server, used only when creating a clone from an existing server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--cpu",
					[]string{"Number of processors to configure the server with (1-16). Ignored for bare metal servers."},
				},
				{
					"--cpu-autoscale-policy-id",
					[]string{
						"ID of the vertical CPU Autoscale policy to associate the server with.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--memory-gb",
					[]string{
						"Number of GB of memory to configure the server with (1-128).",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--type",
					[]string{"Whether to create a standard, hyperscale, or bareMetal server."},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage.",
						"If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--anti-affinity-policy-id",
					[]string{
						"For standard servers, whether to use standard or premium storage. If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale. Ignored for bare metal servers.",
					},
				},
				{
					"--custom-fields",
					[]string{"Collection of custom field ID-value pairs to set for the server."},
				},
				{
					"--additional-disks",
					[]string{"Collection of disk parameters. Ignored for bare metal servers."},
				},
				{
					"--ttl",
					[]string{"Date/time that the server should be deleted. Ignored for bare metal servers."},
				},
				{
					"--packages",
					[]string{"Collection of packages to run on the server after it has been built. Ignored for bare metal servers."},
				},
				{
					"--configuration-id",
					[]string{
						"Specifies the identifier for the specific configuration type of bare metal server to deploy.",
						"Ignored for standard and hyperscale servers.",
					},
				},
				{
					"--os-type",
					[]string{
						"Specifies the OS to provision with the bare metal server. Currently, the only supported OS types",
						"are redHat6_64Bit, centOS6_64Bit, windows2012R2Standard_64Bit.",
						"Ignored for standard and hyperscale servers.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.DeleteReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Sends the delete operation to a given server and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server to be deleted."},
				},
			},
		},
	})
	registerCommandBase(&server.UpdateReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Changes the amount of CPU cores, memory (in GB), server credentials, custom fields, description, disks and server's group."},
			Arguments: []help.Argument{
				{
					"--cpu",
					[]string{"The amount of CPU cores to set for the given server."},
				},
				{
					"--memory",
					[]string{"The amount of memory (in GB) to set for the given server."},
				},
				{
					"--root-password",
					[]string{
						"The current and new administrator/root password values.",
						"Has to be an object with 2 fields:",
						"1) current: the current administrator/root password used to login;",
						"2) password: the new administrator/root password to change to.",
					},
				},
				{
					"--custom-fields",
					[]string{
						"A list of id-value pairs for all custom fields including all required values",
						"and other custom field values that you wish to set.",
						"",
						"Note: You must specify the complete list of custom field values to set on the server.",
						"If you want to change only one value, specify all existing field values",
						"along with the new value for the field you wish to change.",
						"To unset the value for an unrequired field, you may leave the field id-value pairing out,",
						"however all required fields must be included.",
					},
				},
				{
					"--description",
					[]string{"The description of the server to set"},
				},
				{
					"--group-id",
					[]string{"The unique identifier of the group to set as the parent."},
				},
				{
					"--disks",
					[]string{
						"A list of information for all disks to be on the server including type (raw or partition), size, and path",
						"",
						"Note: You must specify the complete list of disks to be on the server.",
						"If you want to add or resize a disk, specify all existing disks/sizes",
						"along with a new entry for the disk to add or the new size of an existing disk.",
						"To delete a disk, just specify all the disks that should remain.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.GetReq{}, &server.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details for a individual server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server being queried."},
				},
			},
		},
	})
	registerCommandBase(&server.GetCredentialsReq{}, &server.GetCredentialsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/credentials",
		Resource: "server",
		Command:  "get-credentials",
		Help: help.Command{
			Brief: []string{"Retrieves the administrator/root password on an existing server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server with the credentials to return."},
				},
			},
		},
	})
	registerCommandBase(&server.GetImportsReq{}, &server.GetImportsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}/{LocationId}/available",
		Resource: "server",
		Command:  "get-imports",
		Help: help.Command{
			Brief: []string{"Gets the list of available servers that can be imported."},
			Arguments: []help.Argument{
				{
					"--location-id",
					[]string{"Data center location identifier."},
				},
			},
		},
	})
	registerCommandBase(&server.GetIPAddressReq{}, &server.GetIPAddressRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "get-public-ip-address",
		Help: help.Command{
			Brief: []string{"Gets the details for the public IP address of a server, including the specific set of protocols and ports allowed and any source IP restrictions."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"The specific public IP to return details about."},
				},
			},
		},
	})
	registerCommandBase(&server.AddIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses",
		Resource: "server",
		Command:  "add-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Claims a public IP address and associates it with a server, allowing access to it on a given set of protocols and ports.",
				"It may also be set to restrict access based on a source IP range.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server being queried."},
				},
				{
					"--internal-ip-address",
					[]string{
						"The internal (private) IP address to map to the new public IP address.",
						"If not provided, one will be assigned for you.",
					},
				},
				{
					"--ports",
					[]string{
						"The set of ports and protocols to allow access to for the new public IP address.",
						"Only these specified ports on the respective protocols will be accessible",
						"when accessing the server using the public IP address claimed here.",
						"Has to be a list of objects with fields port, portTo and protocol.",
					},
				},
				{
					"--source-restrictions",
					[]string{
						"A list of the source IP address ranges allowed to access the new public IP address.",
						"Used to restrict access to only the specified ranges of source IPs.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.RemoveIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "remove-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Releases the given public IP address of a server so that it is no longer associated with the server",
				"and available to be claimed again by another server.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"The specific public IP to remove."},
				},
			},
		},
	})
	registerCommandBase(&server.UpdateIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "update-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Updates a public IP address on an existing server, allowing access to it on a given set of protocols and ports",
				"as well as restricting access based on a source IP range.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"The specific public IP to update."},
				},
				{
					"--ports",
					[]string{
						"The set of ports and protocols to allow access to for the public IP address.",
						"Only these specified ports on the respective protocols will be accessible",
						"when accessing the server using the public IP address claimed here.",
						"Has to be a list of objects with fields port, portTo and protocol.",
					},
				},
				{
					"--source-restrictions",
					[]string{
						"A list of the source IP address ranges allowed to access the public IP address.",
						"Used to restrict access to only the specified ranges of source IPs.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOn",
		Resource: "server",
		Command:  "power-on",
		Help: help.Command{
			Brief: []string{"Sends the power on operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform power on operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOff",
		Resource: "server",
		Command:  "power-off",
		Help: help.Command{
			Brief: []string{"Sends the power off operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform power off operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/pause",
		Resource: "server",
		Command:  "pause",
		Help: help.Command{
			Brief: []string{"Sends the pause operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform pause operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/reset",
		Resource: "server",
		Command:  "reset",
		Help: help.Command{
			Brief: []string{"Sends the reset operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform reset operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/shutDown",
		Resource: "server",
		Command:  "shut-down",
		Help: help.Command{
			Brief: []string{"Sends the shut-down operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform shut-down operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/archive",
		Resource: "server",
		Command:  "archive",
		Help: help.Command{
			Brief: []string{"Sends the archive operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to perform archive operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.RestoreReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/restore",
		Resource: "server",
		Command:  "restore",
		Help: help.Command{
			Brief: []string{"Restores a given archived server to a specified group."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the archived server to restore."},
				},
				{
					"--target-group-id",
					[]string{"The unique identifier of the target group to restore the server to."},
				},
			},
		},
	})
	registerCommandBase(&server.CreateSnapshotReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/createSnapshot",
		Resource: "server",
		Command:  "create-snapshot",
		Help: help.Command{
			Brief: []string{"Sends the create snapshot operation to a list of servers (along with the number of days to keep the snapshot for) and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server names to perform create snapshot operation on."},
				},
				{
					"--snapshot-expiration-days",
					[]string{"Number of days to keep the snapshot for (must be between 1 and 10)."},
				},
			},
		},
	})
	registerCommandBase(&server.RevertToSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}/restore",
		Resource: "server",
		Command:  "revert-to-snapshot",
		Help: help.Command{
			Brief: []string{"Reverts a server to a snapshot."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server with the snapshot to restore."},
				},
				{
					"--snapshot-id",
					[]string{"ID of the snapshot to restore."},
				},
			},
		},
	})
	registerCommandBase(&server.DeleteSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}",
		Resource: "server",
		Command:  "delete-snapshot",
		Help: help.Command{
			Brief: []string{"Deletes a given server snapshot."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server with the snapshot to delete."},
				},
				{
					"--snapshot-id",
					[]string{"ID of the snapshot to delete."},
				},
			},
		},
	})
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/startMaintenance",
		Resource: "server",
		Command:  "start-maintenance-mode",
		Help: help.Command{
			Brief: []string{"Sends a start maintenance mode operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to start maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/stopMaintenance",
		Resource: "server",
		Command:  "stop-maintenance-mode",
		Help: help.Command{
			Brief: []string{"Sends a stop maintenance mode operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to stop maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.Import{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}",
		Resource: "server",
		Command:  "import",
		Help: help.Command{
			Brief: []string{"Imports a new server from an uploaded OVF."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server."},
				},
				{
					"--group-id",
					[]string{"ID of the parent group."},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--root-password",
					[]string{
						"Password of administrator or root user on server. This password must match",
						"the one set on the server being imported or the import will fail.",
					},
				},
				{
					"--cpu",
					[]string{
						"Number of processors to configure the server with (1-16). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--memoryGB",
					[]string{
						"Number of GB of memory to configure the server with (1-128). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--type",
					[]string{"Whether to create standard or hyperscale server"},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage. If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
					},
				},
				{
					"--custom-fields",
					[]string{"Collection of custom field ID-value pairs to set for the server."},
				},
				{
					"--ovf-id",
					[]string{"The identifier of the OVF that defines the server to import."},
				},
				{
					"--ovf-os-type",
					[]string{
						"The OS type of the server being imported. Currently, the only supported OS types",
						"are redHat6_64Bit, windows2008R2DataCenter_64bit, and windows2012R2DataCenter_64Bit.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.AddNetwork{}, &models.Status{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/networks",
		Resource: "server",
		Command:  "add-secondary-network",
		Help: help.Command{
			Brief: []string{"Adds a secondary network adapter to a given server in a given account."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server."},
				},
				{
					"--network-id",
					[]string{"ID of the network."},
				},
				{
					"--ip-address",
					[]string{"Optional IP address for the network ID."},
				},
			},
		},
	})
	registerCommandBase(&server.RemoveNetwork{}, &models.Status{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/networks/{NetworkId}",
		Resource: "server",
		Command:  "remove-secondary-network",
		Help: help.Command{
			Brief: []string{"Removes a secondary network adapter from a given server in a given account."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"ID of the server."},
				},
				{
					"--network-id",
					[]string{"ID of the network."},
				},
			},
		},
	})

	registerCommandBase(&group.GetReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details for a individual server group and any sub-groups and servers that it contains."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.CreateReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}",
		Resource: "group",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a new group."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{"Name of the group to create."},
				},
				{
					"--description",
					[]string{"User-defined description of this group."},
				},
				{
					"--parent-group-id",
					[]string{"ID of the parent group."},
				},
				{
					"--custom-fields",
					[]string{"Collection of custom field ID-value pairs to set for the server."},
				},
			},
		},
	})
	registerCommandBase(&group.DeleteReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Sends the delete operation to a given group and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group to be deleted."},
				},
			},
		},
	})
	registerCommandBase(&group.GetBillingReq{}, &group.GetBillingRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/billing",
		Resource: "group",
		Command:  "get-billing-details",
		Help: help.Command{
			Brief: []string{"Gets the current and estimated charges for each server in a designated group hierarchy."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.GetStatsReq{}, &[]group.GetStatsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/statistics?start={Start}&end={End}&sampleInterval={SampleInterval}&type={Type}",
		Resource: "group",
		Command:  "get-monitoring-statistics",
		Help: help.Command{
			Brief: []string{
				"Gets the resource consumption details for whatever window specified in the request.",
				"Data can be retrieved for a variety of time windows and intervals.",
			},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group being queried."},
				},
				{
					"--type",
					[]string{
						"Valid values are latest, hourly, or realtime.",
						"",
						"'latest' will return a single data point that reflects the last monitoring data collected.",
						"No start, end, or sampleInterval values are required for this type.",
						"",
						"'hourly' returns data points for each sampleInterval value between the start and end times provided.",
						"The start and sampleInterval parameters are both required for this type.",
						"",
						"'realtime' will return data from the last 4 hours, available in smaller increments.",
						"To use realtime type, start parameter must be within the last 4 hours.",
						"The start and sampleInterval parameters are both required for this type.",
					},
				},
				{
					"--start",
					[]string{
						"DateTime (UTC) of the query window. Note that statistics are only held for 14 days.",
						"Start date (and optional end date) must be within the past 14 days.",
						"Value is not required if choosing the latest query type.",
					},
				},
				{
					"--end",
					[]string{
						"DateTime (UTC) of the query window. Default is the current time in UTC.",
						"End date (and start date) must be within the past 14 days.",
						"Not a required value if results should be up to the current time.",
					},
				},
				{
					"--sample-interval",
					[]string{
						"Result interval. For the default hourly type, the minimum value is 1 hour (01:00:00)",
						"and maximum is the full window size of 14 days. Note that interval must fit within start/end window,",
						"or you will get an exception that states: 'The 'end' parameter must represent a time that occurs at least one 'sampleInterval' before 'start.'",
						"If realtime type is specified, interval can be as small as 5 minutes (05:00).",
					},
				},
			},
		},
	})
	registerCommandBase(&group.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Changes the custom fields, name, description and parent group of the given group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group being updated."},
				},
				{
					"--custom-fields",
					[]string{
						"A list of id-value pairs for all custom fields including all required values and other custom field",
						"values that you wish to set.",
						"",
						"Note: You must specify the complete list of custom field values",
						"to set on the group. If you want to change only one value,",
						"specify all existing field values along with the new value for the field you wish to change.",
						"To unset the value for an unrequired field, you may leave the field id-value pairing out,",
						"however all required fields must be included",
					},
				},
				{
					"--name",
					[]string{"The name to set for the group."},
				},
				{
					"--description",
					[]string{"The description to set for the group."},
				},
				{
					"--parent-group-id",
					[]string{"The group identifier for the new parent group."},
				},
			},
		},
	})
	registerCommandBase(&group.GetReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/archive",
		Resource: "group",
		Command:  "archive",
		Help: help.Command{
			Brief: []string{"Sends the archive operation to a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group to archive."},
				},
			},
		},
	})
	registerCommandBase(&group.RestoreReq{}, &group.RestoreRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/restore",
		Resource: "group",
		Command:  "restore",
		Help: help.Command{
			Brief: []string{"Sends the restore operation to an archived group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"ID of the group to restore."},
				},
				{
					"--target-group-id",
					[]string{"The unique identifier of the target group to restore the group to."},
				},
			},
		},
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
	registerCommandBase(&balancer.Get{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "get",
	})
	registerCommandBase(&balancer.UpdatePool{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "update",
	})
	registerCommandBase(&balancer.Update{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "update",
	})
	registerCommandBase(&balancer.GetNodes{}, &[]balancer.Node{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "get-nodes",
	})
	registerCommandBase(&balancer.UpdateNodes{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "update-nodes",
	})
	registerCommandBase(&balancer.DeletePool{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "delete",
	})
	registerCommandBase(&balancer.Delete{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "delete",
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
