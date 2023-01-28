package ecs

import (
	"fmt"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type batchRestartEcs struct {
	Type    string      `json:"type"`
	Servers []serverIds `json:"servers"`
}

type batchRestartRequestEcs struct {
	ServersToStart batchRestartEcs `json:"reboot"`
}

func BatchRestartEcs(projectID, restartType string, ecsIdsToRestart []string) (ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/action", projectID)
	serverIds := make([]serverIds, len(ecsIdsToRestart))
	for i := 0; i < len(ecsIdsToRestart); i++ {
		serverIds[i].Id = ecsIdsToRestart[i]
	}
	restartReq := batchRestartEcs{Type: restartType, Servers: serverIds}
	request := batchRestartRequestEcs{ServersToStart: restartReq}
	var response ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &response, nil)
	return response, err
}
