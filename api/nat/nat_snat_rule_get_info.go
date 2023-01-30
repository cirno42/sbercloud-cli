package nat

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
)

type snatRulesListResp struct {
	SnatRules []natModels.SnatRuleModel `json:"snat_rules"`
}

type snatRuleResp struct {
	SnatRule natModels.SnatRuleModel `json:"snat_rule"`
}

func ListSNATRules(projectID, natGatewayId, ipAddress string, limit int) ([]natModels.SnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/snat_rules?", projectID)
	if natGatewayId != "" {
		endpoint += "&nat_gateway_id=" + natGatewayId
	}
	if limit != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&limit=" + s
	}
	if ipAddress != "" {
		endpoint += "&floating_ip_address=" + ipAddress
	}
	var resp snatRulesListResp
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.SnatRules, err
}

func GetSNATRule(projectID, snatRuleId string) (natModels.SnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/snat_rules/%s", projectID, snatRuleId)
	var resp snatRuleResp
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.SnatRule, err
}
