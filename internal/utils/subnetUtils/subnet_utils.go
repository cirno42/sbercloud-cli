package subnetUtils

import "github.com/cirno42/sbercloud-api/api/subnets"

func GetSubnetId(subnetId, subnetName, projectId string) (string, error) {
	if subnetId != "" {
		return subnetId, nil
	}
	s, err := subnets.GetSubnetByName(projectId, subnetName)
	if err == nil {
		return s.Id, err
	} else {
		return "", err
	}
}
