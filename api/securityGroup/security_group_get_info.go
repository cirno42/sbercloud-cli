package securityGroup

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/securityGroupModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type getInfoAboutSecGroupResponse struct {
	SecGroup securityGroupModels.SecurityGroupModel `json:"security_group"`
}

type getSecGroupListResponse struct {
	SecGroups []securityGroupModels.SecurityGroupModel `json:"security_groups"`
}

func GetInfoAboutSecurityGroup(projectID, securityGroupID string) (*securityGroupModels.SecurityGroupModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-groups/%s", projectID, securityGroupID)
	var secGroup getInfoAboutSecGroupResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &secGroup, nil)
	return &secGroup.SecGroup, err
}

func GetSecurityGroupsList(projectID string, limit int, marker, vpcID string) ([]securityGroupModels.SecurityGroupModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-groups?limit=%d&marker=%s&vpcID=%s", projectID, limit, marker, vpcID)
	var sgList getSecGroupListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &sgList, nil)
	return sgList.SecGroups, err
}

func GetInfoAboutSecurityGroupByName(projectID, name string) (*securityGroupModels.SecurityGroupModel, error) {
	secGroups, err := GetSecurityGroupsList(projectID, 0, "", "")
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(secGroups); i++ {
		if secGroups[i].Name == name {
			return &secGroups[i], nil
		}
	}
	return nil, errors.New("no security group with such name")
}
