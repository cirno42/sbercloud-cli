package nat

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type createDNATRuleRequest struct {
	DnatRule createDNATRuleParameters `json:"dnat_rule"`
}

type createDNATRuleParameters struct {
	FloatingIPID             string `json:"floating_ip_id"`
	NatGatewayID             string `json:"nat_gateway_id"`
	PortID                   string `json:"port_id"`
	InternalServicePort      int    `json:"internal_service_port"`
	Protocol                 string `json:"protocol"`
	ExternalServicePort      int    `json:"external_service_port"`
	Description              string `json:"description"`
	InternalServicePortRange string `json:"internal_service_port_range"`
	ExternalServicePortRange string `json:"external_service_port_range"`
}

type createDNATRuleResponse struct {
	DnatRule natModels.DnatRuleModel `json:"dnat_rule"`
}

func AddDNATRule(projectID, natID, portID, floatingIpId, protocol, description, internalServicePortRange, externalServicePortRange string, internalServicePort, externalServicePort int) (natModels.DnatRuleModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/dnat_rules", projectID)
	params := createDNATRuleParameters{
		FloatingIPID:             floatingIpId,
		PortID:                   portID,
		NatGatewayID:             natID,
		InternalServicePort:      internalServicePort,
		Protocol:                 protocol,
		ExternalServicePort:      externalServicePort,
		Description:              description,
		InternalServicePortRange: internalServicePortRange,
		ExternalServicePortRange: externalServicePortRange,
	}
	requestBody := createDNATRuleRequest{DnatRule: params}
	var resp createDNATRuleResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &requestBody, &resp, nil)
	return resp.DnatRule, err
}
