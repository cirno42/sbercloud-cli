package nat

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type getInfoAboutNatResponse struct {
	NatGateway natModels.NatModel `json:"nat_gateway"`
}
type getNatListResponse struct {
	NatGateways []natModels.NatModel `json:"nat_gateways"`
}

func GetInfoAboutNat(projectID, natID string) (*natModels.NatModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways/%s", projectID, natID)
	var natResp getInfoAboutNatResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &natResp, nil)
	return &natResp.NatGateway, err
}

func GetNatList(projectID string) ([]natModels.NatModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways", projectID)
	var natResp getNatListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &natResp, nil)
	return natResp.NatGateways, err
}

func GetNatByName(projectID, name string) (*natModels.NatModel, error) {
	nats, err := GetNatList(projectID)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(nats); i++ {
		if nats[i].Name == name {
			return &nats[i], nil
		}
	}
	return nil, errors.New("No such NAT")
}
