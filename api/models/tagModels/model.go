package tagModels

type Tag struct {
	Key   string `json:"key" header:"key"`
	Value string `json:"value" header:"value"`
}

type Resource struct {
	ResourceID     string `json:"resource_id" header:"resource_id"`
	ResourceName   string `json:"resource_name" header:"resource_name"`
	ResourceDetail string `json:"resource_detail" header:"resource_detail"`
	Tags           []Tag  `json:"tags"`
}

type ResourceInfo struct {
	Resources  []Resource `json:"resources"`
	TotalCount int        `json:"total_count"`
}
