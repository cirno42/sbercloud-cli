package subnetModels

type SubnetEntity struct {
	Subnet SubnetModel `json:"subnet"`
}

type SubnetModel struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Cidr              string   `json:"cidr"`
	GatewayIp         string   `json:"gateway_ip"`
	Ipv6Enable        bool     `json:"ipv6_enable"`
	DhcpEnable        bool     `json:"dhcp_enable"`
	PrimaryDns        string   `json:"primary_dns"`
	SecondaryDns      string   `json:"secondary_dns"`
	DnsList           []string `json:"dns_list"`
	AvailabilityZone  string   `json:"availability_zone"`
	VpcId             string   `json:"vpc_id"`
	Status            string   `json:"status"`
	NeutronNetworkId  string   `json:"neutron_Network_id"`
	NeutronSubnetId   string   `json:"neutron_subnet_id"`
	NeutronSubnetIdV6 string   `json:"neutron_subnet_id_v_6"`
}
