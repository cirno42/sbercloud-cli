package ecs

import (
	"sbercloud-cli/api/models/ecsModels"
)

type escCreateRequest struct {
	Server ecsCreateParameters `json:"server"`
}

type ecsCreateParameters struct {
	AvailabilityZone string `json:"availability_zone"`
	Name             string `json:"name"`
	ImageRef         string `json:"imageRef"`
	RootVolume       struct {
		Volumetype string `json:"volumetype"`
	} `json:"root_volume"`
	DataVolumes []struct {
		Volumetype    string `json:"volumetype"`
		Size          int    `json:"size"`
		Multiattach   bool   `json:"multiattach"`
		HwPassthrough bool   `json:"hw:passthrough"`
	} `json:"data_volumes"`
	FlavorRef      string `json:"flavorRef"`
	Vpcid          string `json:"vpcid"`
	SecurityGroups []struct {
		ID string `json:"id"`
	} `json:"security_groups"`
	Nics []struct {
		SubnetID string `json:"subnet_id"`
	} `json:"nics"`
	Publicip struct {
		Eip struct {
			Iptype    string `json:"iptype"`
			Bandwidth struct {
				Size      int    `json:"size"`
				Sharetype string `json:"sharetype"`
			} `json:"bandwidth"`
		} `json:"eip"`
	} `json:"publicip"`
	KeyName    string `json:"key_name"`
	Count      int    `json:"count"`
	ServerTags []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"server_tags"`
	Metadata struct {
		OpSvcUserid string `json:"op_svc_userid"`
		AgencyName  string `json:"agency_name"`
	} `json:"metadata"`
}

func CreateECS(projectID, vpcID, imageRef, name, flavorRef, volumeType, subnetID string) (*ecsModels.ESCJobID, error) {
	/*
		endpoint := fmt.Sprintf("https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/%s/cloudservers", projectID)
		ecsRequest := escCreateRequest{Server: ecsCreateParameters{
			AvailabilityZone: "",
			Name:             name,
			ImageRef:         imageRef,
			RootVolume: struct {
				Volumetype string `json:"volumetype"`
			}{},
			DataVolumes:    nil,
			FlavorRef:      flavorRef,
			Vpcid:          vpcID,
			SecurityGroups: nil,
			Nics:           nil,
			Publicip: struct {
				Eip struct {
					Iptype    string `json:"iptype"`
					Bandwidth struct {
						Size      int    `json:"size"`
						Sharetype string `json:"sharetype"`
					} `json:"bandwidth"`
				} `json:"eip"`
			}{},
			KeyName:    "",
			Count:      0,
			ServerTags: nil,
			Metadata: struct {
				OpSvcUserid string `json:"op_svc_userid"`
				AgencyName  string `json:"agency_name"`
			}{},
		}}
		var createdJobID ecsModels.ESCJobID
		err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, &ecsRequest, &createdJobID, nil)
		return &createdJobID, err
	*/
	return nil, nil
}
