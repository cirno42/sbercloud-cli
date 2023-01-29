package evs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/evsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
)

type evsListResponse struct {
	Volumes []evsModels.EvsModel `json:"volumes"`
}

func GetDisksList(projectID, status string, limit, offset int) ([]evsModels.EvsModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/volumes/detail?", projectID)
	if status != "" {
		endpoint += "&status=" + status
	}
	if limit != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&limit=" + s
	}
	if offset != 0 {
		s := strconv.FormatInt(int64(limit), 10)
		endpoint += "&offset=" + s
	}
	var resp evsListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Volumes, err
}
