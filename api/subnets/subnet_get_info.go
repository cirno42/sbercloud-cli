package subnets

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/subnetModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type subnetQueryingResponse struct {
	Subnets []subnetModels.SubnetModel `json:"subnets"`
}

func GetSubnetsList(projectID string, limit int, marker, vpcID string) ([]subnetModels.SubnetModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/subnets?limit=%d&marker=%s&vpc_id=%s", projectID, limit, marker, vpcID)
	var subnetsArray subnetQueryingResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &subnetsArray, nil)
	return subnetsArray.Subnets, err
}

func GetInfoAboutSubnet(projectID, subnetID string) (*subnetModels.SubnetModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/subnets/%s", projectID, subnetID)
	var subnet subnetModels.SubnetEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &subnet, nil)
	return &subnet.Subnet, err
}

func GetSubnetByName(projectID, subnetName string) (*subnetModels.SubnetModel, error) {
	subnetsList, err := GetSubnetsList(projectID, 0, "", "")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(subnetsList); i++ {
		if subnetsList[i].Name == subnetName {
			return &subnetsList[i], nil
		}
	}
	return nil, errors.New("no subnet with such name")
}

func GetSubnetsByNames(projectID string, subnetNames []string) ([]subnetModels.SubnetModel, error) {
	subnetsList, err := GetSubnetsList(projectID, 0, "", "")
	if err != nil {
		return nil, err
	}
	res := make([]subnetModels.SubnetModel, 0)
	names := make(map[string]bool)
	for _, subnetName := range subnetNames {
		names[subnetName] = true
	}
	for _, subnet := range subnetsList {
		if names[subnet.Name] {
			res = append(res, subnet)
		}
	}
	return res, nil
}
