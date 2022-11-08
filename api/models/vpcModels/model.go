package vpcModels

type VpcEntity struct {
	Vpc VpcModel `json:"vpc"`
}

type VpcModel struct {
	Id                  string       `json:"id"`
	Name                string       `json:"name"`
	Description         string       `json:"description"`
	Cidr                string       `json:"cidr"`
	Status              string       `json:"status"`
	Routes              []VpcsRoutes `json:"routes"`
	EnterpriseProjectId string       `json:"enterprise_project_id"`
}

type VpcsRoutes struct {
	Destination string `json:"destination"`
	Nexthop     string `json:"nexthop"`
}
