package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type VolumeInfoResponse struct {
	Volume ecsModels.VolumeAttachments `json:"volume_attachments"`
}

func GetListAttachedDisks(projectID, serverID string) (ecsModels.AttachedDisks, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/block_device", projectID, serverID)
	var disks ecsModels.AttachedDisks
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &disks, nil)
	return disks, err
}

func GetInfoAboutAttachedDisk(projectID, serverID, volumeID string) (ecsModels.VolumeAttachments, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/block_device/%s", projectID, serverID, volumeID)
	var disk VolumeInfoResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &disk, nil)
	return disk.Volume, err
}
