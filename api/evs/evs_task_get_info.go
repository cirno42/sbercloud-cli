package evs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func GetInfoAboutTask(projectID, jobID string) (*ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}
