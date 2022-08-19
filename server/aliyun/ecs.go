package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

func createClient(endpoint, accessKeyId, accessKeySecret string) (result *ecs20140526.Client, err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: tea.String(accessKeyId),
		// 您的 AccessKey Secret
		AccessKeySecret: tea.String(accessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	result = &ecs20140526.Client{}
	result, err = ecs20140526.NewClient(config)
	return
}

func CreateEcs(endpoint, regionId, securityGroupId, accessKeyId, accessKeySecret, ip, name string) error {
	client, err := createClient(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return err
	}

	authorizeSecurityGroupRequest := &ecs20140526.AuthorizeSecurityGroupRequest{}

	abc := ecs20140526.AuthorizeSecurityGroupRequestPermissions{
		SourceCidrIp: tea.String(ip),
		IpProtocol:   tea.String("ALL"),
		PortRange:    tea.String("-1/-1"),
		Description:  tea.String(name),
	}

	authorizeSecurityGroupRequest.RegionId = tea.String(regionId)
	authorizeSecurityGroupRequest.SecurityGroupId = tea.String(securityGroupId)
	authorizeSecurityGroupRequest.Permissions = append(authorizeSecurityGroupRequest.Permissions, &abc)

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err := client.AuthorizeSecurityGroupWithOptions(authorizeSecurityGroupRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	return tryErr
}

func DeleteEcs(endpoint, regionId, securityGroupId, accessKeyId, accessKeySecret, ip string) error {
	client, err := createClient(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return err
	}

	authorizeSecurityGroupRequest := &ecs20140526.RevokeSecurityGroupRequest{}

	abc := ecs20140526.RevokeSecurityGroupRequestPermissions{
		SourceCidrIp: tea.String(ip),
		IpProtocol:   tea.String("ALL"),
		PortRange:    tea.String("-1/-1"),
	}

	authorizeSecurityGroupRequest.RegionId = tea.String(regionId)
	authorizeSecurityGroupRequest.SecurityGroupId = tea.String(securityGroupId)
	authorizeSecurityGroupRequest.Permissions = append(authorizeSecurityGroupRequest.Permissions, &abc)

	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err := client.RevokeSecurityGroupWithOptions(authorizeSecurityGroupRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	return tryErr
}
