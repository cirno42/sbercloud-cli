package securityGroup

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/securityGroupModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type securityGroupCreateRequest struct {
	SecGroup securityGroupCreateParameters `json:"security_group"`
}

type securityGroupCreateParameters struct {
	Name string `json:"name"`
}

type securityGroupCreateResponse struct {
	SecGroup securityGroupModels.SecurityGroupModel `json:"security_group"`
}

func CreateSecurityGroup(projectID, name string) (*securityGroupModels.SecurityGroupModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-groups", projectID)
	sgParams := securityGroupCreateParameters{Name: name}
	sgRequest := securityGroupCreateRequest{SecGroup: sgParams}
	var createdSG securityGroupCreateResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &sgRequest, &createdSG, nil)
	return &createdSG.SecGroup, err
}
