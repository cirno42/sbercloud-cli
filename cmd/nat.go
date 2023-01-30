package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/api/nat"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var natCmd = &cobra.Command{
	Use:   "nat",
	Short: "Command to interact with NAT",
	Long:  `Command to interact with NAT`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var natCreateName string
var natCreateDesc string
var natCreateRouterID string
var natCreateInternalNetworkID string
var natCreateSpec string
var natCreateEnterpriseProjectID string
var natCreateRouterName string
var natCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create NAT",
	Long:  `Create NAT`,
	Run: func(cmd *cobra.Command, args []string) {
		var createdNat *natModels.NatModel
		var err error
		var vpcID string
		if natCreateRouterID != "" {
			vpcID = natCreateRouterID
		} else if natCreateRouterName != "" {
			vpc, err := vpcs.GetVpcByName(ProjectID, natCreateRouterName)
			if err != nil {
				vpcID = vpc.Id
			}
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			createdNat, err = nat.CreateNAT(ProjectID, natCreateName, natCreateDesc, vpcID, natCreateInternalNetworkID, natCreateSpec, natCreateEnterpriseProjectID)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(createdNat, jmesPathQuery)
			}
		}

	},
}

var natUpdateId string
var natUpdateName string
var natUpdateDesc string
var natUpdateSpec string
var natUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update NAT",
	Long:  `Update NAT`,
	Run: func(cmd *cobra.Command, args []string) {
		updatedNat, err := nat.UpdateNAT(ProjectID, natUpdateId, natUpdateName, natUpdateDesc, natUpdateSpec)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(updatedNat, jmesPathQuery)
		}
	},
}

var natListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get NAT list",
	Long:  `Get NAT list`,
	Run: func(cmd *cobra.Command, args []string) {
		nats, err := nat.GetNatList(ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(nats, jmesPathQuery)
		}
	},
}

var natGetInfoNatID string
var natGetInfoNatName string
var natGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about NAT",
	Long:  `Get info about NAT`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var natEntity *natModels.NatModel
		if natGetInfoNatID != "" {
			natEntity, err = nat.GetInfoAboutNat(ProjectID, natGetInfoNatID)
		} else if natGetInfoNatName != "" {
			natEntity, err = nat.GetNatByName(ProjectID, natGetInfoNatName)
		} else {
			err = errors.New("NAT name and id both are not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(natEntity, jmesPathQuery)
		}
	},
}

var natDeleteNatID string
var natDeleteNatName string
var natDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete NAT",
	Long:  `Delete NAT`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if natDeleteNatID != "" {
			err = nat.DeleteNat(ProjectID, natDeleteNatID)
		} else if natDeleteNatName != "" {
			err = nat.DeleteNatByName(ProjectID, natDeleteNatName)
		} else {
			err = errors.New("NAT name and id both are not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		}

	},
}

var createSNATRuleNatID string
var createSNATRuleVpcID string
var createSNATRuleEipID string
var createSNATRuleDescription string
var createSNATRuleSourceType int
var natCreateSNATRuleCmd = &cobra.Command{
	Use:   "add-snat-rule",
	Short: "Create SNAT Rule",
	Long:  `Create SNAT Rule`,
	Run: func(cmd *cobra.Command, args []string) {
		rule, err := nat.CreateSNATRule(ProjectID, createSNATRuleNatID, createSNATRuleVpcID, createSNATRuleEipID, createSNATRuleDescription, createSNATRuleSourceType)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(rule, jmesPathQuery)
		}
	},
}

var listSNATRulesGatewayId string
var listSNATRulesIpAddress string
var listSNATRulesLimit int
var natListSNATRuleCmd = &cobra.Command{
	Use:   "list-snat-rule",
	Short: "List SNAT Rules",
	Long:  `List SNAT Rules`,
	Run: func(cmd *cobra.Command, args []string) {
		rules, err := nat.ListSNATRules(ProjectID, listSNATRulesGatewayId, listSNATRulesIpAddress, listSNATRulesLimit)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(rules, jmesPathQuery)
		}
	},
}

var getSNATRuleId string
var natGetInfoSNATRuleCmd = &cobra.Command{
	Use:   "get-snat-rule",
	Short: "This command is used to query details about a specified SNAT rule.",
	Long:  `This command is used to query details about a specified SNAT rule.`,
	Run: func(cmd *cobra.Command, args []string) {
		rule, err := nat.GetSNATRule(ProjectID, getSNATRuleId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(rule, jmesPathQuery)
		}
	},
}

var deleteSNATRuleId string
var deleteRuleNATId string
var natDeleteSNATRuleCmd = &cobra.Command{
	Use:   "delete-snat-rule",
	Short: "This command is used to delete an SNAT rule.",
	Long:  `This command is used to delete an SNAT rule.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := nat.DeleteSNATRule(ProjectID, deleteRuleNATId, deleteSNATRuleId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var updateSNATRuleId string
var updateRuleNATId string
var updateRulePublicIP string
var updateRuleDesc string
var natUpdateSNATRuleCmd = &cobra.Command{
	Use:   "update-snat-rule",
	Short: "This command is used to update an SNAT rule..",
	Long:  `This command is used to update an SNAT rule..`,
	Run: func(cmd *cobra.Command, args []string) {
		rule, err := nat.UpdateSNATRule(ProjectID, updateSNATRuleId, updateRuleNATId, updateRulePublicIP, updateRuleDesc)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(rule, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(natCmd)
	natCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	natCmd.AddCommand(natCreateCmd)
	natCmd.AddCommand(natListCmd)
	natCmd.AddCommand(natGetInfoCmd)
	natCmd.AddCommand(natDeleteCmd)
	natCmd.AddCommand(natUpdateCmd)
	natCmd.AddCommand(natCreateSNATRuleCmd)
	natCmd.AddCommand(natListSNATRuleCmd)
	natCmd.AddCommand(natGetInfoSNATRuleCmd)
	natCmd.AddCommand(natDeleteSNATRuleCmd)
	natCmd.AddCommand(natUpdateSNATRuleCmd)

	natCreateCmd.Flags().StringVarP(&natCreateName, "name", "n", "", "Specifies the NAT gateway name. The name can contain only digits, letters, underscores (_), and hyphens (-).")
	natCreateCmd.Flags().StringVarP(&natCreateDesc, "description", "d", "", "Provides supplementary information about the NAT gateway.")
	natCreateCmd.Flags().StringVarP(&natCreateRouterID, "router-id", "i", "", "Specifies the VPC ID.")
	natCreateCmd.Flags().StringVar(&natCreateInternalNetworkID, "network-id", "", "Specifies the network ID of the downstream interface (the next hop of the DVR) of the NAT gateway. ")
	natCreateCmd.Flags().StringVarP(&natCreateSpec, "spec", "s", "1", "Specifies the NAT gateway type. The value can be: 1: small type, which supports up to 10,000 SNAT connections. 2: medium type, which supports up to 50,000 SNAT connections. 3: large type, which supports up to 200,000 SNAT connections. 4: extra-large type, which supports up to 1,000,000 SNAT connections.")
	natCreateCmd.Flags().StringVarP(&natCreateEnterpriseProjectID, "ent-project-id", "p", "", "Specifies the enterprise project ID. When creating a NAT gateway, associate an enterprise project ID with the NAT gateway. The value 0 indicates the default enterprise project.")
	natCreateCmd.Flags().StringVarP(&natCreateRouterName, "router-name", "r", "", "Specifies the VPC name.")

	natUpdateCmd.Flags().StringVarP(&natUpdateId, "id", "i", "", "Specifies the NAT gateway ID")
	natUpdateCmd.Flags().StringVarP(&natUpdateName, "name", "n", "", "Specifies the NAT gateway name. The name can contain only digits, letters, underscores (_), and hyphens (-).")
	natUpdateCmd.Flags().StringVarP(&natUpdateDesc, "description", "d", "", "Provides supplementary information about the NAT gateway.")
	natUpdateCmd.Flags().StringVarP(&natUpdateSpec, "spec", "s", "1", "Specifies the NAT gateway type. The value can be: 1: small type, which supports up to 10,000 SNAT connections. 2: medium type, which supports up to 50,000 SNAT connections. 3: large type, which supports up to 200,000 SNAT connections. 4: extra-large type, which supports up to 1,000,000 SNAT connections.")

	natGetInfoCmd.Flags().StringVarP(&natGetInfoNatID, "id", "i", "", "Specifies the NAT ID.")
	natGetInfoCmd.Flags().StringVarP(&natGetInfoNatName, "name", "n", "", "Specifies the NAT name.")

	natDeleteCmd.Flags().StringVarP(&natDeleteNatID, "id", "i", "", "Specifies the NAT ID.")
	natDeleteCmd.Flags().StringVarP(&natDeleteNatName, "name", "n", "", "Specifies the NAT name.")

	natCreateSNATRuleCmd.Flags().StringVarP(&createSNATRuleNatID, "nat-id", "n", "", "Specifies the NAT ID")
	natCreateSNATRuleCmd.Flags().StringVarP(&createSNATRuleVpcID, "subnet-id", "i", "", "Specifies the VPC ID")
	natCreateSNATRuleCmd.Flags().StringVarP(&createSNATRuleEipID, "eip-id", "e", "", "Specifies the EIP ID")
	natCreateSNATRuleCmd.Flags().StringVarP(&createSNATRuleDescription, "desc", "d", "", "Provides supplementary information about the NAT gateway")
	natCreateSNATRuleCmd.Flags().IntVarP(&createSNATRuleSourceType, "source-type", "s", 0, "0: Either network_id or cidr can be specified in a VPC. 1: Only cidr can be specified over a Direct Connect connection.")

	natListSNATRuleCmd.Flags().StringVarP(&listSNATRulesGatewayId, "nat-id", "n", "", "Specifies the NAT gateway ID.")
	natListSNATRuleCmd.Flags().StringVarP(&listSNATRulesIpAddress, "ip-addr", "a", "", "Specifies the EIP.")
	natListSNATRuleCmd.Flags().IntVarP(&listSNATRulesLimit, "limit", "l", 0, "Specifies the number of records displayed on each page.")

	natGetInfoSNATRuleCmd.Flags().StringVarP(&getSNATRuleId, "rule-id", "i", "", "Specifies the SNAT rule ID.")

	natDeleteSNATRuleCmd.Flags().StringVarP(&deleteRuleNATId, "nat-id", "n", "", "Specifies the NAT gateway ID.")
	natDeleteSNATRuleCmd.Flags().StringVarP(&deleteSNATRuleId, "rule-id", "r", "", "Specifies the SNAT rule ID.")

	natUpdateSNATRuleCmd.Flags().StringVarP(&updateSNATRuleId, "rule-id", "r", "", "Specifies the SNAT rule ID.")
	natUpdateSNATRuleCmd.Flags().StringVarP(&updateRuleNATId, "nat-id", "n", "", "Specifies the NAT gateway ID.")
	natUpdateSNATRuleCmd.Flags().StringVarP(&updateRulePublicIP, "ip", "i", "", "Specifies the EIP. Multiple EIPs are separated using commas (,).")
	natUpdateSNATRuleCmd.Flags().StringVarP(&updateRuleDesc, "desc", "d", "", "Provides supplementary information about the SNAT rule.")

}
