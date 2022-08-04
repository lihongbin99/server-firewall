package tencent

import (
	"security-network/common/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

var (
	Protocol = "ALL"
	Port     = "ALL"
	Action   = "ACCEPT"
)

func Create(ip, name string) error {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(credential, "ap-guangzhou", cpf)

	request := lighthouse.NewCreateFirewallRulesRequest()
	request.InstanceId = &config.InstanceId
	firewallRules := []*lighthouse.FirewallRule{{
		Protocol:                &Protocol,
		Port:                    &Port,
		CidrBlock:               &ip,
		Action:                  &Action,
		FirewallRuleDescription: &name,
	}}
	request.FirewallRules = firewallRules

	_, err := client.CreateFirewallRules(request)
	return err
}

func Delete(ip string) error {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "lighthouse.tencentcloudapi.com"
	client, _ := lighthouse.NewClient(credential, "ap-guangzhou", cpf)

	request := lighthouse.NewDeleteFirewallRulesRequest()
	request.InstanceId = &config.InstanceId
	firewallRules := []*lighthouse.FirewallRule{{
		Protocol:  &Protocol,
		Port:      &Port,
		CidrBlock: &ip,
		Action:    &Action,
	}}
	request.FirewallRules = firewallRules

	_, err := client.DeleteFirewallRules(request)
	return err
}
