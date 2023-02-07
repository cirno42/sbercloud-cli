package availabilityZoneModels

type AvailabilityZoneInfo struct {
	ZoneState AvailabilityZoneState `json:"zoneState"`
	ZoneName  string                `json:"zoneName" header:"zoneName"`
}

type AvailabilityZoneState struct {
	Available bool `json:"available" header:"available"`
}
