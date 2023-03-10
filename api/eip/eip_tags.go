package eip

import (
	"errors"
	"fmt"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/tagModels"
	"sbercloud-cli/internal/handlers/requestMakers"
)

type tagActionRequest struct {
	Tags   []tagModels.Tag `json:"tags"`
	Action string          `json:"action"`
}

type tagListResp struct {
	Tags []tagModels.Tag `json:"tags"`
}

type tagFilterParams struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type tagFilterRequest struct {
	Action string            `json:"action"`
	Tags   []tagFilterParams `json:"tags"`
}

func CreateEipTag(projectID, eipId string, tagKeys, tagValues []string) error {
	if len(tagKeys) != len(tagValues) {
		return errors.New("{\"error\" : \"Amount of keys and values must be equal\"}")
	}
	err := doActionWithTag(projectID, eipId, "create", tagKeys, tagValues)
	return err
}

func DeleteEipTag(projectID, eipId string, tagKeys, tagValues []string) error {
	if len(tagKeys) != len(tagValues) {
		return errors.New("{\"error\" : \"Amount of keys and values must be equal\"}")
	}
	err := doActionWithTag(projectID, eipId, "delete", tagKeys, tagValues)
	return err
}

func GetEipTags(projectID, eipId string) ([]tagModels.Tag, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v2.0/%s/publicips/%s/tags", projectID, eipId)
	var tags tagListResp
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &tags, nil)
	return tags.Tags, err
}

func FilterEipByTags(projectID string, tagKeys, tagValues []string) ([]tagModels.Resource, error) {
	res, err := filterEipByTags(projectID, "filter", tagKeys, tagValues)
	if err != nil {
		return nil, err
	}
	return res.Resources, err
}

func CountEipByTags(projectID string, tagKeys, tagValues []string) (tagModels.ResourceInfo, error) {
	res, err := filterEipByTags(projectID, "count", tagKeys, tagValues)
	if err != nil {
		return res, err
	}
	return res, err
}

func filterEipByTags(projectID, action string, tagKeys, tagValues []string) (tagModels.ResourceInfo, error) {
	if len(tagKeys) != len(tagValues) {
		return tagModels.ResourceInfo{}, errors.New("{\"error\" : \"Amount of keys and values must be equal\"}")
	}
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v2.0/%s/publicips/resource_instances/action", projectID)
	tags := make([]tagFilterParams, len(tagKeys))
	for i, key := range tagKeys {
		tags[i].Key = key
		tags[i].Values = make([]string, 1)
		tags[i].Values[0] = tagValues[i]

	}
	req := tagFilterRequest{
		Action: action,
		Tags:   tags,
	}
	var resp tagModels.ResourceInfo
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &req, &resp, nil)
	return resp, err
}

func doActionWithTag(projectID, eipId, action string, tagKeys, tagValues []string) error {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v2.0/%s/publicips/%s/tags/action", projectID, eipId)
	tags := make([]tagModels.Tag, len(tagKeys))
	for i, key := range tagKeys {
		tags[i].Key = key
	}
	for i, value := range tagValues {
		tags[i].Value = value
	}
	req := tagActionRequest{
		Tags:   tags,
		Action: action,
	}
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_POST, &req, nil, nil)
	return err
}
