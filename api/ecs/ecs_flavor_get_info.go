package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type getECSFlavorList struct {
	Flavors []ecsModels.FlavorModel `json:"flavors"`
}

func GetESCFlavorList(projectID, availabilityZone string) ([]ecsModels.FlavorModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/flavors?availability_zone=%s", projectID, availabilityZone)
	var flavorsArray getECSFlavorList
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &flavorsArray, nil)
	return flavorsArray.Flavors, err
}
