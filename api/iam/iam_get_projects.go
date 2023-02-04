package iam

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/iamModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type listProjectResponse struct {
	Projects []iamModels.ProjectModel `json:"projects"`
}

func ListProjects() ([]iamModels.ProjectModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.IamEndpoint) + "/v3/auth/projects")
	var resp listProjectResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Projects, err
}
