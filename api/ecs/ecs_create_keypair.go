package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type createKeyPairRequest struct {
	KeyPair createKeyPairParams `json:"keypair"`
}

type createKeyPairParams struct {
	Name string `json:"name"`
}

type importKeyPairRequest struct {
	KeyPair importKeyPairParams `json:"keypair"`
}

type importKeyPairParams struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

type createKeyPairResponse struct {
	KeyPair ecsModels.Keypair `json:"keypair"`
}

func CreateKeyPair(projectID, name string) (ecsModels.Keypair, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-keypairs", projectID)
	params := createKeyPairParams{Name: name}
	req := createKeyPairRequest{KeyPair: params}
	var resp createKeyPairResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, &resp, nil)
	return resp.KeyPair, err
}

func ImportKeyPair(projectID, name, publicKey string) (ecsModels.Keypair, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v2.1/%s/os-keypairs", projectID)
	params := importKeyPairParams{
		Name:      name,
		PublicKey: publicKey,
	}
	req := importKeyPairRequest{KeyPair: params}
	var resp createKeyPairResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, &resp, nil)
	return resp.KeyPair, err
}
