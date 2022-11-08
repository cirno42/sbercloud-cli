package securityGroup

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func DeleteSecurityGroup(projectID, securityGroupID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-groups/%s", projectID, securityGroupID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeleteSecurityGroupByName(projectID, securityGroupName string) error {
	sg, err := GetInfoAboutSecurityGroupByName(projectID, securityGroupName)
	if err != nil {
		return err
	}
	err = DeleteSecurityGroup(projectID, sg.Id)
	return err
}
