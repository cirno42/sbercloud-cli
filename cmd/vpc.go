package cmd

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/models/vpcModels"
	"github.com/cirno42/sbercloud-api/api/vpcs"
	"github.com/spf13/cobra"
	"sbercloud-cli/internal/beautyfulPrints"
	"sbercloud-cli/internal/utils/vpcUtils"
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
			beautyfulPrints.PrintError(err)
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
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(*vpc, jmesPathQuery)
		}
	},
}

var vpcUpdateVpcId string
var vpcUpdateVpcDesc string
var vpcUpdateOldVpcName string
var vpcUpdateVpcName string
var vpcUpdateVpcCidr string
var vpcUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update VPC",
	Long:  `Update VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := vpcUtils.GetVpcId(vpcUpdateVpcId, vpcUpdateOldVpcName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		vpc, err := vpcs.UpdateVpc(ProjectID, id, vpcUpdateVpcName, vpcUpdateVpcDesc, vpcUpdateVpcCidr)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(vpc.Vpc, jmesPathQuery)
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
					beautyfulPrints.PrintError(err)
					return
				}
				vpcID = vpc.Id
			} else {
				vpcID = deleteID
			}
			err := vpcs.DeleteVpcRecursive(ProjectID, vpcID)
			if err != nil {
				beautyfulPrints.PrintError(err)

			}
			return
		}
		var id string
		if deleteID != "" {
			id = deleteID
		} else {
			vpc, err := vpcs.GetVpcByName(ProjectID, deleteName)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			}
			id = vpc.Id
		}
		err := vpcs.DeleteVpc(ProjectID, id)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var vpcCreateTagKeys []string
var vpcCreateTagValues []string
var vpcCreateTagVpcId string
var vpcCreateTagVpcName string
var vpcCreateTagCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Command to add tags to VPC",
	Long:  `Command to add tags to VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := vpcUtils.GetVpcId(vpcCreateTagVpcId, vpcCreateTagVpcName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		err = vpcs.CreateVpcTag(ProjectID, id, vpcCreateTagKeys, vpcCreateTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var vpcDeleteTagKeys []string
var vpcDeleteTagValues []string
var vpcDeleteTagVpcId string
var vpcDeleteTagVpcName string
var vpcDeleteTagCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Command to delete tags from VPC",
	Long:  `Command to delete tags from VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := vpcUtils.GetVpcId(vpcDeleteTagVpcId, vpcDeleteTagVpcName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		err = vpcs.DeleteVpcTag(ProjectID, id, vpcDeleteTagKeys, vpcDeleteTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var vpcGetVpcTagsVpcId string
var vpcGetVpcTagsVpcName string
var vpcGetTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Command to get VPC tags",
	Long:  `Command to get VPC tags`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := vpcUtils.GetVpcId(vpcGetVpcTagsVpcId, vpcGetVpcTagsVpcName, ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		tags, err := vpcs.GetVpcTags(ProjectID, id)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(tags, jmesPathQuery)
		}
	},
}

var vpcGetByTagsKeys []string
var vpcGetByTagsValues []string
var vpcGetByTagsAction string
var vpcGetByTagsCmd = &cobra.Command{
	Use:   "get-by-tags",
	Short: "Command to get VPC tags",
	Long:  `Command to get VPC tags`,
	Run: func(cmd *cobra.Command, args []string) {
		if vpcGetByTagsAction == "count" {
			count, err := vpcs.CountVpcByTags(ProjectID, vpcGetByTagsKeys, vpcGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(count, jmesPathQuery)
			}
		} else {
			vpcs, err := vpcs.FilterVpcByTags(ProjectID, vpcGetByTagsKeys, vpcGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(vpcs, jmesPathQuery)
			}
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
	vpcUpdateCmd.Flags().StringVar(&vpcUpdateOldVpcName, "old-name", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcName, "new-name", "n", "", "")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcCidr, "cidr", "c", "", "Specifies the available IP address ranges for subnets in the VPC. Possible values are as follows: 10.0.0.0/8-24, 172.16.0.0/12-24, 192.168.0.0/16-24")
	vpcUpdateCmd.Flags().StringVarP(&vpcUpdateVpcDesc, "desc", "d", "", "Provides supplementary information about the VPC.")

	vpcCmd.AddCommand(vpcDeleteCmd)
	vpcDeleteCmd.Flags().StringVarP(&deleteName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcDeleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	vpcDeleteCmd.Flags().BoolVarP(&vpcDeleteIsRecursive, "rec", "r", false, "Specifies recursive delete flag")

	vpcCmd.AddCommand(vpcCreateTagCmd)
	vpcCreateTagCmd.Flags().StringVarP(&vpcCreateTagVpcName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcCreateTagCmd.Flags().StringVarP(&vpcCreateTagVpcId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	vpcCreateTagCmd.Flags().StringSliceVarP(&vpcCreateTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	vpcCreateTagCmd.Flags().StringSliceVarP(&vpcCreateTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	vpcCmd.AddCommand(vpcDeleteTagCmd)
	vpcDeleteTagCmd.Flags().StringVarP(&vpcDeleteTagVpcName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcDeleteTagCmd.Flags().StringVarP(&vpcDeleteTagVpcId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")
	vpcDeleteTagCmd.Flags().StringSliceVarP(&vpcDeleteTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	vpcDeleteTagCmd.Flags().StringSliceVarP(&vpcDeleteTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	vpcCmd.AddCommand(vpcGetTagsCmd)
	vpcGetTagsCmd.Flags().StringVarP(&vpcGetVpcTagsVpcName, "name", "n", "", "Specifies the VPC name, which uniquely identifies the VPC.")
	vpcGetTagsCmd.Flags().StringVarP(&vpcGetVpcTagsVpcId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the VPC.")

	vpcCmd.AddCommand(vpcGetByTagsCmd)
	vpcGetByTagsCmd.Flags().StringVarP(&vpcGetByTagsAction, "action", "a", "filter", "Specifies the operation to perform. The value can only be filter (filtering) or count (querying the total number).")
	vpcGetByTagsCmd.Flags().StringSliceVarP(&vpcGetByTagsKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	vpcGetByTagsCmd.Flags().StringSliceVarP(&vpcGetByTagsValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")
}

//todo: Add "vpc-" prefix to every flag variable
