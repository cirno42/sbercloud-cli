package nat

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
)

type listDNATRules struct {
	DnatRules []natModels.DnatRuleModel `json:"dnat_rules"`
}

type getDNATRule struct {
	DnatRule natModels.DnatRuleModel `json:"dnat_rule"`
}

func ListDNATRules(projectID, natID, portID, eipId, eipAddress, protocol, internalPortRange, externalPortRange string, internalServicePort, externalServicePort int) ([]natModels.DnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/dnat_rules?", projectID)
	if natID != "" {
		endpoint += "&nat_gateway_id=" + natID
	}
	if portID != "" {
		endpoint += "&port_id=" + portID
	}
	if eipId != "" {
		endpoint += "&floating_ip_id=" + eipId
	}
	if eipAddress != "" {
		endpoint += "&floating_ip_address=" + eipAddress
	}
	if internalServicePort != 0 {
		s := strconv.FormatInt(int64(internalServicePort), 10)
		endpoint += "&internal_service_port=" + s
	}
	if externalServicePort != 0 {
		s := strconv.FormatInt(int64(externalServicePort), 10)
		endpoint += "&external_service_port=" + s
	}
	if protocol != "" {
		endpoint += "&protocol=" + protocol
	}
	if internalPortRange != "" {
		endpoint += "&internal_service_port_range=" + internalPortRange
	}
	if externalPortRange != "" {
		endpoint += "&external_service_port_range=" + externalPortRange
	}
	var resp listDNATRules
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.DnatRules, err
}

func GetDNATRule(projectID, ruleID string) (natModels.DnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/dnat_rules/%s", projectID, ruleID)
	var resp getDNATRule
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.DnatRule, err

}
