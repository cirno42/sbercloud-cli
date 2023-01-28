package ecs

import (
	"fmt"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type batchStopEcs struct {
	Type    string      `json:"type"`
	Servers []serverIds `json:"servers"`
}

type batchStopRequestEcs struct {
	ServersToStop batchStopEcs `json:"os-stop"`
}

func BatchStopEcs(projectID, stopType string, ecsIdsToStop []string) (ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/action", projectID)
	serverIds := make([]serverIds, len(ecsIdsToStop))
	for i := 0; i < len(ecsIdsToStop); i++ {
		serverIds[i].Id = ecsIdsToStop[i]
	}
	restartReq := batchStopEcs{Type: stopType, Servers: serverIds}
	request := batchStopRequestEcs{ServersToStop: restartReq}
	var response ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &response, nil)
	return response, err
}
