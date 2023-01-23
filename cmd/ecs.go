package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/ecs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "commands to interact with ECS instances",
	Long:  `commands to interact with ECS instances`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("esc called")
	},
}

var ecsFlavorListAvailabilityZone string
var ecsFlavorListCmd = &cobra.Command{
	Use:   "flavor-list",
	Short: "Get flavor list",
	Long:  `Get flavor list`,
	Run: func(cmd *cobra.Command, args []string) {
		flavors, err := ecs.GetESCFlavorList(ProjectID, ecsFlavorListAvailabilityZone)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(flavors, jmesPathQuery)
		}
	},
}

var ecsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get ECS list",
	Long:  `Get ECS list`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.GetECSList(ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ecs, jmesPathQuery)
		}
	},
}

var ecsGetInfoId string
var ecsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about ECS",
	Long:  `Get info about ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.GetInfoAboutEcs(ProjectID, ecsGetInfoId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ecs, jmesPathQuery)
		}
	},
}

var ecsGetDeleteIds []string
var deletePublicIp bool
var deleteVolume bool
var ecsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ECS",
	Long:  `Delete ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.DeleteEcs(ProjectID, ecsGetDeleteIds, deletePublicIp, deleteVolume)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ecs, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(ecsCmd)
	ecsCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	ecsCmd.AddCommand(ecsFlavorListCmd)
	ecsCmd.AddCommand(ecsListCmd)
	ecsCmd.AddCommand(ecsInfoCmd)
	ecsCmd.AddCommand(ecsDeleteCmd)

	ecsFlavorListCmd.Flags().StringVarP(&ecsFlavorListAvailabilityZone, "availability_zone", "a", "", "")

	ecsInfoCmd.Flags().StringVarP(&ecsGetInfoId, "id", "i", "", "")

	ecsDeleteCmd.Flags().StringSliceVarP(&ecsGetDeleteIds, "id", "i", nil, "")
	ecsDeleteCmd.Flags().BoolVar(&deletePublicIp, "del-ip", false, "")
	ecsDeleteCmd.Flags().BoolVar(&deleteVolume, "del-vol", false, "")
}
