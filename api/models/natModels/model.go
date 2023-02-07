package natModels

type NatModel struct {
	RouterID              string `json:"router_id" header:"router_id"`
	Status                string `json:"status" header:"status"`
	Description           string `json:"description" header:"description"`
	AdminStateUp          bool   `json:"admin_state_up" header:"admin_state_up"`
	TenantID              string `json:"tenant_id" header:"tenant_id"`
	CreatedAt             string `json:"created_at" header:"created_at"`
	Spec                  string `json:"spec" header:"spec"`
	InternalNetworkID     string `json:"internal_network_id" header:"internal_network_id"`
	ID                    string `json:"id" header:"id"`
	Name                  string `json:"name" header:"name"`
	DnatRulesLimit        int    `json:"dnat_rules_limit" header:"dnat_rules_limit"`
	SnatRulePublicIPLimit int    `json:"snat_rule_public_ip_limit" header:"snat_rule_public_ip_limit"`
	EnterpriseProjectID   string `json:"enterprise_project_id" header:"enterprise_project_id"`
	BillingInfo           string `json:"billing_info" header:"billing_info"`
}

type SnatRuleModel struct {
	FloatingIPID      string `json:"floating_ip_id" header:"floating_ip_id"`
	Status            string `json:"status" header:"status"`
	NatGatewayID      string `json:"nat_gateway_id" header:"nat_gateway_id"`
	AdminStateUp      bool   `json:"admin_state_up" header:"admin_state_up"`
	NetworkID         string `json:"network_id" header:"network_id"`
	Description       string `json:"description" header:"description"`
	SourceType        int    `json:"source_type" header:"source_type"`
	TenantID          string `json:"tenant_id" header:"tenant_id"`
	CreatedAt         string `json:"created_at" header:"created_at"`
	ID                string `json:"id" header:"id"`
	FloatingIPAddress string `json:"floating_ip_address" header:"floating_ip_address"`
}

type DnatRuleModel struct {
	FloatingIPID             string `json:"floating_ip_id" header:"floating_ip_id"`
	Status                   string `json:"status" header:"status"`
	NatGatewayID             string `json:"nat_gateway_id" header:"nat_gateway_id"`
	AdminStateUp             bool   `json:"admin_state_up" header:"admin_state_up"`
	PortID                   string `json:"port_ID" header:"port_ID"`
	PrivateIP                string `json:"private_ip" header:"private_ip"`
	InternalServicePort      int    `json:"internal_service_port" header:"internal_service_port"`
	Protocol                 string `json:"protocol" header:"protocol"`
	TenantID                 string `json:"tenant_id" header:"tenant_id"`
	CreatedAt                string `json:"created_at" header:"created_at"`
	ID                       string `json:"id" header:"id"`
	FloatingIPAddress        string `json:"floating_ip_address" header:"floating_ip_address"`
	ExternalServicePort      int    `json:"external_service_port" header:"external_service_port"`
	Description              string `json:"description" header:"description"`
	ExternalServicePortRange string `json:"external_service_port_range" header:"external_service_port_range"`
	InternalServicePortRange string `json:"internal_service_port_range" header:"internal_service_port_range"`
}
