package eip

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/eipModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type eipAssignRequest struct {
	PublicIP  eipAssignPublicIP  `json:"publicip"`
	Bandwidth eipAssignBandwidth `json:"bandwidth"`
}

type eipAssignBandwidth struct {
	Name      string `json:"name"`
	Size      int    `json:"size"`
	ShareType string `json:"share_type"`
}

type eipAssignPublicIP struct {
	EipType   string `json:"type"`
	IpVersion int    `json:"ip_version"`
}

type eipAssignResponse struct {
	EipEntity eipModels.EipModel `json:"publicip"`
}

func AssignEIP(projectID string, eipType string, ipVersion int, bandwidthName string, size int, shareType string) (*eipModels.EipModel, error) {
	if (ipVersion != 4) && (ipVersion != 6) {
		return nil, errors.New("wrong version of IP")
	}
	ip := eipAssignPublicIP{
		EipType:   eipType,
		IpVersion: ipVersion,
	}

	bandwidth := eipAssignBandwidth{
		Name:      bandwidthName,
		Size:      size,
		ShareType: shareType,
	}
	request := eipAssignRequest{
		PublicIP:  ip,
		Bandwidth: bandwidth,
	}
	var createdEIP eipAssignResponse
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips", projectID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &request, &createdEIP, nil)
	return &createdEIP.EipEntity, err
}
