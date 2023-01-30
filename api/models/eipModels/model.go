package eipModels

type EipModel struct {
	Id                  string `json:"id" header:"id"`
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
