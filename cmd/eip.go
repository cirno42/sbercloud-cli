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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
	Short: "A brief description of your command",
	Long: `add
details
here`,
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

	eipCmdAssign.Flags().StringVar(&eipAssignPublicIPType, "eip-type", "5_bgp", "add details here")
	eipCmdAssign.Flags().IntVar(&eipAssignPublicIPVersion, "ip-version", 4, "add details here")
	eipCmdAssign.Flags().StringVarP(&eipAssignBandwidthName, "bandwidth-name", "b", "", "")
	eipCmdAssign.Flags().IntVar(&eipAssignBandwidthSize, "bandwidth-size", 0, "")
	eipCmdAssign.Flags().StringVar(&eipAssignBandwidthShareType, "bandwidth-share-type", "", "")

	eipCmdRelease.Flags().StringVarP(&eipReleasePublicIpID, "id", "i", "", "")
	eipCmdRelease.Flags().StringVarP(&eipReleasePublicIp, "address", "a", "", "")

	eipCmdList.Flags().IntVarP(&eipCmdListLimit, "limit", "l", 0, "")
	eipCmdList.Flags().StringVarP(&eipCmdListMarker, "marker", "m", "", "")

	eipCmdGetInfo.Flags().StringVarP(&eipCmdGetInfoAddress, "address", "a", "", "")
	eipCmdGetInfo.Flags().StringVarP(&eipCmdGetInfoPublicIpID, "id", "i", "", "")
}

//todo: вынести значения по умолчанию в константы
