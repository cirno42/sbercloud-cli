package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
)

type getResizeFlavorsResponse struct {
	Flavors []ecsModels.FlavorModel `json:"flavors"`
}

func ListResizeFlavors(projectID, sortKey, sortDir, marker, ecsID string, limit int) ([]ecsModels.FlavorModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/resize_flavors?", projectID)
	if ecsID != "" {
		endpoint += "instance_uuid=" + ecsID
	}
	if sortKey != "" {
		endpoint += "&sort_key=" + sortKey
	}
	if sortDir != "" {
		endpoint += "&sort_dir=" + sortDir
	}
	if marker != "" {
		endpoint += "&marker=" + marker
	}
	if limit != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&limit=" + s
	}
	var resp getResizeFlavorsResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Flavors, err
}
