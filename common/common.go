package common

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/memory"
)

var reg registry.Registry

func GetReg() registry.Registry {
	if reg == nil {
		reg = memory.NewRegistry()
		return reg
	}
	return reg
}
