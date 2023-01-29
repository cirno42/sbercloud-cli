package ecs

import (
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type EcsNic struct {
	InterfaceAttachments []struct {
		PortState string `json:"port_state"`
		FixedIps  []struct {
			SubnetID  string `json:"subnet_id"`
			IPAddress string `json:"ip_address"`
		} `json:"fixed_ips"`
		NetID   string `json:"net_id"`
		PortID  string `json:"port_id"`
		MacAddr string `json:"mac_addr"`
	} `json:"interfaceAttachments"`
}

func GetEcsNics(projectID, serverID string) (EcsNic, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/%s/os-interface", projectID, serverID)
	var nics EcsNic
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &nics, nil)
	return nics, err
}
