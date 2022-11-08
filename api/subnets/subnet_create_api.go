package subnets

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/subnetModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type subnetCreationQuery struct {
	Subnet subnetCreationParameters `json:"subnet"`
}

type subnetCreationParameters struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Cidr             string   `json:"cidr"`
	GatewayIp        string   `json:"gateway_ip"`
	Ipv6Enable       bool     `json:"ipv6_enable"`
	DhcpEnable       bool     `json:"dhcp_enable"`
	PrimaryDns       string   `json:"primary_dns"`
	SecondaryDns     string   `json:"secondary_dns"`
	DnsList          []string `json:"dns_list"`
	AvailabilityZone string   `json:"availability_zone"`
	VpcId            string   `json:"vpc_id"`
}

func CreateSubnet(projectID, name, description, cidr, gatewayIp string, ipv6Enable, dhcpEnable bool,
	primaryDns, secondaryDns string, dnsList []string, availabilityZones, vpcId string) (*subnetModels.SubnetEntity, error) {

	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/subnets", projectID)
	subnetParameters := subnetCreationParameters{
		Name:             name,
		Description:      description,
		Cidr:             cidr,
		GatewayIp:        gatewayIp,
		Ipv6Enable:       ipv6Enable,
		DhcpEnable:       dhcpEnable,
		PrimaryDns:       primaryDns,
		SecondaryDns:     secondaryDns,
		DnsList:          dnsList,
		AvailabilityZone: availabilityZones,
		VpcId:            vpcId,
	}
	subnetQuery := subnetCreationQuery{Subnet: subnetParameters}
	var createdSubnet subnetModels.SubnetEntity
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &subnetQuery, &createdSubnet, nil)
	return &createdSubnet, err
}
