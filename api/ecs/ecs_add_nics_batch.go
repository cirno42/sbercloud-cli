package ecs

import (
	"fmt"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type nicRequest struct {
	SubnetId string `json:"subnet_id"`
}

type nics struct {
	Nics []nicRequest `json:"nics"`
}

func AddNicsBatchToEcs(projectID, serverId string, subnetIds []string) error {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/%s/nics", projectID, serverId)
	subnets := make([]nicRequest, len(subnetIds))
	for i := 0; i < len(subnetIds); i++ {
		subnets[i].SubnetId = subnetIds[i]
	}
	request := nics{Nics: subnets}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, nil, nil)
	return err
}
