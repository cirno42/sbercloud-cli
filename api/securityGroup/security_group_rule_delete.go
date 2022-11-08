package securityGroup

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func DeleteSecurityGroupRule(projectID, securityGroupRuleID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/security-group-rules/%s", projectID, securityGroupRuleID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}
