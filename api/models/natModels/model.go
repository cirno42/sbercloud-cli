package natModels

type NatModel struct {
	RouterID              string `json:"router_id"`
	Status                string `json:"status"`
	Description           string `json:"description"`
	AdminStateUp          bool   `json:"admin_state_up"`
	TenantID              string `json:"tenant_id"`
	CreatedAt             string `json:"created_at"`
	Spec                  string `json:"spec"`
	InternalNetworkID     string `json:"internal_network_id"`
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	DnatRulesLimit        string `json:"dnat_rules_limit"`
	SnatRulePublicIPLimit string `json:"snat_rule_public_ip_limit"`
	EnterpriseProjectID   string `json:"enterprise_project_id"`
	BillingInfo           string `json:"billing_info"`
}
