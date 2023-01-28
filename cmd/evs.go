package cmd

import (
	"github.com/spf13/cobra"
	"sbercloud-cli/api/evs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var evsCmd = &cobra.Command{
	Use:   "evs",
	Short: "commands to interact with EVS instances",
	Long:  `commands to interact with EVS instances`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var evsCreateCount int
var evsCreateSize int
var evsCreateVolumeType string
var evsCreateName string
var evsCreateMultiattach bool
var evsCreateAvailabilityZone string
var evsCreateDiskCmd = &cobra.Command{
	Use:   "create-disk",
	Short: "commands to interact with EVS instances",
	Long:  `commands to interact with EVS instances`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := evs.CreateDisk(ProjectID, evsCreateName, evsCreateVolumeType, evsCreateAvailabilityZone, evsCreateCount, evsCreateSize, evsCreateMultiattach)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var evsJobId string
var evsJobInfoCmd = &cobra.Command{
	Use:   "job-info",
	Short: "commands to interact with ECS instances",
	Long:  `commands to interact with ECS instances`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := evs.GetInfoAboutTask(ProjectID, evsJobId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(evsCmd)

	evsCmd.AddCommand(evsCreateDiskCmd)
	evsCmd.AddCommand(evsJobInfoCmd)

	evsCreateDiskCmd.Flags().IntVarP(&evsCreateCount, "count", "c", 1, "")
	evsCreateDiskCmd.Flags().IntVarP(&evsCreateSize, "size", "s", 0, "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateVolumeType, "volume-type", "t", "SAS", "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateName, "name", "n", "", "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateAvailabilityZone, "az", "a", "AZ1", "")
	evsCreateDiskCmd.Flags().BoolVarP(&evsCreateMultiattach, "multiattach", "m", false, "")

	evsJobInfoCmd.Flags().StringVarP(&evsJobId, "id", "i", "", "")
}
