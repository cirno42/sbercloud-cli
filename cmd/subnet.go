package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/subnetModels"
	"sbercloud-cli/api/subnets"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
	"sbercloud-cli/internal/utils/subnetUtils"
	"sbercloud-cli/internal/utils/vpcUtils"
)

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Commands to interact with subnet",
	Long:  `Commands to interact with subnet`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var subnetCreateName string
var subnetCreateDescription string
var subnetCreateCIDR string
var subnetCreateGatewayIP string
var subnetCreateIPv6Enable bool
var subnetCreateDHCPEnable bool
var subnetCreatePrimaryDNS string
var subnetCreateSecondaryDNS string
var subnetCreateAvailabilityZones string
var subnetCreateVpcName string
var subnetCreateVpcId string

var subnetCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create subnet",
	Long:  `Create subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		var vpcId string
		if subnetCreateVpcId != "" {
			vpcId = subnetCreateVpcId
		} else if subnetCreateVpcName != "" {
			vpc, err := vpcs.GetVpcByName(ProjectID, subnetCreateVpcName)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())
				return
			}
			vpcId = vpc.Id
		} else {
			fmt.Println("ERROR: subnet ID or name is not specified")
			return
		}
		sn, err := subnets.CreateSubnet(ProjectID, subnetCreateName, subnetCreateDescription, subnetCreateCIDR, subnetCreateGatewayIP,
			subnetCreateIPv6Enable, subnetCreateDHCPEnable, subnetCreatePrimaryDNS, subnetCreateSecondaryDNS, nil, subnetCreateAvailabilityZones, vpcId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sn, jmesPathQuery)
		}
	},
}

var subnetListLimit int
var subnetListMarker string
var subnetListVpcID string
var subnetListVpcName string
var subnetGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of subnets",
	Long:  `Get list of subnets`,
	Run: func(cmd *cobra.Command, args []string) {
		subnetId, err := vpcUtils.GetVpcId(subnetListVpcID, subnetListVpcName, ProjectID)
		if err != nil && subnetListVpcName != "" {
			beautyfulPrints.PrintError(err)
			return
		}
		subnets, err := subnets.GetSubnetsList(ProjectID, subnetListLimit, subnetListMarker, subnetId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(subnets, jmesPathQuery)
		}
	},
}

var subnetInfoSubnetID string
var subnetInfoSubnetName string
var subnetGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about subnet",
	Long:  `Get info about subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		var subnet *subnetModels.SubnetModel
		var err error
		if subnetInfoSubnetID != "" {
			subnet, err = subnets.GetInfoAboutSubnet(ProjectID, subnetInfoSubnetID)
		} else {
			subnet, err = subnets.GetSubnetByName(ProjectID, subnetInfoSubnetName)
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(subnet, jmesPathQuery)
		}
	},
}

var subnetUpdateName string
var subnetUpdateDescription string
var subnetUpdateIPv6Enable bool
var subnetUpdateDHCPEnable bool
var subnetUpdatePrimaryDNS string
var subnetUpdateSecondaryDNS string
var subnetUpdateVpcId string
var subnetUpdateSubnetId string
var subnetUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update subnet",
	Long:  `Update subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		sn, err := subnets.UpdateSubnet(ProjectID, subnetUpdateName, subnetUpdateDescription, subnetUpdateIPv6Enable,
			subnetUpdateDHCPEnable, subnetUpdatePrimaryDNS, subnetUpdateSecondaryDNS, subnetUpdateVpcId, subnetUpdateSubnetId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(sn, jmesPathQuery)
		}
	},
}

var subnetDeleteName string
var subnetDeleteSubnetID string
var subnetDeleteVpcID string
var subnetDeleteVpcName string
var subnetDeleteIsRecursive bool
var subnetDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete subnet",
	Long:  `Delete subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		var subnet *subnetModels.SubnetModel
		var err error
		if subnetDeleteSubnetID != "" {
			subnet, err = subnets.GetInfoAboutSubnet(ProjectID, subnetDeleteSubnetID)
		} else if subnetDeleteName != "" {
			subnet, err = subnets.GetSubnetByName(ProjectID, subnetDeleteName)
		} else {
			fmt.Println("ERROR: subnet ID or name is not specified")
			return
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		err = subnets.DeleteSubnet(ProjectID, subnet.VpcId, subnet.Id)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var subnetCreateTagKeys []string
var subnetCreateTagValues []string
var subnetCreateTagSubnetId string
var subnetCreateTagSubnetName string
var subnetCreateTagCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Command to add tags to subnet",
	Long:  `Command to add tags to subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := subnetUtils.GetSubnetId(subnetCreateTagSubnetId, subnetCreateTagSubnetName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		err = subnets.CreateSubnetTag(ProjectID, id, subnetCreateTagKeys, subnetCreateTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var subnetDeleteTagKeys []string
var subnetDeleteTagValues []string
var subnetDeleteTagSubnetId string
var subnetDeleteTagSubnetName string
var subnetDeleteTagCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Command to delete tags from subnet",
	Long:  `Command to delete tags from subnet`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := subnetUtils.GetSubnetId(subnetDeleteTagSubnetId, subnetDeleteTagSubnetName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		err = subnets.DeleteSubnetTag(ProjectID, id, subnetDeleteTagKeys, subnetDeleteTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var subnetGetSubnetTagsSubnetId string
var subnetGetSubnetTagsSubnetName string
var subnetGetTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Command to get VPC tags",
	Long:  `Command to get VPC tags`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := subnetUtils.GetSubnetId(subnetGetSubnetTagsSubnetId, subnetGetSubnetTagsSubnetName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		tags, err := subnets.GetSubnetTags(ProjectID, id)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(tags, jmesPathQuery)
		}
	},
}

var subnetGetByTagsKeys []string
var subnetGetByTagsValues []string
var subnetGetByTagsAction string
var subnetGetByTagsCmd = &cobra.Command{
	Use:   "get-by-tags",
	Short: "Command to get subnet tags",
	Long:  `Command to get subnet tags`,
	Run: func(cmd *cobra.Command, args []string) {
		if subnetGetByTagsAction == "count" {
			count, err := subnets.CountSubnetByTags(ProjectID, subnetGetByTagsKeys, subnetGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(count, jmesPathQuery)
			}
		} else {
			subnets, err := subnets.FilterSubnetByTags(ProjectID, subnetGetByTagsKeys, subnetGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(subnets, jmesPathQuery)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(subnetCmd)

	subnetCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	subnetCmd.AddCommand(subnetCreateCmd)
	subnetCreateCmd.Flags().StringVarP(&subnetCreateName, "name", "n", "", "Specifies the subnet name.")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateDescription, "decs", "d", "", "Provides supplementary information about the subnet.")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateCIDR, "cidr", "c", "192.168.0.0/16", "Specifies the subnet CIDR block.")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateGatewayIP, "gateway-ip", "g", "192.168.0.1", "Specifies the gateway of the subnet.")
	subnetCreateCmd.Flags().BoolVar(&subnetCreateIPv6Enable, "ipv6-en", false, "Specifies whether IPv6 is enabled. If IPv6 is enabled, you can use IPv6 CIDR blocks. The value can be true (enabled) or false (disabled).")
	subnetCreateCmd.Flags().BoolVar(&subnetCreateDHCPEnable, "dhcp-en", true, "Specifies whether DHCP is enabled for the subnet. The value can be true (enabled) or false (disabled).")
	subnetCreateCmd.Flags().StringVar(&subnetCreatePrimaryDNS, "primary-dns", "100.125.13.59", "Specifies the IP address of DNS server 1 on the subnet.")
	subnetCreateCmd.Flags().StringVar(&subnetCreateSecondaryDNS, "secondary-dns", "100.125.65.14", "Specifies the IP address of DNS server 2 on the subnet.")
	subnetCreateCmd.Flags().StringVar(&subnetCreateAvailabilityZones, "availability-zones", "", "Specifies the AZ to which the subnet belongs, which can be obtained from endpoints.")
	subnetCreateCmd.Flags().StringVar(&subnetCreateVpcName, "vpc-name", "", "Specifies the name of the VPC to which the subnet belongs.")
	subnetCreateCmd.Flags().StringVar(&subnetCreateVpcId, "vpc-id", "", "Specifies the ID of the VPC to which the subnet belongs.")

	subnetCmd.AddCommand(subnetUpdateCmd)
	subnetUpdateCmd.Flags().StringVarP(&subnetUpdateSubnetId, "id", "i", "", "Specifies the subnet ID.")
	subnetUpdateCmd.Flags().StringVarP(&subnetUpdateName, "name", "n", "", "Specifies the subnet name.")
	subnetUpdateCmd.Flags().StringVarP(&subnetUpdateDescription, "decs", "d", "", "Provides supplementary information about the subnet.")
	subnetUpdateCmd.Flags().BoolVar(&subnetUpdateIPv6Enable, "ipv6-en", false, "Specifies whether IPv6 is enabled. If IPv6 is enabled, you can use IPv6 CIDR blocks. The value can be true (enabled) or false (disabled).")
	subnetUpdateCmd.Flags().BoolVar(&subnetUpdateDHCPEnable, "dhcp-en", false, "Specifies whether DHCP is enabled for the subnet. The value can be true (enabled) or false (disabled).")
	subnetUpdateCmd.Flags().StringVar(&subnetUpdatePrimaryDNS, "primary-dns", "", "Specifies the IP address of DNS server 1 on the subnet.")
	subnetUpdateCmd.Flags().StringVar(&subnetUpdateSecondaryDNS, "secondary-dns", "", "Specifies the IP address of DNS server 2 on the subnet.")
	subnetUpdateCmd.Flags().StringVar(&subnetUpdateVpcId, "vpc-id", "", "Specifies the ID of the VPC to which the subnet belongs.")

	subnetCmd.AddCommand(subnetGetListCmd)
	subnetGetListCmd.Flags().IntVarP(&subnetListLimit, "limit", "l", 0, "Specifies the number of records that will be returned on each page. The value is from 0 to intmax.")
	subnetGetListCmd.Flags().StringVarP(&subnetListMarker, "marker", "m", "", "Specifies a resource ID for pagination query, indicating that the query starts from the next record of the specified resource ID.")
	subnetGetListCmd.Flags().StringVarP(&subnetListVpcID, "vpc-id", "v", "", "Specifies the ID of the VPC to which the subnet belongs")
	subnetGetListCmd.Flags().StringVar(&subnetListVpcName, "vpc-name", "", "Specifies the name of the VPC to which the subnet belongs")

	subnetCmd.AddCommand(subnetGetInfoCmd)
	subnetGetInfoCmd.Flags().StringVarP(&subnetInfoSubnetName, "name", "n", "", "Specifies the ID of the subnet.")
	subnetGetInfoCmd.Flags().StringVarP(&subnetInfoSubnetID, "id", "i", "", "Specifies the name of the subnet")

	subnetCmd.AddCommand(subnetDeleteCmd)
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteName, "name", "n", "", "Specifies the name of the subnet")
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteSubnetID, "id", "i", "", "Specifies the ID of the subnet")
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteVpcName, "vpc-name", "v", "", "Specifies the name of the VPC to which the subnet belongs")
	subnetDeleteCmd.Flags().StringVar(&subnetDeleteVpcID, "vpc-id", "", "Specifies the ID of the VPC to which the subnet belongs")
	subnetDeleteCmd.Flags().BoolVarP(&subnetDeleteIsRecursive, "rec", "r", false, "Specifies recursive delete flag")

	subnetCmd.AddCommand(subnetCreateTagCmd)
	subnetCreateTagCmd.Flags().StringVarP(&subnetCreateTagSubnetName, "name", "n", "", "Specifies the subnet name, which uniquely identifies the subnet.")
	subnetCreateTagCmd.Flags().StringVarP(&subnetCreateTagSubnetId, "id", "i", "", "Specifies the subnet ID, which uniquely identifies the subnet.")
	subnetCreateTagCmd.Flags().StringSliceVarP(&subnetCreateTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a subnet must be unique")
	subnetCreateTagCmd.Flags().StringSliceVarP(&subnetCreateTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	subnetCmd.AddCommand(subnetDeleteTagCmd)
	subnetDeleteTagCmd.Flags().StringVarP(&subnetDeleteTagSubnetName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	subnetDeleteTagCmd.Flags().StringVarP(&subnetDeleteTagSubnetId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	subnetDeleteTagCmd.Flags().StringSliceVarP(&subnetDeleteTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	subnetDeleteTagCmd.Flags().StringSliceVarP(&subnetDeleteTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	subnetCmd.AddCommand(subnetGetTagsCmd)
	subnetGetTagsCmd.Flags().StringVarP(&subnetGetSubnetTagsSubnetName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	subnetGetTagsCmd.Flags().StringVarP(&subnetGetSubnetTagsSubnetId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")

	subnetCmd.AddCommand(subnetGetByTagsCmd)
	subnetGetByTagsCmd.Flags().StringVarP(&subnetGetByTagsAction, "action", "a", "filter", "Specifies the operation to perform. The value can only be filter (filtering) or count (querying the total number).")
	subnetGetByTagsCmd.Flags().StringSliceVarP(&subnetGetByTagsKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	subnetGetByTagsCmd.Flags().StringSliceVarP(&subnetGetByTagsValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

}

//todo: Some flags might be persistent
