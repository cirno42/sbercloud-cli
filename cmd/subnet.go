package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/subnetModels"
	"sbercloud-cli/api/subnets"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subnet called")
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
			fmt.Println(err)
		} else {
			beautyfulPrints.PrintStruct(sn, jmesPathQuery)
		}
	},
}

var subnetListLimit int
var subnetListMarker string
var subnetListVpcID string

var subnetGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		subnets, err := subnets.GetSubnetsList(ProjectID, subnetListLimit, subnetListMarker, subnetListVpcID)
		if err != nil {
			fmt.Println(err)
		} else {
			beautyfulPrints.PrintStruct(subnets, jmesPathQuery)
		}
	},
}

var subnetInfoSubnetID string
var subnetInfoSubnetName string
var subnetGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		var subnet *subnetModels.SubnetModel
		var err error
		if subnetInfoSubnetID != "" {
			subnet, err = subnets.GetInfoAboutSubnet(ProjectID, subnetInfoSubnetID)
		} else {
			subnet, err = subnets.GetSubnetByName(ProjectID, subnetInfoSubnetName)
		}
		if err != nil {
			fmt.Println(err)
		} else {
			beautyfulPrints.PrintStruct(subnet, jmesPathQuery)
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
			fmt.Println("ERROR: " + err.Error())
			return
		}
		err = subnets.DeleteSubnet(ProjectID, subnet.VpcId, subnet.Id)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(subnetCmd)

	subnetCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	subnetCmd.AddCommand(subnetCreateCmd)
	subnetCreateCmd.Flags().StringVarP(&subnetCreateName, "name", "n", "", "add details here")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateDescription, "decs", "d", "", "add details here")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateCIDR, "cidr", "c", "192.168.0.0/16", "add details here")
	subnetCreateCmd.Flags().StringVarP(&subnetCreateGatewayIP, "gateway-ip", "g", "192.168.0.1", "add details here")
	subnetCreateCmd.Flags().BoolVar(&subnetCreateIPv6Enable, "ipv6-en", false, "add details here")
	subnetCreateCmd.Flags().BoolVar(&subnetCreateDHCPEnable, "dhcp-en", false, "add details here")
	subnetCreateCmd.Flags().StringVar(&subnetCreatePrimaryDNS, "primary-dns", "", "add details here")
	subnetCreateCmd.Flags().StringVar(&subnetCreateSecondaryDNS, "secondary-dns", "", "add details here")
	subnetCreateCmd.Flags().StringVar(&subnetCreateAvailabilityZones, "availability-zones", "", "")
	subnetCreateCmd.Flags().StringVar(&subnetCreateVpcName, "vpc-name", "", "")
	subnetCreateCmd.Flags().StringVar(&subnetCreateVpcId, "vpc-id", "", "")

	subnetCmd.AddCommand(subnetGetListCmd)
	subnetGetListCmd.Flags().IntVarP(&subnetListLimit, "limit", "l", 0, "add details here")
	subnetGetListCmd.Flags().StringVarP(&subnetListMarker, "marker", "m", "", "add details here")
	subnetGetListCmd.Flags().StringVarP(&subnetListVpcID, "vpc-id", "v", "", "add details here")

	subnetCmd.AddCommand(subnetGetInfoCmd)
	subnetGetInfoCmd.Flags().StringVarP(&subnetInfoSubnetName, "name", "n", "", "add details here")
	subnetGetInfoCmd.Flags().StringVarP(&subnetInfoSubnetID, "id", "i", "", "add details here")

	subnetCmd.AddCommand(subnetDeleteCmd)
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteName, "name", "n", "", "add details")
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteSubnetID, "id", "i", "", "add details here")
	subnetDeleteCmd.Flags().StringVarP(&subnetDeleteVpcName, "vpc-name", "v", "", "add details here")
	subnetDeleteCmd.Flags().StringVar(&subnetDeleteVpcID, "vpc-id", "", "add details here")
	subnetDeleteCmd.Flags().BoolVarP(&subnetDeleteIsRecursive, "rec", "r", false, "add details here")
}

//todo: Some flags might be persistent
