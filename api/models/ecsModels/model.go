package ecsModels

import "time"

type FlavorModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Vcpus string `json:"vcpus"`
	RAM   int    `json:"ram"`
	Disk  string `json:"disk"`
	Swap  string `json:"swap"`
	Links []struct {
		Rel  string      `json:"rel"`
		Href string      `json:"href"`
		Type interface{} `json:"type"`
	} `json:"links"`
	OSFLVEXTDATAEphemeral  int         `json:"OS-FLV-EXT-DATA:ephemeral"`
	RxtxFactor             int         `json:"rxtx_factor"`
	OSFLVDISABLEDDisabled  bool        `json:"OS-FLV-DISABLED:disabled"`
	RxtxQuota              interface{} `json:"rxtx_quota"`
	RxtxCap                interface{} `json:"rxtx_cap"`
	OsFlavorAccessIsPublic bool        `json:"os-flavor-access:is_public"`
	OsExtraSpecs           struct {
		EcsVirtualizationEnvTypes string `json:"ecs:virtualization_env_types"`
		EcsGeneration             string `json:"ecs:generation"`
		EcsPerformancetype        string `json:"ecs:performancetype"`
		ResourceType              string `json:"resource_type"`
	} `json:"os_extra_specs"`
}

type ESCJobID struct {
	JobID string `json:"job_id"`
}

type ECSJob struct {
	Status   string `json:"status"`
	Entities struct {
		SubJobsTotal int `json:"sub_jobs_total"`
		SubJobs      []struct {
			Status   string `json:"status"`
			Entities struct {
				ServerID string `json:"server_id"`
			} `json:"entities"`
			JobID      string    `json:"job_id"`
			JobType    string    `json:"job_type"`
			BeginTime  time.Time `json:"begin_time"`
			EndTime    time.Time `json:"end_time"`
			ErrorCode  string    `json:"error_code"`
			FailReason string    `json:"fail_reason"`
		} `json:"sub_jobs"`
	} `json:"entities"`
	JobID      string    `json:"job_id"`
	JobType    string    `json:"job_type"`
	BeginTime  time.Time `json:"begin_time"`
	EndTime    time.Time `json:"end_time"`
	ErrorCode  string    `json:"error_code"`
	FailReason string    `json:"fail_reason"`
}
