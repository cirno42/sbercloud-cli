package subnetModels

type SubnetEntity struct {
	Subnet SubnetModel `json:"subnet"`
}

type SubnetModel struct {
	Id                string   `json:"id" header:"id"`
	Name              string   `json:"name" header:"name"`
	Description       string   `json:"description" header:"description"`
	Cidr              string   `json:"cidr" header:"cidr"`
	GatewayIp         string   `json:"gateway_ip" header:"gateway_ip"`
	Ipv6Enable        bool     `json:"ipv6_enable" header:"ipv6_enable"`
	DhcpEnable        bool     `json:"dhcp_enable" header:"dhcp_enable"`
	PrimaryDns        string   `json:"primary_dns" header:"primary_dns"`
	SecondaryDns      string   `json:"secondary_dns" header:"secondary_dns"`
	DnsList           []string `json:"dns_list"`
	AvailabilityZone  string   `json:"availability_zone" header:"availability_zone"`
	VpcId             string   `json:"vpc_id" header:"vpc_id"`
	Status            string   `json:"status" header:"status"`
	NeutronNetworkId  string   `json:"neutron_Network_id" header:"neutron_Network_id"`
	NeutronSubnetId   string   `json:"neutron_subnet_id" header:"neutron_subnet_id"`
	NeutronSubnetIdV6 string   `json:"neutron_subnet_id_v_6" header:"neutron_subnet_id_v_6"`
}
