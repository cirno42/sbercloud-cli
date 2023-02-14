package quotaModels

type QuotaModel struct {
	Type  string `json:"type" header:"type"`
	Used  int    `json:"used" header:"used"`
	Quota int    `json:"quota" header:"quota"`
	Min   int    `json:"min" header:"min"`
}
