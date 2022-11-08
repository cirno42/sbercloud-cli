package eip

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/eipModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type eipQueryingResponse struct {
	PublicIPs []eipModels.EipModel `json:"publicips"`
}

func GetInfoAboutEIPByAddress(projectID, ipAddress string) (*eipModels.EipModel, error) {
	ipArray, err := GetEIPsList(projectID, 1000, "")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ipArray); i++ {
		if ipArray[i].PublicIP == ipAddress {
			return &ipArray[i], nil
		}
	}
	return nil, nil
}

func GetEIPInfo(projectID, publicIpId string) (*eipModels.EipModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips/%s", projectID, publicIpId)
	var publicIP eipModels.EipModel
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &publicIP, nil)
	return &publicIP, err
}

func GetEIPsList(projectID string, limit int, marker string) ([]eipModels.EipModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips?limit=%d&marker=%s", projectID, limit, marker)
	var ipsArray eipQueryingResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &ipsArray, nil)
	return ipsArray.PublicIPs, err
}
