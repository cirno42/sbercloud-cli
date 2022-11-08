package vpcs

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/vpcModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type vpcsQueringResponse struct {
	Vpcs []vpcModels.VpcModel `json:"vpcs"`
}

func GetVpcsList(projectID string, limit int, marker string) ([]vpcModels.VpcModel, error) {

	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs?limit=%d&marker=%s", projectID, limit, marker)
	var vpcsArray vpcsQueringResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &vpcsArray, nil)

	return vpcsArray.Vpcs, err
}

func GetInfoAboutVpc(projectID, vpcID string) (*vpcModels.VpcModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s", projectID, vpcID)
	var returnedVpc vpcModels.VpcEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &returnedVpc, nil)
	return &returnedVpc.Vpc, err
}

func GetVpcByName(projectID, name string) (*vpcModels.VpcModel, error) {
	vpcs, err := GetVpcsList(projectID, 0, "")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(vpcs); i++ {
		if name == vpcs[i].Name {
			return &vpcs[i], nil
		}
	}
	return nil, errors.New("No such VPC")
}
