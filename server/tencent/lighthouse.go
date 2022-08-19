package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
)

var (
	Protocol = "ALL"
	Port     = "ALL"
	Action   = "ACCEPT"
)

func CreateLighthouse(endpoint, region, instanceId, secretId, secretKey, ip, name string) error {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, _ := lighthouse.NewClient(credential, region, cpf)

	request := lighthouse.NewCreateFirewallRulesRequest()
	request.InstanceId = &instanceId
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

func DeleteLighthouse(endpoint, region, instanceId, secretId, secretKey, ip string) error {
	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpoint
	client, _ := lighthouse.NewClient(credential, region, cpf)

	request := lighthouse.NewDeleteFirewallRulesRequest()
	request.InstanceId = &instanceId
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
