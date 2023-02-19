package vpcUtils

import (
	"sbercloud-cli/api/vpcs"
)

func GetVpcId(id, name, projectID string) (string, error) {
	if id != "" {
		return id, nil
	}
	vpc, err := vpcs.GetVpcByName(projectID, name)
	if err != nil {
		return "", err
	}
	return vpc.Id, nil
}
