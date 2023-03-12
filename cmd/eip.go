package cmd

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/eip"
	"github.com/cirno42/sbercloud-api/api/models/eipModels"
	"github.com/spf13/cobra"
	"sbercloud-cli/internal/beautyfulPrints"
)

var eipCmd = &cobra.Command{
	Use:   "eip",
	Short: "commands to interact with Elastic IPs",
	Long:  `commands to interact with Elastic IPs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("eip called")
	},
}

var eipAssignPublicIPType string
var eipAssignPublicIPVersion int
var eipAssignBandwidthName string
var eipAssignBandwidthSize int
var eipAssignBandwidthShareType string

var eipCmdAssign = &cobra.Command{
	Use:   "assign",
	Short: "Assign Elastic IP",
	Long:  `Assign Elastic IP`,
	Run: func(cmd *cobra.Command, args []string) {
		eip, err := eip.AssignEIP(ProjectID, eipAssignPublicIPType, eipAssignPublicIPVersion, eipAssignBandwidthName, eipAssignBandwidthSize, eipAssignBandwidthShareType)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(eip, jmesPathQuery)
		}
	},
}

var eipReleasePublicIpID string
var eipReleasePublicIp string
var eipCmdRelease = &cobra.Command{
	Use:   "release",
	Short: "Release Elastic IP",
	Long:  `Release Elastic IP`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if eipReleasePublicIpID != "" {
			err = eip.DeletePublicIP(ProjectID, eipReleasePublicIpID)
		} else if eipReleasePublicIp != "" {
			err = eip.DeletePublicIPByAddress(ProjectID, eipReleasePublicIp)
		} else {
			fmt.Printf("ERROR: public IP address and ID are both not specified\n")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("Public ip released successfully")
		}
	},
}

var eipCmdGetInfoAddress string
var eipCmdGetInfoPublicIpID string
var eipCmdGetInfo = &cobra.Command{
	Use:   "info",
	Short: "Get info about Elastic IP you choose",
	Long:  `Get info about Elastic IP you choose`,
	Run: func(cmd *cobra.Command, args []string) {
		var eipEntity *eipModels.EipModel
		var err error
		if eipCmdGetInfoPublicIpID != "" {
			eipEntity, err = eip.GetEIPInfo(ProjectID, eipCmdGetInfoPublicIpID)
		} else if eipCmdGetInfoAddress != "" {
			eipEntity, err = eip.GetInfoAboutEIPByAddress(ProjectID, eipCmdGetInfoAddress)
		} else {
			fmt.Printf("ERROR: public IP address and ID are both not specified\n")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(eipEntity, jmesPathQuery)
		}
	},
}

var eipCmdListLimit int
var eipCmdListMarker string
var eipCmdList = &cobra.Command{
	Use:   "list",
	Short: "Get list of Elastic IPs",
	Long:  `Get list of Elastic IPs`,
	Run: func(cmd *cobra.Command, args []string) {
		eips, err := eip.GetEIPsList(ProjectID, eipCmdListLimit, eipCmdListMarker)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(eips, jmesPathQuery)
		}
	},
}

var listActiveIPsAllProject bool
var eipActiveListCmd = &cobra.Command{
	Use:   "list-active",
	Short: "Get list of active Elastic IPs",
	Long:  `Get list of active Elastic IPs`,
	Run: func(cmd *cobra.Command, args []string) {
		if listActiveIPsAllProject {
			eips, err := eip.GetActiveIPsInAllProjects()
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(eips, jmesPathQuery)
			}
		} else {
			eips, err := eip.GetActiveIPsInSpecifiedProject(ProjectID)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(eips, jmesPathQuery)
			}
		}
	},
}

var eipCreateTagKeys []string
var eipCreateTagValues []string
var eipCreateTagEipId string
var eipCreateTagCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Command to add tags to EIP",
	Long:  `Command to add tags to EIP`,
	Run: func(cmd *cobra.Command, args []string) {
		err := eip.CreateEipTag(ProjectID, eipCreateTagEipId, eipCreateTagKeys, eipCreateTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var eipDeleteTagKeys []string
var eipDeleteTagValues []string
var eipDeleteTagEipId string
var eipDeleteTagCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Command to delete tags from EIP",
	Long:  `Command to delete tags from EIP`,
	Run: func(cmd *cobra.Command, args []string) {
		err := eip.DeleteEipTag(ProjectID, eipDeleteTagEipId, eipDeleteTagKeys, eipDeleteTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var eipGetEipTagsEipId string
var eipGetTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Command to get VPC tags",
	Long:  `Command to get VPC tags`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := eip.GetEipTags(ProjectID, eipGetEipTagsEipId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(tags, jmesPathQuery)
		}
	},
}

var eipGetByTagsKeys []string
var eipGetByTagsValues []string
var eipGetByTagsAction string
var eipGetByTagsCmd = &cobra.Command{
	Use:   "get-by-tags",
	Short: "Command to get subnet tags",
	Long:  `Command to get subnet tags`,
	Run: func(cmd *cobra.Command, args []string) {
		if eipGetByTagsAction == "count" {
			count, err := eip.CountEipByTags(ProjectID, eipGetByTagsKeys, eipGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(count, jmesPathQuery)
			}
		} else {
			eips, err := eip.FilterEipByTags(ProjectID, eipGetByTagsKeys, eipGetByTagsValues)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(eips, jmesPathQuery)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(eipCmd)
	eipCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	eipCmd.AddCommand(eipCmdAssign)
	eipCmd.AddCommand(eipCmdRelease)
	eipCmd.AddCommand(eipCmdGetInfo)
	eipCmd.AddCommand(eipCmdList)
	eipCmd.AddCommand(eipActiveListCmd)

	eipCmdAssign.Flags().StringVar(&eipAssignPublicIPType, "eip-type", "5_bgp", "Specifies Type of EIP. The value can be 5_bgp, default is 5_bgp")
	eipCmdAssign.Flags().IntVar(&eipAssignPublicIPVersion, "ip-version", 4, "Specifies Version of IP. The value can be 4 or 6, default is 4.")
	eipCmdAssign.Flags().StringVarP(&eipAssignBandwidthName, "bandwidth-name", "b", "", "Specifies Bandwidth name. The value is a string of 1 to 64 characters that can contain letters, digits, underscores (_), hyphens (-), and periods (.). This parameter is mandatory when share_type is set to PER. This parameter will be ignored when share_type is set to WHOLE with an ID specified.")
	eipCmdAssign.Flags().IntVar(&eipAssignBandwidthSize, "bandwidth-size", 0, "Specifies the bandwidth size. The value ranges from 1 Mbit/s to 300 Mbit/s by default. (The specific range may vary depending on the configuration in each region. You can see the bandwidth range of each region on the management console.) This parameter is mandatory when share_type is set to PER. This parameter will be ignored when share_type is set to WHOLE with an ID specified.")
	eipCmdAssign.Flags().StringVar(&eipAssignBandwidthShareType, "bandwidth-share-type", "", "Specifies the bandwidth type. Possible values are PER (Dedicated bandwidth) and WHOLE (Shared bandwidth). If this parameter is set to WHOLE, the bandwidth ID must be specified")

	eipCmdRelease.Flags().StringVarP(&eipReleasePublicIpID, "id", "i", "", "Specifies ID of IP to release")
	eipCmdRelease.Flags().StringVarP(&eipReleasePublicIp, "address", "a", "", "Specifies IP address to release")

	eipCmdList.Flags().IntVarP(&eipCmdListLimit, "limit", "l", 1000, "Specifies the number of records returned on each page")
	eipCmdList.Flags().StringVarP(&eipCmdListMarker, "marker", "m", "", "Specifies the start resource ID of pagination query. If the parameter is left blank, only resources on the first page are queried.")

	eipCmdGetInfo.Flags().StringVarP(&eipCmdGetInfoAddress, "address", "a", "", "Specifies IP address of EIP")
	eipCmdGetInfo.Flags().StringVarP(&eipCmdGetInfoPublicIpID, "id", "i", "", "Specifies ID of EIP")

	eipActiveListCmd.Flags().BoolVarP(&listActiveIPsAllProject, "all-projects", "a", false, "")

	eipCmd.AddCommand(eipCreateTagCmd)
	eipCreateTagCmd.Flags().StringVarP(&eipCreateTagEipId, "id", "i", "", "Specifies the EIP ID, which uniquely identifies the EIP.")
	eipCreateTagCmd.Flags().StringSliceVarP(&eipCreateTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a subnet must be unique")
	eipCreateTagCmd.Flags().StringSliceVarP(&eipCreateTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	eipCmd.AddCommand(eipDeleteTagCmd)
	eipDeleteTagCmd.Flags().StringVarP(&eipDeleteTagEipId, "id", "i", "", "Specifies the VPC ID, which uniquely identifies the EIP.")
	eipDeleteTagCmd.Flags().StringSliceVarP(&eipDeleteTagKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	eipDeleteTagCmd.Flags().StringSliceVarP(&eipDeleteTagValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")

	eipCmd.AddCommand(eipGetTagsCmd)
	eipGetTagsCmd.Flags().StringVarP(&eipGetEipTagsEipId, "id", "i", "", "Specifies the EIP ID, which uniquely identifies the EIP.")

	eipCmd.AddCommand(eipGetByTagsCmd)
	eipGetByTagsCmd.Flags().StringVarP(&eipGetByTagsAction, "action", "a", "filter", "Specifies the operation to perform. The value can only be filter (filtering) or count (querying the total number).")
	eipGetByTagsCmd.Flags().StringSliceVarP(&eipGetByTagsKeys, "keys", "k", nil, "Specifies comma-separated tag keys. Key cannot be left blank. Key can contain a maximum of 36 characters. The tag key of a VPC must be unique")
	eipGetByTagsCmd.Flags().StringSliceVarP(&eipGetByTagsValues, "values", "v", nil, "Specifies comma-separated tag values. Tag value can contain a maximum of 43 characters.")
}

//todo: вынести значения по умолчанию в константы
