package datacenter

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type ListReq struct{}

type ListRes struct {
	Id    string
	Name  string
	Links []models.LinkEntity
}