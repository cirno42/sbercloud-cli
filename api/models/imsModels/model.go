package imsModels

type ImageModel struct {
	Schema             string        `json:"schema"`
	MinDisk            int           `json:"min_disk"`
	CreatedAt          string        `json:"created_at" header:"created_at"`
	ImageSourceType    string        `json:"__image_source_type"`
	ContainerFormat    string        `json:"container_format"`
	File               string        `json:"file"`
	UpdatedAt          string        `json:"updated_at" header:"updated_at"`
	Protected          bool          `json:"protected"`
	Checksum           string        `json:"checksum"`
	SupportKvmFpgaType string        `json:"__support_kvm_fpga_type"`
	ID                 string        `json:"id" header:"schema"`
	Isregistered       string        `json:"__isregistered"`
	MinRAM             int           `json:"min_ram"`
	Lazyloading        string        `json:"__lazyloading"`
	Owner              string        `json:"owner"`
	OsType             string        `json:"__os_type"`
	Imagetype          string        `json:"__imagetype"`
	Visibility         string        `json:"visibility"`
	VirtualEnvType     string        `json:"virtual_env_type"`
	Tags               []interface{} `json:"tags"`
	Platform           string        `json:"__platform" header:"platform"`
	Size               int           `json:"size" header:"Size"`
	OsBit              string        `json:"__os_bit"`
	OsVersion          string        `json:"__os_version" header:"os_version"`
	Name               string        `json:"name" header:"name"`
	Self               string        `json:"self"`
	DiskFormat         string        `json:"disk_format"`
	VirtualSize        interface{}   `json:"virtual_size"`
	HwFirmwareType     string        `json:"hw_firmware_type"`
	Status             string        `json:"status"`
	SupportFcInject    string        `json:"__support_fc_inject"`
	SupportXenGpuType  string        `json:"__support_xen_gpu_type"`
}
