package ecs

import (
	"fmt"
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type escCreateRequest struct {
	Server ecsCreateParameters `json:"server"`
}

type dataVolume struct {
	Volumetype    string `json:"volumetype"`
	Size          int    `json:"size"`
	Multiattach   bool   `json:"multiattach"`
	HwPassthrough bool   `json:"hw:passthrough"`
}

type secGroupId struct {
	ID string `json:"id"`
}

type nic struct {
	SubnetID string `json:"subnet_id"`
}

type bandwidth struct {
	Size      int    `json:"size"`
	Sharetype string `json:"sharetype"`
}

type eip struct {
	Iptype    string    `json:"iptype"`
	Bandwidth bandwidth `json:"bandwidth"`
}

type publicIp struct {
	Eip eip `json:"eip"`
}

type serverTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type rootVolume struct {
	Volumetype string `json:"volumetype"`
}

type ecsCreateParameters struct {
	AvailabilityZone string       `json:"availability_zone"`
	Name             string       `json:"name"`
	ImageRef         string       `json:"imageRef"`
	RootVolume       rootVolume   `json:"root_volume"`
	DataVolumes      []dataVolume `json:"data_volumes"`
	FlavorRef        string       `json:"flavorRef"`
	Vpcid            string       `json:"vpcid"`
	SecurityGroups   []secGroupId `json:"security_groups"`
	Nics             []nic        `json:"nics"`
	AdminPass        string       `json:"adminPass"`
	//Publicip         publicIp     `json:"publicip"`
	KeyName    string      `json:"key_name"`
	Count      int         `json:"count"`
	ServerTags []serverTag `json:"server_tags"`
}

func CreateECS(projectID, vpcID, imageRef, name, flavorRef, rootVolumeType string,
	subnetIds []string, secGroupIds []string, adminPass string, count int) (*ecsModels.ESCJobID, error) {

	endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers", projectID)
	rv := rootVolume{Volumetype: rootVolumeType}
	subnets := make([]nic, len(subnetIds))
	for i := 0; i < len(subnetIds); i++ {
		subnets[i].SubnetID = subnetIds[i]
	}
	secGroups := make([]secGroupId, len(secGroupIds))
	for i := 0; i < len(secGroupIds); i++ {
		secGroups[i].ID = secGroupIds[i]
	}
	ecsRequest := escCreateRequest{Server: ecsCreateParameters{
		AvailabilityZone: "",
		Name:             name,
		ImageRef:         imageRef,
		RootVolume:       rv,
		FlavorRef:        flavorRef,
		Vpcid:            vpcID,
		SecurityGroups:   secGroups,
		Nics:             subnets,
		KeyName:          "",
		AdminPass:        adminPass,
		Count:            count,
		ServerTags:       nil,
	}}
	var createdJobID ecsModels.ESCJobID
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &ecsRequest, &createdJobID, nil)
	return &createdJobID, err
}
