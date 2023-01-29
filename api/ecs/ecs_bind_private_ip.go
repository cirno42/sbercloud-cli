package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type bindPrivateIpParameters struct {
	SubnetID       string `json:"subnet_id"`
	IPAddress      string `json:"ip_address"`
	ReverseBinding bool   `json:"reverse_binding"`
}

type bindPrivateIpRequest struct {
	Nic bindPrivateIpParameters `json:"nic"`
}

func BindPrivateIp(projectID, nicID, subnetID, ipAddress string, reverseBinding bool) (ecsModels.BindPrivateIpResponse, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/nics/%s", projectID, nicID)
	params := bindPrivateIpParameters{
		SubnetID:       subnetID,
		IPAddress:      ipAddress,
		ReverseBinding: reverseBinding,
	}
	req := bindPrivateIpRequest{Nic: params}
	var resp ecsModels.BindPrivateIpResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, req, &resp, nil)
	return resp, err
}
