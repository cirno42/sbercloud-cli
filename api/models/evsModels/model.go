package evsModels

type EvsAttachment struct {
	ServerId     string `json:"server_id"`
	AttachmentId string `json:"attachment_id"`
	AttachedAt   string `json:"attached_at"`
	VolumeId     string `json:"volume_id"`
	Device       string `json:"device"`
}

type EvsModel struct {
	Attachments      []EvsAttachment `json:"attachments"`
	AvailabilityZone string          `json:"availability_zone" header:"availability_zone"`
	Bootable         string          `json:"bootable" header:"bootable"`
	CreatedAt        string          `json:"created_at" header:"created_at"`
	Description      string          `json:"description" header:"description"`
	Encrypted        bool            `json:"encrypted" header:"encrypted"`
	ID               string          `json:"id" header:"id"`
	Links            []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name              string `json:"name" header:"name"`
	OsVolHostAttrHost string `json:"os-vol-host-attr:host"`
	ReplicationStatus string `json:"replication_status"`
	Multiattach       bool   `json:"multiattach"`
	Size              int    `json:"size" header:"size"`
	SnapshotID        string `json:"snapshot_id"`
	SourceVolid       string `json:"source_volid"`
	Status            string `json:"status" header:"status"`
	UpdatedAt         string `json:"updated_at" header:"updated_at"`
	UserID            string `json:"user_id"`
	VolumeType        string `json:"volume_type" header:"volume_type"`
}
