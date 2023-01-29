package evs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func DeleteDisk(projectID, volumeID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/volumes/%s", projectID, volumeID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}
