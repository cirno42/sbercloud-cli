package ecs

import (
	"fmt"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type volumeAttachmentRequestParameters struct {
	VolumeId string `json:"volumeId"`
	Device   string `json:"device"`
}

type volumeAttachmentRequest struct {
	VolumeAttachment volumeAttachmentRequestParameters `json:"volumeAttachment"`
}

func AttachDiskEcs(projectID, ecsId, volumeId, device string) (ecsModels.ECSJob, error) {
	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers/%s/attachvolume", projectID, ecsId)
	params := volumeAttachmentRequestParameters{
		VolumeId: volumeId,
		Device:   device,
	}
	request := volumeAttachmentRequest{
		VolumeAttachment: params,
	}
	var response ecsModels.ECSJob
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, request, &response, nil)
	return response, err
}
