package ecs

import (
	"fmt"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type deleteNicId struct {
	SubnetId string `json:"id"`
}

type deleteNicsRequest struct {
	Nics []deleteNicId `json:"nics"`
}

func DeleteNicsBatchToEcs(projectID, serverId string, subnetIds []string) error {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/%s/nics/delete", projectID, serverId)
	subnets := make([]deleteNicId, len(subnetIds))
	for i := 0; i < len(subnetIds); i++ {
		subnets[i].SubnetId = subnetIds[i]
	}
	request := deleteNicsRequest{Nics: subnets}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, nil, nil)
	return err
}
