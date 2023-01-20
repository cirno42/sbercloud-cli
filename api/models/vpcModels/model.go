package vpcModels

type VpcEntity struct {
	Vpc VpcModel `json:"vpc"`
}

type VpcModel struct {
	Id                  string       `json:"id" header:"id"`
	Name                string       `json:"name" header:"name"`
	Description         string       `json:"description" header:"description"`
	Cidr                string       `json:"cidr" header:"cidr"`
	Status              string       `json:"status" header:"status"`
	Routes              []VpcsRoutes `json:"routes" header:"routes"`
	EnterpriseProjectId string       `json:"enterprise_project_id" header:"enterprise_project_id"`
}

type VpcsRoutes struct {
	Destination string `json:"destination"`
	Nexthop     string `json:"nexthop"`
}
