package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/eip"
	"sbercloud-cli/api/models/eipModels"
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
			fmt.Printf("ERROR: %s\n", err.Error())
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
			fmt.Printf("%s\n", err)
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
			fmt.Printf("ERROR: ", err.Error())
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
			fmt.Printf("ERROR: %s\n", err.Error())
		} else {
			beautyfulPrints.PrintStruct(eips, jmesPathQuery)
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
}

//todo: вынести значения по умолчанию в константы
