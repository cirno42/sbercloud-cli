package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type listKeypairsResponse struct {
	Keypairs []getKeypairResponse `json:"keypairs"`
}

type getKeypairResponse struct {
	Keypair ecsModels.Keypair `json:"keypair"`
}

func ListKeypairs(projectID string) ([]ecsModels.Keypair, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-keypairs", projectID)
	var resp listKeypairsResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	pairs := make([]ecsModels.Keypair, len(resp.Keypairs))
	for i := 0; i < len(resp.Keypairs); i++ {
		pairs[i] = resp.Keypairs[i].Keypair
	}
	return pairs, err
}

func GetKeypair(projectID, keypairName string) (ecsModels.Keypair, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-keypairs/%s", projectID, keypairName)
	var resp getKeypairResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Keypair, err
}
