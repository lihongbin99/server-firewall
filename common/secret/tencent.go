package secret

import (
	"fmt"
	"gopkg.in/ini.v1"
	"security-network/server/tencent"
)

type TencentLighthouse struct {
	Name       string
	Endpoint   string
	Region     string
	InstanceId string
	SecretId   string
	SecretKey  string
}

func (that *TencentLighthouse) ParseIni(section *ini.Section) {
	that.Name = section.Name()
	if that.Endpoint = section.Key("endpoint").String(); that.Endpoint == "" {
		panic(fmt.Errorf("place input Endpoint"))
	}
	if that.Region = section.Key("region").String(); that.Region == "" {
		panic(fmt.Errorf("place input Region"))
	}
	if that.InstanceId = section.Key("instance").String(); that.InstanceId == "" {
		panic(fmt.Errorf("place input InstanceId"))
	}
	if that.SecretId = section.Key("secretId").String(); that.SecretId == "" {
		panic(fmt.Errorf("place input SecretId"))
	}
	if that.SecretKey = section.Key("secretKey").String(); that.SecretKey == "" {
		panic(fmt.Errorf("place input SecretKey"))
	}
}

func (that *TencentLighthouse) GetName() string {
	return that.Name
}
func (that *TencentLighthouse) Create(ip, name string) error {
	return tencent.CreateLighthouse(that.Endpoint, that.Region, that.InstanceId, that.SecretId, that.SecretKey, ip, name)
}
func (that *TencentLighthouse) Delete(ip string) error {
	return tencent.DeleteLighthouse(that.Endpoint, that.Region, that.InstanceId, that.SecretId, that.SecretKey, ip)
}
