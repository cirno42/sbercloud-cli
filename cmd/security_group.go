package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/securityGroupModels"
	"sbercloud-cli/api/securityGroup"
	"sbercloud-cli/internal/beautyfulPrints"
)

var sgCmd = &cobra.Command{
	Use:   "sg",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eip called")
	},
}

var sgCreateName string
var sgCreateVpcID string
var sgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		sg, err := securityGroup.CreateSecurityGroup(ProjectID, sgCreateName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sg, jmesPathQuery)
		}
	},
}

var sgGetInfoSGID string
var sgGetInfoName string
var sgGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var sg *securityGroupModels.SecurityGroupModel
		if sgGetInfoSGID != "" {
			sg, err = securityGroup.GetInfoAboutSecurityGroup(ProjectID, sgGetInfoSGID)
		} else if sgGetInfoName != "" {
			sg, err = securityGroup.GetInfoAboutSecurityGroupByName(ProjectID, sgGetInfoName)
		} else {
			beautyfulPrints.PrintError(errors.New("ID and name are both not specified"))
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sg, jmesPathQuery)
		}
	},
}

var sgGetListLimit int
var sgGetListMarker string
var sgGetListVpcID string
var sgGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sgs, err := securityGroup.GetSecurityGroupsList(ProjectID, sgGetListLimit, sgGetListMarker, sgGetListVpcID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sgs, jmesPathQuery)
		}
	},
}

var sgDeleteSGID string
var sgDeleteName string
var sgDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if sgDeleteSGID != "" {
			err = securityGroup.DeleteSecurityGroup(ProjectID, sgDeleteSGID)
		} else if sgDeleteName != "" {
			err = securityGroup.DeleteSecurityGroupByName(ProjectID, sgDeleteName)
		} else {
			err = errors.New("ID and name are both not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("Deleted successfully")
		}
	},
}

var sgAddRuleSGID string
var sgAddRuleDescription string
var sgAddRuleDirection string
var sgAddRuleEtherType string
var sgAddRuleProtocol string
var sgAddRuleRemoteIpPrefix string
var sgAddRuleRemoteGroupId string
var sgAddRulePortRangeMin int
var sgAddRulePortRangeMax int
var sgCmdAddRule = &cobra.Command{
	Use:   "add-rule",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		sgRule, err := securityGroup.CreateSecurityGroupRule(ProjectID, sgAddRuleSGID, sgAddRuleDescription, sgAddRuleDirection, sgAddRuleEtherType,
			sgAddRuleProtocol, sgAddRulePortRangeMin, sgAddRulePortRangeMax, sgAddRuleRemoteIpPrefix, sgAddRuleRemoteGroupId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sgRule, jmesPathQuery)
		}
	},
}

var sgRuleGetInfoRuleID string
var sgRuleGetInfoCmd = &cobra.Command{
	Use:   "rule-info",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		sgRule, err := securityGroup.GetInfoAboutSecurityGroupRule(ProjectID, sgRuleGetInfoRuleID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sgRule, jmesPathQuery)
		}
	},
}

var sgRuleListMarker string
var sgRuleListLimit int
var sgRuleListSGID string
var sgRuleListSGName string
var sgRuleListCmd = &cobra.Command{
	Use:   "rule-list",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var rulesList []securityGroupModels.SecurityGroupRuleModel
		if sgRuleListSGID != "" {
			rulesList, err = securityGroup.GetSecurityGroupRulesList(ProjectID, sgRuleListLimit, sgRuleListMarker, sgRuleListSGID)
		} else if sgRuleListSGName != "" {
			rulesList, err = securityGroup.GetSecurityGroupRulesListBySGName(ProjectID, sgRuleListLimit, sgRuleListMarker, sgRuleListSGName)
		} else {
			err = errors.New("name and id of security group are both not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(rulesList, jmesPathQuery)
		}
	},
}

var sgRuleDeleteRuleID string
var sgRuleDeleteCmd = &cobra.Command{
	Use:   "rule-delete",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eip called")
	},
}

func init() {
	RootCmd.AddCommand(sgCmd)
	sgCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	sgCmd.AddCommand(sgCreateCmd)
	sgCmd.AddCommand(sgGetInfoCmd)
	sgCmd.AddCommand(sgGetListCmd)
	sgCmd.AddCommand(sgDeleteCmd)
	sgCmd.AddCommand(sgCmdAddRule)
	sgCmd.AddCommand(sgRuleGetInfoCmd)
	sgCmd.AddCommand(sgRuleListCmd)
	sgCmd.AddCommand(sgRuleDeleteCmd)

	sgCreateCmd.Flags().StringVarP(&sgCreateName, "name", "n", "", "")
	sgCreateCmd.Flags().StringVarP(&sgCreateVpcID, "vpc", "v", "", "")

	sgGetInfoCmd.Flags().StringVarP(&sgGetInfoSGID, "id", "i", "", "")
	sgGetInfoCmd.Flags().StringVarP(&sgGetInfoName, "name", "n", "", "")

	sgGetListCmd.Flags().IntVarP(&sgGetListLimit, "limit", "l", 0, "")
	sgGetListCmd.Flags().StringVarP(&sgGetListMarker, "marker", "m", "", "")
	sgGetListCmd.Flags().StringVarP(&sgGetListVpcID, "vpc_id", "v", "", "")

	sgDeleteCmd.Flags().StringVarP(&sgDeleteName, "name", "n", "", "")
	sgDeleteCmd.Flags().StringVarP(&sgDeleteSGID, "id", "i", "", "")

	sgCmdAddRule.Flags().StringVar(&sgAddRuleSGID, "sg-id", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleDescription, "desc", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleDirection, "direction", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleEtherType, "ethertype", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleProtocol, "protocol", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleRemoteIpPrefix, "remote_ip_prefix", "", "")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleRemoteGroupId, "remote_group_id", "", "")
	sgCmdAddRule.Flags().IntVar(&sgAddRulePortRangeMin, "port_min", 1, "")
	sgCmdAddRule.Flags().IntVar(&sgAddRulePortRangeMax, "port_max", 65535, "")

	sgRuleGetInfoCmd.Flags().StringVarP(&sgRuleGetInfoRuleID, "id", "i", "", "")

	sgRuleListCmd.Flags().StringVarP(&sgRuleListSGID, "id", "i", "", "")
	sgRuleListCmd.Flags().StringVarP(&sgRuleListSGName, "name", "n", "", "")
	sgRuleListCmd.Flags().StringVarP(&sgRuleListMarker, "marker", "m", "", "")
	sgRuleListCmd.Flags().IntVarP(&sgRuleListLimit, "limit", "l", 0, "")

	sgRuleDeleteCmd.Flags().StringVarP(&sgRuleDeleteRuleID, "id", "i", "", "")
}
