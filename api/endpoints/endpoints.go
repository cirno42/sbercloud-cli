package endpoints

import (
	"os"
)

type ServicesEndpoints int

const (
	VpcEndpoint ServicesEndpoints = iota
	EscEndpoint
	ImsEndpoint
	NatEndpoint
	EvsEndpoint
	IamEndpoint
)

var ruMoscow1RegionEndpoints map[ServicesEndpoints]string = map[ServicesEndpoints]string{
	VpcEndpoint: "https://vpc.ru-moscow-1.hc.sbercloud.ru",
	EscEndpoint: "https://ecs.ru-moscow-1.hc.sbercloud.ru",
	ImsEndpoint: "https://ims.ru-moscow-1.hc.sbercloud.ru",
	NatEndpoint: "https://nat.ru-moscow-1.hc.sbercloud.ru",
	EvsEndpoint: "https://evs.ru-moscow-1.hc.sbercloud.ru",
	IamEndpoint: "https://iam.ru-moscow-1.hc.sbercloud.ru",
}

func GetEndpointAddress(endpoint ServicesEndpoints) string {
	region := os.Getenv("REGION")
	if region == "RU-Moscow" {
		return ruMoscow1RegionEndpoints[endpoint]
	}
	return ""
}
