package eipModels

type EipModel struct {
	ID                  string `json:"id" header:"id"`
	EipType             string `json:"type" header:"type"`
	PublicIPAddress     string `json:"public_ip_address" header:"public_ip_address"`
	Status              string `json:"status" header:"status"`
	TenantID            string `json:"tenant_id" header:"tenant_id"`
	CreateTime          string `json:"create_time" header:"create_time"`
	BandwidthSize       int    `json:"bandwidth_size" header:"bandwidth_size"`
	EnterpriseProjectID string `json:"enterprise_project_id" header:"enterprise_project_id"`
	IPVersion           int    `json:"ip_version" header:"ip_version"`
	PrivateIP           string `json:"private_ip_address" header:"private_ip_address"`
	PortID              string `json:"port_id" header:"port_id"`
	BandwidthID         string `json:"bandwidth_id" header:"bandwidth_id"`
	BandwidthShareType  string `json:"bandwidth_share_type" header:"bandwidth_share_type"`
	BandwidthName       string `json:"bandwidth_name" header:"bandwidth_name"`
}

type ActiveEIP struct {
	ID                 string `json:"id" header:"id"`
	Address            string `json:"address" header:"address"`
	InstanceID         string `json:"instance_id" header:"instance_id"`
	InstanceType       string `json:"instance_type" header:"instance_type"`
	ParentInstanceID   string `json:"parent_instance" header:"parent_instance"`
	ParentInstanceType string `json:"parent_instance_type" header:"parent_instance_type"`
}

type ProjectActiveEIP struct {
	ProjectID   string      `json:"project_id" header:"project_id"`
	ProjectName string      `json:"project_name" header:"project_name"`
	ActiveIP    []ActiveEIP `json:"active_ip"`
}
