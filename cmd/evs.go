package cmd

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/evs"
	"github.com/spf13/cobra"
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
var evsCreateWaitUntilSuccess bool
var evsCreateDiskCmd = &cobra.Command{
	Use:   "create-disk",
	Short: "commands to interact with EVS instances",
	Long:  `commands to interact with EVS instances`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := evs.CreateDisk(ProjectID, evsCreateName, evsCreateVolumeType, evsCreateAvailabilityZone, evsCreateCount, evsCreateSize, evsCreateMultiattach)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if evsCreateWaitUntilSuccess {
			diskIds, err := evs.WaitUntilJobSuccess(ProjectID, job.JobID)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				disks, err := evs.GetDisksById(ProjectID, diskIds)
				if err != nil {
					beautyfulPrints.PrintError(err)
				} else {
					beautyfulPrints.PrintStruct(disks, jmesPathQuery)
				}
			}
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
		job, err := evs.GetInfoAboutBatchTask(ProjectID, evsJobId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var evsListLimit int
var evsListOffset int
var evsListStatus string
var evsEvsGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "This command is used to query details about all disks.",
	Long:  `This command is used to query details about all disks.`,
	Run: func(cmd *cobra.Command, args []string) {
		evs, err := evs.GetDisksList(ProjectID, evsListStatus, evsListLimit, evsListOffset)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(evs, jmesPathQuery)
		}
	},
}

var evsDeleteVolumeId string
var evsEvsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This command is used to delete an EVS disk.",
	Long:  `This command is used to delete an EVS disk.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := evs.DeleteDisk(ProjectID, evsDeleteVolumeId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var evsExpandVolumeId string
var evsExpandNewSize int
var evsExpandCmd = &cobra.Command{
	Use:   "expand",
	Short: "This command is used to expand the capacity of an EVS disk.",
	Long:  `This command is used to expand the capacity of an EVS disk.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := evs.ExpandDisk(ProjectID, evsExpandVolumeId, evsExpandNewSize)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

func init() {
	RootCmd.AddCommand(evsCmd)
	evsCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	evsCmd.AddCommand(evsCreateDiskCmd)
	evsCmd.AddCommand(evsJobInfoCmd)
	evsCmd.AddCommand(evsEvsGetListCmd)
	evsCmd.AddCommand(evsEvsDeleteCmd)
	evsCmd.AddCommand(evsExpandCmd)

	evsCreateDiskCmd.Flags().IntVarP(&evsCreateCount, "count", "c", 1, "")
	evsCreateDiskCmd.Flags().IntVarP(&evsCreateSize, "size", "s", 0, "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateVolumeType, "volume-type", "t", "SAS", "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateName, "name", "n", "", "")
	evsCreateDiskCmd.Flags().StringVarP(&evsCreateAvailabilityZone, "az", "a", "AZ1", "")
	evsCreateDiskCmd.Flags().BoolVarP(&evsCreateMultiattach, "multiattach", "m", false, "")
	evsCreateDiskCmd.Flags().BoolVar(&evsCreateWaitUntilSuccess, "wait-until-success", true, "")

	evsJobInfoCmd.Flags().StringVarP(&evsJobId, "id", "i", "", "")

	evsEvsGetListCmd.Flags().IntVarP(&evsListLimit, "limit", "l", 0, "")
	evsEvsGetListCmd.Flags().IntVarP(&evsListOffset, "offset", "o", 0, "")
	evsEvsGetListCmd.Flags().StringVarP(&evsListStatus, "status", "s", "", "")

	evsEvsDeleteCmd.Flags().StringVarP(&evsDeleteVolumeId, "id", "i", "", "")

	evsExpandCmd.Flags().StringVarP(&evsExpandVolumeId, "id", "i", "", "Specifies the ID of the disk.")
	evsExpandCmd.Flags().IntVarP(&evsExpandNewSize, "size", "s", 0, "Specifies the size of the disk after capacity expansion, in GB.")
}
