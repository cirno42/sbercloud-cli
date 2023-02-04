package iamModels

type ProjectModel struct {
	IsDomain    bool   `json:"is_domain"`
	Description string `json:"description"`
	Links       struct {
		Self string `json:"self"`
	} `json:"links"`
	Enabled  bool   `json:"enabled"`
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
	DomainID string `json:"domain_id"`
	Name     string `json:"name"`
}
