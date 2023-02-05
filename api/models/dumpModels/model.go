package dumpModels

import (
	"sbercloud-cli/api/models/ecsModels"
	"sbercloud-cli/api/models/eipModels"
	"sbercloud-cli/api/models/evsModels"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/api/models/securityGroupModels"
	"sbercloud-cli/api/models/subnetModels"
	"sbercloud-cli/api/models/vpcModels"
)

type DumpModel struct {
	Vpcs      []vpcModels.VpcModel                     `json:"vpcs"`
	Subnets   []subnetModels.SubnetModel               `json:"subnets"`
	Eips      []eipModels.EipModel                     `json:"eips"`
	Nats      []natModels.NatModel                     `json:"nats"`
	SecGroups []securityGroupModels.SecurityGroupModel `json:"secGroups"`
	SnatRules []natModels.SnatRuleModel                `json:"snatRules"`
	DnatRules []natModels.DnatRuleModel                `json:"dnatRules"`
	ECSs      []ecsModels.ECSModel                     `json:"ECSs"`
	Disks     []evsModels.EvsModel                     `json:"disks"`
	KeyPairs  []ecsModels.Keypair                      `json:"keyPairs"`
}
