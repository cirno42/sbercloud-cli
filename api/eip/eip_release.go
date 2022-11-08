package eip

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func DeletePublicIP(projectID, publicIpID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/publicips/%s", projectID, publicIpID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeletePublicIPByAddress(projectID, publicIpAddress string) error {
	eip, err := GetInfoAboutEIPByAddress(projectID, publicIpAddress)
	if err != nil {
		return err
	}
	if eip == nil {
		return errors.New("No public IP with such address" + publicIpAddress)
	}
	err = DeletePublicIP(projectID, eip.Id)
	return err
}
