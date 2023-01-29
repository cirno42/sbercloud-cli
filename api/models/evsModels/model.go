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
	AvailabilityZone string          `json:"availability_zone"`
	Bootable         string          `json:"bootable"`
	CreatedAt        string          `json:"created_at"`
	Description      string          `json:"description"`
	Encrypted        bool            `json:"encrypted"`
	ID               string          `json:"id"`
	Links            []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name              string `json:"name"`
	OsVolHostAttrHost string `json:"os-vol-host-attr:host"`
	ReplicationStatus string `json:"replication_status"`
	Multiattach       bool   `json:"multiattach"`
	Size              int    `json:"size"`
	SnapshotID        string `json:"snapshot_id"`
	SourceVolid       string `json:"source_volid"`
	Status            string `json:"status"`
	UpdatedAt         string `json:"updated_at"`
	UserID            string `json:"user_id"`
	VolumeType        string `json:"volume_type"`
}
