package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
	"sync"
)

type listEcsQueryingResponse struct {
	Count   int                  `json:"count"`
	Servers []ecsModels.ECSModel `json:"servers"`
}

type ecsQueryingResponse struct {
	Server ecsModels.ECSModel `json:"server"`
}

func GetECSList(projectID string, offset, limit int) ([]ecsModels.ECSModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/detail?", projectID)
	if offset != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&offset=" + s
	}
	if limit != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&limit=" + s
	}
	var ecsArray listEcsQueryingResponse
	err := requestMakers.CreateAndDoRequestInSpecifiedProject(endpoint, projectID, requestMakers.HTTP_METHOD_GET, nil, &ecsArray)
	return ecsArray.Servers, err
}

func GetInfoAboutEcs(projectID, ecsID string) (ecsModels.ECSModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s", projectID, ecsID)
	var queriedEcs ecsQueryingResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &queriedEcs, nil)
	return queriedEcs.Server, err
}

func GetListEcsById(projectID string, ecsIds []string) ([]ecsModels.ECSModel, error) {

	ecses := make([]ecsModels.ECSModel, len(ecsIds))
	var wg sync.WaitGroup
	for i, id := range ecsIds {
		wg.Add(1)
		go func(projectID, ecsID string, ecs *ecsModels.ECSModel) {
			defer wg.Done()
			server, _ := GetInfoAboutEcs(projectID, ecsID)
			*ecs = server
		}(projectID, id, &ecses[i])
	}
	wg.Wait()
	return ecses, nil
}
