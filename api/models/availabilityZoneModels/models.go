package availabilityZoneModels

type AvailabilityZoneInfo struct {
	ZoneState struct {
		Available bool `json:"available"`
	} `json:"zoneState"`
	ZoneName string `json:"zoneName"`
}
