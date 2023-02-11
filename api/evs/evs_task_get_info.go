package evs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/evsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func GetInfoAboutBatchTask(projectID, jobID string) (*evsModels.CreateBatchDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job evsModels.CreateBatchDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}

func GetInfoAboutSingleTask(projectID, jobID string) (*evsModels.CreateSingleDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v1/%s/jobs/%s", projectID, jobID)
	var job evsModels.CreateSingleDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &job, nil)
	return &job, err
}
