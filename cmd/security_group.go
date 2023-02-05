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
	Short: "Commands to interact with security groups",
	Long:  `Commands to interact with security groups`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var sgCreateName string
var sgCreateVpcID string
var sgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create security group",
	Long:  `Create security group`,
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
	Short: "Get info about security group",
	Long:  `Get info about security group`,
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
	Short: "Get list of security groups",
	Long:  `Get list of security groups`,
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
	Short: "Delete security group",
	Long:  `Delete security group`,
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
	Short: "Add rule to security group",
	Long:  `Add rule to security group`,
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
	Short: "Get info about security group rule",
	Long:  `Get info about security group rule`,
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
	Short: "Get list of security rules",
	Long:  `Get list of security rules`,
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
	Short: "Delete security rule",
	Long:  `Delete security rule`,
	Run: func(cmd *cobra.Command, args []string) {

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

	sgCreateCmd.Flags().StringVarP(&sgCreateName, "name", "n", "", "Specifies Security Group name")
	sgCreateCmd.Flags().StringVarP(&sgCreateVpcID, "vpc", "v", "", "Specifies VPC ID")

	sgGetInfoCmd.Flags().StringVarP(&sgGetInfoSGID, "id", "i", "", "Specifies Security Group ID")
	sgGetInfoCmd.Flags().StringVarP(&sgGetInfoName, "name", "n", "", "Specifies Security Group name")

	sgGetListCmd.Flags().IntVarP(&sgGetListLimit, "limit", "l", 0, "Specifies the number of records that will be returned on each page. The value is from 0 to intmax.")
	sgGetListCmd.Flags().StringVarP(&sgGetListMarker, "marker", "m", "", "Specifies a resource ID for pagination query, indicating that the query starts from the next record of the specified resource ID.")
	sgGetListCmd.Flags().StringVarP(&sgGetListVpcID, "vpc_id", "v", "", "Specifies VPC ID")

	sgDeleteCmd.Flags().StringVarP(&sgDeleteName, "name", "n", "", "Specifies Security Group name")
	sgDeleteCmd.Flags().StringVarP(&sgDeleteSGID, "id", "i", "", "Specifies Security Group ID")

	sgCmdAddRule.Flags().StringVar(&sgAddRuleSGID, "sg-id", "", "Specifies Security Group ID")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleDescription, "desc", "", "Provides supplementary information about the security group rule.")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleDirection, "direction", "", "The value can be: egress or ingress")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleEtherType, "ethertype", "", "Specifies the IP protocol version. The value can be IPv4 or IPv6. If you do not set this parameter, IPv4 is used by default.")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleProtocol, "protocol", "", "Specifies the protocol type. The value can be icmpv6, tcp, udp, icmp or an IP protocol number (0 to 255). If the parameter is left blank, all protocols are supported. When the protocol is icmpv6, IP protocol version should be IPv6. When the protocol is icmp, IP protocol version should be IPv4.")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleRemoteIpPrefix, "remote_ip_prefix", "", "Specifies the remote IP address. If the access control direction is set to egress, the parameter specifies the source IP address. If the access control direction is set to ingress, the parameter specifies the destination IP address. The value can be in the CIDR format or IP addresses. The parameter is exclusive with parameter remote_group_id.")
	sgCmdAddRule.Flags().StringVar(&sgAddRuleRemoteGroupId, "remote_group_id", "", "Specifies the ID of the peer security group. The value is exclusive with parameter remote_ip_prefix.")
	sgCmdAddRule.Flags().IntVar(&sgAddRulePortRangeMin, "port_min", -1, "Specifies the start port number. The value ranges from 1 to 65535.The value cannot be greater than the port_range_max value. An empty value indicates all ports. If the protocol is icmp, the value range is shown in ICMP-Port Range Relationship Table.")
	sgCmdAddRule.Flags().IntVar(&sgAddRulePortRangeMax, "port_max", -1, "Specifies the end port number. The value ranges from 1 to 65535. If the protocol is not icmp, the value cannot be smaller than the port_range_min value. An empty value indicates all ports. If the protocol is icmp, the value range is shown in ICMP-Port Range Relationship Table.")

	sgRuleGetInfoCmd.Flags().StringVarP(&sgRuleGetInfoRuleID, "id", "i", "", "Specifies the Security Group ID")

	sgRuleListCmd.Flags().StringVarP(&sgRuleListSGID, "id", "i", "", "Specifies the Security Group ID")
	sgRuleListCmd.Flags().StringVarP(&sgRuleListSGName, "name", "n", "", "Specifies the Security Group name")
	sgRuleListCmd.Flags().StringVarP(&sgRuleListMarker, "marker", "m", "", "Specifies a resource ID for pagination query, indicating that the query starts from the next record of the specified resource ID.")
	sgRuleListCmd.Flags().IntVarP(&sgRuleListLimit, "limit", "l", 0, "Specifies the number of records that will be returned on each page. The value is from 0 to intmax.")

	sgRuleDeleteCmd.Flags().StringVarP(&sgRuleDeleteRuleID, "id", "i", "", "Specifies the Security Group ID")
}
