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
