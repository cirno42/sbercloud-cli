package ims

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/imsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type getImagesListResponse struct {
	Images []imsModels.ImageModel `json:"images"`
}

func GetImagesList(platform string) ([]imsModels.ImageModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.ImsEndpoint)+"/v2/cloudimages?__platform=%s", platform)
	var imgs getImagesListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &imgs, nil)
	return imgs.Images, err
}
