package securityGroup

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/securityGroupModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type getInfoAboutSecGroupRule struct {
	SecurityGroupRule securityGroupModels.SecurityGroupRuleModel `json:"security_group_rule"`
}

type getSecGroupRulesListResponse struct {
	SecurityGroupRules []securityGroupModels.SecurityGroupRuleModel `json:"security_group_rules"`
}

func GetInfoAboutSecurityGroupRule(projectID, securityGroupRuleID string) (*securityGroupModels.SecurityGroupRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-group-rules/%s", projectID, securityGroupRuleID)
	var secGroupRule getInfoAboutSecGroupRule
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &secGroupRule, nil)
	return &secGroupRule.SecurityGroupRule, err
}

func GetSecurityGroupRulesList(projectID string, limit int, marker, securityGroupId string) ([]securityGroupModels.SecurityGroupRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-group-rules?limit=%d&marker=%s&security_group_id=%s", projectID, limit, marker, securityGroupId)
	var sgListRules getSecGroupRulesListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &sgListRules, nil)
	return sgListRules.SecurityGroupRules, err
}

func GetSecurityGroupRulesListBySGName(projectID string, limit int, marker string, securityGroupName string) ([]securityGroupModels.SecurityGroupRuleModel, error) {
	sg, err := GetInfoAboutSecurityGroupByName(projectID, securityGroupName)
	if err != nil {
		return nil, err
	}
	rulesList, err := GetSecurityGroupRulesList(projectID, limit, marker, sg.Id)
	return rulesList, err
}
