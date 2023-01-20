package securityGroupModels

type SecurityGroupModel struct {
	Name               string                   `json:"name" header:"name"`
	Description        string                   `json:"description" header:"description"`
	Id                 string                   `json:"id" header:"id"`
	VpcID              string                   `json:"vpc_id" header:"vpc_id"`
	SecurityGroupRules []SecurityGroupRuleModel `json:"security_group_rules" header:"security_group_rules"`
}

type SecurityGroupRuleModel struct {
	Id              string `json:"id"`
	Description     string `json:"description"`
	SecurityGroupID string `json:"security_group_id"`
	Direction       string `json:"direction"`
	Ethertype       string `json:"ethertype"`
	Protocol        string `json:"protocol"`
	PortRangeMin    int    `json:"port_range_min"`
	PortRangeMax    int    `json:"port_range_max"`
	RemoteIpPrefix  string `json:"remote_ip_prefix"`
	RemoteGroupId   string `json:"remote_group_id"`
	TenantId        string `json:"tenant_id"`
}
