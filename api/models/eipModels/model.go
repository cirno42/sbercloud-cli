package eipModels

type EipModel struct {
	Id                  string `json:"id"`
	EipType             string `json:"type"`
	PublicIP            string `json:"public_ip_address"`
	Status              string `json:"status"`
	TenantID            string `json:"tenant_id"`
	CreateTime          string `json:"create_time"`
	BandwidthSize       int    `json:"bandwidth_size"`
	EnterpriseProjectID string `json:"enterprise_project_id"`
	IPVersion           int    `json:"ip_version"`
	PrivateIP           string `json:"private_ip_address"`
	PortID              string `json:"port_id"`
	BandwidthID         string `json:"bandwidth_id"`
	BandwidthShareType  string `json:"bandwidth_share_type"`
	BandwidthName       string `json:"bandwidth_name"`
}
