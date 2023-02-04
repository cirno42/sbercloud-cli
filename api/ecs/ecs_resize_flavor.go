package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type resizeRequestParameters struct {
	FlavorRef string `json:"flavorRef"`
}

type resizeRequest struct {
	Resize resizeRequestParameters `json:"resize"`
	DryRun bool                    `json:"dry_run"`
}

func ResizeECS(projectID, ecsID, flavorRef string, dryRun bool) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/resize", projectID, ecsID)
	params := resizeRequestParameters{FlavorRef: flavorRef}
	req := resizeRequest{Resize: params, DryRun: dryRun}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, req, nil, nil)
	return err
}
