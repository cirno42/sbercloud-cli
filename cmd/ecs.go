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

var ecsDeleteIds []string
var ecsDeletePublicIp bool
var ecsDeleteVolume bool
var ecsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ECS",
	Long:  `Delete ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.DeleteEcs(ProjectID, ecsDeleteIds, ecsDeletePublicIp, ecsDeleteVolume)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ecs, jmesPathQuery)
		}
	},
}

var ecsCreateVpcID string
var ecsCreateImageRef string
var ecsCreateName string
var ecsCreateFlavorRef string
var ecsCreateRootVolumeType string
var ecsCreateSubnetIds []string
var ecsCreateSecGroupIds []string
var ecsCreateAdminPass string
var ecsCreateCount int
var ecsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Delete ECS",
	Long:  `Delete ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.CreateECS(ProjectID, ecsCreateVpcID, ecsCreateImageRef, ecsCreateName, ecsCreateFlavorRef,
			ecsCreateRootVolumeType, ecsCreateSubnetIds, ecsCreateSecGroupIds, ecsCreateAdminPass, ecsCreateCount)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ecs, jmesPathQuery)
		}
	},
}

var ecsBatchStartServerIds []string
var ecsBatchStartCmd = &cobra.Command{
	Use:   "batch-start",
	Short: "This command is used to start ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time.",
	Long:  `This command is used to start ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchStartEcs(ProjectID, ecsBatchStartServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsBatchRestartServerIds []string
var ecsBatchRestartType string
var ecsBatchRestartCmd = &cobra.Command{
	Use:   "batch-restart",
	Short: "This command is used to restart  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time",
	Long:  `This command is used to restart  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchRestartEcs(ProjectID, ecsBatchRestartType, ecsBatchRestartServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsBatchStopServerIds []string
var ecsBatchStopType string
var ecsBatchStopCmd = &cobra.Command{
	Use:   "batch-stop",
	Short: "This command is used to stop  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time",
	Long:  `This command is used to stop  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchStopEcs(ProjectID, ecsBatchStopType, ecsBatchStopServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
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
	ecsCmd.AddCommand(ecsCreateCmd)
	ecsCmd.AddCommand(ecsBatchStartCmd)
	ecsCmd.AddCommand(ecsBatchRestartCmd)
	ecsCmd.AddCommand(ecsBatchStopCmd)

	ecsFlavorListCmd.Flags().StringVarP(&ecsFlavorListAvailabilityZone, "availability_zone", "a", "", "")

	ecsInfoCmd.Flags().StringVarP(&ecsGetInfoId, "id", "i", "", "")

	ecsDeleteCmd.Flags().StringSliceVarP(&ecsDeleteIds, "id", "i", nil, "")
	ecsDeleteCmd.Flags().BoolVar(&ecsDeletePublicIp, "del-ip", false, "")
	ecsDeleteCmd.Flags().BoolVar(&ecsDeleteVolume, "del-vol", false, "")

	ecsCreateCmd.Flags().StringVar(&ecsCreateVpcID, "vpc-id", "", "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateImageRef, "image-ref", "", "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateName, "name", "", "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateFlavorRef, "flavor-ref", "", "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateRootVolumeType, "root-volume-type", "", "")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSubnetIds, "subnet-ids", nil, "")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSecGroupIds, "sg-ids", nil, "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateAdminPass, "admin-pass", "", "")
	ecsCreateCmd.Flags().IntVar(&ecsCreateCount, "count", 1, "")

	ecsBatchStartCmd.Flags().StringSliceVarP(&ecsBatchStartServerIds, "id", "i", nil, "Specifies ECS IDs")

	ecsBatchRestartCmd.Flags().StringSliceVarP(&ecsBatchRestartServerIds, "id", "i", nil, "Specifies ECS IDs")
	ecsBatchRestartCmd.Flags().StringVarP(&ecsBatchRestartType, "type", "t", "SOFT", "Specifies the type of the restart operation.")

	ecsBatchStopCmd.Flags().StringSliceVarP(&ecsBatchStopServerIds, "id", "i", nil, "Specifies ECS IDs")
	ecsBatchStopCmd.Flags().StringVarP(&ecsBatchStopType, "type", "t", "SOFT", "Specifies an ECS stop type.")

}
