package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/vpcModels"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var listMarker string
var listLimit int

var createName string
var createCidr string
var createDescription string

var infoName string
var infoVpcID string

var deleteName string
var deleteID string
var vpcDeleteIsRecursive bool

var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "Command to interact with VPC",
	Long:  `Command to interact with VPC`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var vpcCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create VPC",
	Long:  `Create VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		vpc, err := vpcs.CreateVpc(ProjectID, createName, createDescription, createCidr)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(vpc, jmesPathQuery)
		}
	},
}

var vpcGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of VPC",
	Long:  `Get list of VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		vpcs, err := vpcs.GetVpcsList(ProjectID, listLimit, listMarker)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		beautyfulPrints.PrintStruct(vpcs, jmesPathQuery)
	},
}

var vpcGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about VPC",
	Long:  `Get info about VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		var vpc *vpcModels.VpcModel
		var err error
		if infoVpcID != "" {
			vpc, err = vpcs.GetInfoAboutVpc(ProjectID, infoVpcID)
		} else {
			vpc, err = vpcs.GetVpcByName(ProjectID, infoName)
		}
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		} else {
			beautyfulPrints.PrintStruct(*vpc, jmesPathQuery)
		}
	},
}

var vpcUpdateVpcId string
var vpcUpdateVpcDesc string
var vpcUpdateVpcName string
var vpcUpdateVpcCidr string
var vpcUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update VPC",
	Long:  `Update VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		vpc, err := vpcs.UpdateVpc(ProjectID, vpcUpdateVpcId, vpcUpdateVpcName, vpcUpdateVpcDesc, vpcUpdateVpcCidr)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(vpc, jmesPathQuery)
		}
	},
}

var vpcDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete VPC",
	Long:  `Delete VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		if vpcDeleteIsRecursive {
			var vpcID string
			if deleteID == "" {
				vpc, err := vpcs.GetVpcByName(ProjectID, deleteName)
				if err != nil {
					fmt.Println("ERROR: " + err.Error())
					return
				}
				vpcID = vpc.Id
			} else {
				vpcID = deleteID
			}
			err := vpcs.DeleteVpcRecursive(ProjectID, vpcID)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())

			}
			return
		}
		var id string
		if deleteID != "" {
			id = deleteID
		} else {
			vpc, err := vpcs.GetVpcByName(ProjectID, deleteName)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				return
			}
			id = vpc.Id
		}
		err := vpcs.DeleteVpc(ProjectID, id)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(vpcCmd)

	vpcCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	vpcCmd.AddCommand(vpcCreateCmd)
	vpcCreateCmd.Flags().StringVarP(&createName, "name", "n", "", "Specifies the VPC name")
	vpcCreateCmd.Flags().StringVarP(&createCidr, "cidr", "c", "192.168.0.0/16", "Specifies the available IP address ranges for subnets in the VPC. Possible values are as follows: 10.0.0.0/8-24, 172.16.0.0/12-24, 192.168.0.0/16-24")
	vpcCreateCmd.Flags().StringVarP(&createDescription, "description", "d", "", "Provides supplementary information about the VPC")

	vpcCmd.AddCommand(vpcGetListCmd)
	vpcGetListCmd.Flags().IntVarP(&listLimit, "limit", "l", 0, "Specifies the number of records that will be returned on each page. The value is from 0 to intmax.")
	vpcGetListCmd.Flags().StringVarP(&listMarker, "marker", "m", "", "Specifies a resource ID for pagination query, indicating that the query starts from the next record of the specified resource ID.")

	vpcCmd.AddCommand(vpcGetInfoCmd)
	vpcGetInfoCmd.Flags().StringVarP(&infoName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcGetInfoCmd.Flags().StringVarP(&infoVpcID, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")

	vpcCmd.AddCommand(vpcUpdateCmd)
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcCidr, "cidr", "c", "", "Specifies the available IP address ranges for subnets in the VPC. Possible values are as follows: 10.0.0.0/8-24, 172.16.0.0/12-24, 192.168.0.0/16-24")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcDesc, "desc", "d", "", "Provides supplementary information about the VPC.")

	vpcCmd.AddCommand(vpcDeleteCmd)
	vpcDeleteCmd.Flags().StringVarP(&deleteName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcDeleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	vpcDeleteCmd.Flags().BoolVarP(&vpcDeleteIsRecursive, "rec", "r", false, "Specifies recursive delete flag")
}

//todo: Add "vpc-" prefix to every flag variable
