package evs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type volumeRequest struct {
	AvailabilityZone string `json:"availability_zone"`
	Name             string `json:"name"`
	VolumeType       string `json:"volume_type"`
	Count            int    `json:"count"`
	Size             int    `json:"size"`
	Multiattach      bool   `json:"multiattach"`
}

type createVolumeRequest struct {
	Volume volumeRequest `json:"volume"`
}

func CreateDisk(projectId, name, volumeType, availabilityZone string, count, size int, multiattach bool) (*ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EvsEndpoint)+"/v2/%s/cloudvolumes", projectId)
	req := volumeRequest{
		Name:             name,
		AvailabilityZone: availabilityZone,
		VolumeType:       volumeType,
		Count:            count,
		Size:             size,
		Multiattach:      multiattach,
	}
	request := createVolumeRequest{
		Volume: req,
	}
	var createdJob ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &createdJob, nil)
	return &createdJob, err
}
