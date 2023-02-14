package quota

import (
	"fmt"
	"sbercloud-cli/api/ecs"
	"sbercloud-cli/api/endpoints"
	"sbercloud-cli/api/models/quotaModels"
	"sbercloud-cli/internal/handlers/requestMakers"
	"strconv"
)

type getQuotaResponse struct {
	Quotas struct {
		Resources []quotaModels.QuotaModel `json:"resources"`
	} `json:"quotas"`
}

type ecsQuotaResponse struct {
	Absolute struct {
		MaxServerMeta           int `json:"maxServerMeta"`
		MaxPersonality          int `json:"maxPersonality"`
		MaxImageMeta            int `json:"maxImageMeta"`
		MaxPersonalitySize      int `json:"maxPersonalitySize"`
		MaxSecurityGroupRules   int `json:"maxSecurityGroupRules"`
		MaxTotalKeypairs        int `json:"maxTotalKeypairs"`
		TotalRAMUsed            int `json:"totalRAMUsed"`
		TotalInstancesUsed      int `json:"totalInstancesUsed"`
		MaxSecurityGroups       int `json:"maxSecurityGroups"`
		TotalFloatingIpsUsed    int `json:"totalFloatingIpsUsed"`
		MaxTotalCores           int `json:"maxTotalCores"`
		TotalSecurityGroupsUsed int `json:"totalSecurityGroupsUsed"`
		MaxTotalFloatingIps     int `json:"maxTotalFloatingIps"`
		MaxTotalInstances       int `json:"maxTotalInstances"`
		TotalCoresUsed          int `json:"totalCoresUsed"`
		MaxTotalRAMSize         int `json:"maxTotalRAMSize"`
		MaxServerGroups         int `json:"maxServerGroups"`
		MaxServerGroupMembers   int `json:"maxServerGroupMembers"`
		TotalServerGroupsUsed   int `json:"totalServerGroupsUsed"`
	} `json:"absolute"`
}

func GetInfoAboutNetworkQuota(projectID, quotaType string) ([]quotaModels.QuotaModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.VpcEndpoint)+"/v1/%s/quotas", projectID)
	if quotaType != "" {
		endpoint += "?type=" + quotaType
	}
	var resp getQuotaResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	return resp.Quotas.Resources, err
}

func GetInfoAboutServerQuota(projectID string) ([]quotaModels.QuotaModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/limits", projectID)
	var resp ecsQuotaResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &resp, nil)
	if err != nil {
		return nil, err
	}
	servers, err := ecs.GetECSList(projectID, 0, 1000)
	if err != nil {
		return nil, err
	}
	quotas := make([]quotaModels.QuotaModel, 3)
	totalCpus := 0
	totalRam := 0
	for _, server := range servers {
		cpu, _ := strconv.Atoi(server.Flavor.Vcpus)
		ram, _ := strconv.Atoi(server.Flavor.RAM)
		totalCpus += cpu
		totalRam += ram
	}
	quotas[0] = quotaModels.QuotaModel{
		Type:  "ecs",
		Used:  len(servers),
		Quota: resp.Absolute.MaxTotalInstances,
		Min:   0,
	}
	quotas[1] = quotaModels.QuotaModel{
		Type:  "vcpus",
		Used:  totalCpus,
		Quota: resp.Absolute.MaxTotalCores,
		Min:   0,
	}
	quotas[2] = quotaModels.QuotaModel{
		Type:  "ram",
		Used:  totalRam,
		Quota: resp.Absolute.MaxTotalRAMSize,
		Min:   0,
	}
	return quotas, err
}
