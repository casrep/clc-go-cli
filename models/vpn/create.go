package vpn

import (
	"fmt"
	"net"
)

type CreateReq struct {
	Local  LocalPropertiesCreateReq
	Remote RemotePropertiesCreateReq
	Ipsec  IpSecCreateReq
	Ike    IkeCreateReq
}

type LocalPropertiesCreateReq struct {
	Alias   string   `json:"locationAlias" valid:"required"`
	Subnets []string `valid:"required"`
}

type RemotePropertiesCreateReq struct {
	SiteName   string   `valid:"required"`
	DeviceType string   `valid:"required"`
	Address    string   `valid:"required"`
	Subnets    []string `valid:"required"`
}

type IpSecCreateReq struct {
	Encryption string `oneOf:"aes128,aes192,aes256,tripleDES"`
	Hashing    string `oneOf:"sha1_96,sha1_256,md5"`
	Protocol   string `oneOf:"esp,ah"`
	Pfs        string `oneOf:"disabled,group1,group2,group5"`
	Lifetime   int64  `oneOf:"3600,28800,86400"`
}

type IkeCreateReq struct {
	Encryption         string `oneOf:"aes128,aes192,aes256,tripleDES"`
	Hashing            string `oneOf:"sha1_96,sha1_256,md5"`
	DiffieHellmanGroup string `oneOf:"group1,group2,group5"`
	PreSharedKey       string `valid:"required"`
	Lifetime           string `oneOf:"3600,28800,86400"`
	Mode               string `oneOf:"main,aggresive"`
	DeadPeerDetection  string `oneOf:"true,false,optional"`
	NatTraversal       string `oneOf:"true,false,optional"`
	RemoteIdentity     string
}

func (r *IkeCreateReq) Validate() error {
	if r.NatTraversal == "true" {
		ip := r.RemoteIdentity
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			return fmt.Errorf("remoteIdentity: %s is not valid IPv4", ip)
		}
	} else {
		r.RemoteIdentity = ""
	}

	return nil
}
