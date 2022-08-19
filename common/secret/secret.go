package secret

import (
	"fmt"
	"gopkg.in/ini.v1"
)

const (
	TencentLighthouseType = "TencentLighthouse"
	AliyunEcsType         = "AliyunEcs"
)

type Secret interface {
	ParseIni(*ini.Section)
	GetName() string
	Create(ip, name string) error
	Delete(ip string) error
}

func GetSecret(section *ini.Section) Secret {
	var result Secret = nil

	switch section.Key("type").String() {
	case AliyunEcsType:
		result = &AliyunEcs{}
	case TencentLighthouseType:
		result = &TencentLighthouse{}
	default:
		panic(fmt.Errorf("secret type error"))
	}

	result.ParseIni(section)
	return result
}
