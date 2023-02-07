package vpcs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/vpcModels"
	"sbercloud-cli/internal/handlers/errorHandlers"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type vpcCreationParameters struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cidr        string `json:"cidr"`
}

type vpcCreationQuery struct {
	Vpc vpcCreationParameters `json:"vpc"`
}

func CreateVpc(projectID, name, description, cidr string) (*vpcModels.VpcModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs", projectID)

	vpc := vpcCreationParameters{
		Name:        name,
		Description: description,
		Cidr:        cidr,
	}
	vpcQuery := vpcCreationQuery{Vpc: vpc}
	var createdVpc vpcModels.VpcEntity
	var errorResp errorHandlers.ErrorResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &vpcQuery, &createdVpc, &errorResp)
	return &createdVpc.Vpc, err
}
