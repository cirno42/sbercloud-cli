package endpoints

type ServicesEndpoints int

const (
	VpcEndpoint ServicesEndpoints = iota
	EscEndpoint
	ImsEndpoint
	NatEndpoint
)

var ruMoscow1RegionEndpoints map[ServicesEndpoints]string = map[ServicesEndpoints]string{
	VpcEndpoint: "https://vpc.ru-moscow-1.hc.sbercloud.ru",
	EscEndpoint: "https://ecs.ru-moscow-1.hc.sbercloud.ru",
	ImsEndpoint: "https://ims.ru-moscow-1.hc.sbercloud.ru",
	NatEndpoint: "https://nat.ru-moscow-1.hc.sbercloud.ru",
}

func GetEndpointAddress(endpoint ServicesEndpoints) string {
	return ruMoscow1RegionEndpoints[endpoint]
}
