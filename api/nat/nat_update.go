package nat

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func UpdateNAT(projectID, natId, name, desc, spec string) (*natModels.NatModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways/%s", projectID, natId)
	params := NatCreationParameters{
		Name:        name,
		Description: desc,
		Spec:        spec,
	}
	requestBody := NatCreationModel{
		NatGateway: params,
	}
	var createdNat natCreateResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, &requestBody, &createdNat, nil)
	return &createdNat.NatGateway, err
}
