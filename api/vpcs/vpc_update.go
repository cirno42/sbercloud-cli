package vpcs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/vpcModels"
	"sbercloud-cli/internal/handlers/errorHandlers"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func UpdateVpc(projectID, id, name, description, cidr string) (*vpcModels.VpcEntity, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/vpcs/%s", projectID, id)

	vpc := vpcCreationParameters{
		Name:        name,
		Description: description,
		Cidr:        cidr,
	}
	vpcQuery := vpcCreationQuery{Vpc: vpc}
	var createdVpc vpcModels.VpcEntity
	var errorResp errorHandlers.ErrorResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_PUT, &vpcQuery, &createdVpc, &errorResp)
	return &createdVpc, err

}
