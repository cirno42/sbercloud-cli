package ecsModels

import "time"

type ECSNetowrk struct {
	Version            string `json:"version"`
	Addr               string `json:"addr"`
	OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	OSEXTIPSPortID     string `json:"OS-EXT-IPS:port_id"`
	OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
}

type ECSModel struct {
	Fault struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Details string `json:"details"`
		Created string `json:"created"`
	} `json:"fault"`
	Addresses map[string]interface{} `json:"addresses"` //workaround for jmespath query, more info: https://github.com/jmespath/go-jmespath/issues/32
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Flavor    struct {
		Disk  string `json:"disk"`
		Vcpus string `json:"vcpus"`
		RAM   string `json:"ram"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"flavor"`
	AccessIPv4 string `json:"accessIPv4"`
	AccessIPv6 string `json:"accessIPv6"`
	Status     string `json:"status"`
	Image      struct {
		ID string `json:"id"`
	} `json:"image"`
	HostID   string    `json:"hostId"`
	Updated  time.Time `json:"updated"`
	Created  time.Time `json:"created"`
	Metadata struct {
		MeteringImageID          string `json:"metering.image_id"`
		MeteringImagetype        string `json:"metering.imagetype"`
		MeteringResourcespeccode string `json:"metering.resourcespeccode"`
		ImageName                string `json:"image_name"`
		MeteringResourcetype     string `json:"metering.resourcetype"`
		OsBit                    string `json:"os_bit"`
		VpcID                    string `json:"vpc_id"`
		OsType                   string `json:"os_type"`
		ChargingMode             string `json:"charging_mode"`
	} `json:"metadata"`
	Tags                             []interface{} `json:"tags"`
	Description                      string        `json:"description"`
	Locked                           bool          `json:"locked"`
	ConfigDrive                      string        `json:"config_drive"`
	TenantID                         string        `json:"tenant_id"`
	UserID                           string        `json:"user_id"`
	OsExtendedVolumesVolumesAttached []struct {
		Device              string `json:"device"`
		BootIndex           string `json:"bootIndex"`
		ID                  string `json:"id"`
		DeleteOnTermination string `json:"delete_on_termination"`
	} `json:"os-extended-volumes:volumes_attached"`
	OSEXTSTSTaskState              interface{} `json:"OS-EXT-STS:task_state"`
	OSEXTSTSPowerState             int         `json:"OS-EXT-STS:power_state"`
	OSEXTSTSVMState                string      `json:"OS-EXT-STS:vm_state"`
	OSEXTSRVATTRHost               string      `json:"OS-EXT-SRV-ATTR:host"`
	OSEXTSRVATTRInstanceName       string      `json:"OS-EXT-SRV-ATTR:instance_name"`
	OSEXTSRVATTRHypervisorHostname string      `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	OSDCFDiskConfig                string      `json:"OS-DCF:diskConfig"`
	OSEXTAZAvailabilityZone        string      `json:"OS-EXT-AZ:availability_zone"`
	OsSchedulerHints               struct {
	} `json:"os:scheduler_hints"`
	OSEXTSRVATTRRootDeviceName string `json:"OS-EXT-SRV-ATTR:root_device_name"`
	OSEXTSRVATTRRamdiskID      string `json:"OS-EXT-SRV-ATTR:ramdisk_id"`
	OSEXTSRVATTRUserData       string `json:"OS-EXT-SRV-ATTR:user_data"`
	OSSRVUSGLaunchedAt         string `json:"OS-SRV-USG:launched_at"`
	OSEXTSRVATTRKernelID       string `json:"OS-EXT-SRV-ATTR:kernel_id"`
	OSEXTSRVATTRLaunchIndex    int    `json:"OS-EXT-SRV-ATTR:launch_index"`
	HostStatus                 string `json:"host_status"`
	OSEXTSRVATTRReservationID  string `json:"OS-EXT-SRV-ATTR:reservation_id"`
	OSEXTSRVATTRHostname       string `json:"OS-EXT-SRV-ATTR:hostname"`
	SysTags                    []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"sys_tags"`
	SecurityGroups []struct {
		Name string `json:"name"`
	} `json:"security_groups"`
}

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
	RxtxFactor             float64     `json:"rxtx_factor"`
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

type VolumeAttachments struct {
	PciAddress string `json:"pciAddress"`
	VolumeID   string `json:"volumeId"`
	Device     string `json:"device"`
	ServerID   string `json:"serverId"`
	ID         string `json:"id"`
	Size       int    `json:"size"`
	BootIndex  int    `json:"bootIndex"`
	Bus        string `json:"bus"`
}
type AttachedDisks struct {
	AttachableQuantity struct {
		FreeScsi int `json:"free_scsi"`
		FreeBlk  int `json:"free_blk"`
		FreeDisk int `json:"free_disk"`
	} `json:"attachableQuantity"`
	Volumes []VolumeAttachments `json:"volumeAttachments"`
}

type BindPrivateIpResponse struct {
	PortID string `json:"port_id"`
}
