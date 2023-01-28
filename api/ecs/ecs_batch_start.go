package ecs

import (
	"fmt"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type serverIds struct {
	Id string `json:"id"`
}

type serverToStart struct {
	Servers []serverIds `json:"servers"`
}

type batchStartEcs struct {
	ServersToStart serverToStart `json:"os-start"`
}

func BatchStartEcs(projectID string, ecsIdsToStart []string) (ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/action", projectID)
	serverIds := make([]serverIds, len(ecsIdsToStart))
	for i := 0; i < len(ecsIdsToStart); i++ {
		serverIds[i].Id = ecsIdsToStart[i]
	}
	servers := serverToStart{Servers: serverIds}
	request := batchStartEcs{ServersToStart: servers}
	var response ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &response, nil)
	return response, err
}
