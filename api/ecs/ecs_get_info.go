package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type listEcsQueryingResponse struct {
	Count   int                  `json:"count"`
	Servers []ecsModels.ECSModel `json:"servers"`
}

type ecsQueryingResponse struct {
	Server ecsModels.ECSModel `json:"server"`
}

func GetECSList(projectID string) ([]ecsModels.ECSModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/detail", projectID)
	var ecsArray listEcsQueryingResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &ecsArray, nil)
	return ecsArray.Servers, err
}

func GetInfoAboutEcs(projectID, ecsID string) (ecsModels.ECSModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s", projectID, ecsID)
	var queriedEcs ecsQueryingResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &queriedEcs, nil)
	return queriedEcs.Server, err
}
