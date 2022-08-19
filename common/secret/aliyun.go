package secret

import (
	"fmt"
	"gopkg.in/ini.v1"
	"security-network/server/aliyun"
)

type AliyunEcs struct {
	Name            string
	Endpoint        string
	RegionId        string
	SecurityGroupId string
	AccessKeyId     string
	AccessKeySecret string
}

func (that *AliyunEcs) ParseIni(section *ini.Section) {
	that.Name = section.Name()
	if that.Endpoint = section.Key("endpoint").String(); that.Endpoint == "" {
		panic(fmt.Errorf("place input Endpoint"))
	}
	if that.RegionId = section.Key("regionId").String(); that.RegionId == "" {
		panic(fmt.Errorf("place input RegionId"))
	}
	if that.SecurityGroupId = section.Key("securityGroupId").String(); that.SecurityGroupId == "" {
		panic(fmt.Errorf("place input SecurityGroupId"))
	}
	if that.AccessKeyId = section.Key("accessKeyId").String(); that.AccessKeyId == "" {
		panic(fmt.Errorf("place input AccessKeyId"))
	}
	if that.AccessKeySecret = section.Key("accessKeySecret").String(); that.AccessKeySecret == "" {
		panic(fmt.Errorf("place input AccessKeySecret"))
	}
}

func (that *AliyunEcs) GetName() string {
	return that.Name
}
func (that *AliyunEcs) Create(ip, name string) error {
	return aliyun.CreateEcs(that.Endpoint, that.RegionId, that.SecurityGroupId, that.AccessKeyId, that.AccessKeySecret, ip, name)
}
func (that *AliyunEcs) Delete(ip string) error {
	return aliyun.DeleteEcs(that.Endpoint, that.RegionId, that.SecurityGroupId, that.AccessKeyId, that.AccessKeySecret, ip)
}
