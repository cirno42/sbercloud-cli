package nat

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

func DeleteNat(projectID, natID string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.NatEndpoint)+"/v2/%s/nat_gateways/%s", projectID, natID)
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_DELETE, nil, nil, nil)
	return err
}

func DeleteNatByName(projectID, natName string) error {
	nat, err := GetNatByName(projectID, natName)
	if err != nil {
		err = DeleteNat(projectID, nat.ID)
	}
	return err
}
