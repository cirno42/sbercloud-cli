package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/ecs"
	"sbercloud-cli/api/subnets"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "commands to interact with ECS instances",
	Long:  `commands to interact with ECS instances`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var ecsFlavorListAvailabilityZone string
var ecsFlavorListVcpus int
var ecsFlavorListRam int
var ecsFlavorListType string
var ecsFlavorListGen string
var ecsFlavorListCmd = &cobra.Command{
	Use:   "flavor-list",
	Short: "Get flavor list",
	Long:  `Get flavor list`,
	Run: func(cmd *cobra.Command, args []string) {
		flavors, err := ecs.GetFlavorListBySpec(ProjectID, ecsFlavorListGen, ecsFlavorListType, ecsFlavorListAvailabilityZone, ecsFlavorListRam, ecsFlavorListVcpus)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(flavors, jmesPathQuery)
		}
	},
}

var ecsGetFlavorAvailabilityZone string
var ecsGetFlavorVcpus int
var ecsGetFlavorRam int
var ecsGetFlavorType string
var ecsGetFlavorGen string
var ecsGetFlavorCmd = &cobra.Command{
	Use:   "get-flavor",
	Short: "Get flavor",
	Long:  `Get flavor`,
	Run: func(cmd *cobra.Command, args []string) {
		flavor, err := ecs.GetMinimumFlavorBySpec(ProjectID, ecsGetFlavorGen, ecsGetFlavorType, ecsGetFlavorAvailabilityZone, ecsGetFlavorRam, ecsGetFlavorVcpus)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(flavor, jmesPathQuery)
		}
	},
}

var ecsListOffset int
var ecsListLimit int
var ecsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get ECS list",
	Long:  `Get ECS list`,
	Run: func(cmd *cobra.Command, args []string) {
		ecs, err := ecs.GetECSList(ProjectID, ecsListOffset, ecsListLimit)
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
var ecsDeleteWaitUntilSuccess bool
var ecsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ECS",
	Long:  `Delete ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.DeleteEcs(ProjectID, ecsDeleteIds, ecsDeletePublicIp, ecsDeleteVolume)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsDeleteWaitUntilSuccess {
			res, err := ecs.WaitUntilJobSuccessAndGetStatus(ProjectID, job.JobID)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			} else {
				fmt.Println(res)
			}
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsCreateVpcID string
var ecsCreateImageRef string
var ecsCreateName string
var ecsCreateVpcName string
var ecsCreateSubnetNames []string
var ecsCreateSGNames []string
var ecsCreateFlavorRef string
var ecsCreateRootVolumeType string
var ecsCreateSubnetIds []string
var ecsCreateSecGroupIds []string
var ecsCreateAdminPass string
var ecsCreateEipId string
var ecsCreateEipBandwidthSize int
var ecsCreateEipBandwidthType string
var ecsCreateCount int
var ecsCreateEipType string
var ecsCreateVolumeTypes []string
var ecsCreateVolumeSizes []int
var ecsCreateRootVolumeSize int
var ecsCreateAvailabilityZone string
var ecsCreateKeyName string
var ecsCreateVcpus int
var ecsCreateRam int
var ecsCreateFlavorType string
var ecsCreateFlavorGen string
var ecsCreateAssignEip bool
var ecsCreateWaitUntilSuccess bool
var ecsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create ECS",
	Long:  `Create ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		flavorRef := ecsCreateFlavorRef
		if flavorRef == "" {
			flavor, err := ecs.GetMinimumFlavorBySpec(ProjectID, ecsCreateFlavorGen, ecsCreateFlavorType, ecsCreateAvailabilityZone, ecsCreateRam, ecsCreateVcpus)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			}
			flavorRef = flavor.ID
		}
		var vpcId string
		if ecsCreateVpcID == "" {
			vpc, err := vpcs.GetVpcByName(ProjectID, ecsCreateVpcName)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			}
			vpcId = vpc.Id
		} else {
			vpcId = ecsCreateVpcID
		}
		var subnetIds []string
		if ecsCreateSubnetIds == nil {
			subnets, err := subnets.GetSubnetsByNames(ProjectID, ecsCreateSubnetNames)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			}
			subnetIds = make([]string, len(subnets))
			for i, subnet := range subnets {
				subnetIds[i] = subnet.Id
			}
		} else {
			subnetIds = ecsCreateSubnetIds
		}
		if (ecsCreateAssignEip) && (ecsCreateEipBandwidthSize == 0) {
			ecsCreateEipBandwidthSize = 5
		}
		createdEcs, err := ecs.CreateECS(ProjectID, vpcId, ecsCreateImageRef, ecsCreateName, flavorRef,
			ecsCreateRootVolumeType, ecsCreateAvailabilityZone, ecsCreateEipId, ecsCreateEipType, ecsCreateEipBandwidthType, ecsCreateEipBandwidthSize, ecsCreateVolumeTypes,
			subnetIds, ecsCreateSecGroupIds, ecsCreateVolumeSizes, ecsCreateAdminPass, ecsCreateKeyName, ecsCreateRootVolumeSize, ecsCreateCount)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsCreateWaitUntilSuccess {
			res, err := ecs.WaitUntilJobSuccess(ProjectID, createdEcs.JobID)
			servers, err := ecs.GetListEcsById(ProjectID, res)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(servers, jmesPathQuery)
			}
		} else {
			beautyfulPrints.PrintStruct(createdEcs, jmesPathQuery)
		}
	},
}

var ecsBatchStartServerIds []string
var ecsStartWaitUntilSuccess bool
var ecsBatchStartCmd = &cobra.Command{
	Use:   "batch-start",
	Short: "This command is used to start ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time.",
	Long:  `This command is used to start ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchStartEcs(ProjectID, ecsBatchStartServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsStartWaitUntilSuccess {
			res, err := ecs.WaitUntilJobSuccess(ProjectID, job.JobID)
			servers, err := ecs.GetListEcsById(ProjectID, res)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(servers, jmesPathQuery)
			}
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsBatchRestartServerIds []string
var ecsBatchRestartType string
var ecsRestartWaitUntilSuccess bool
var ecsBatchRestartCmd = &cobra.Command{
	Use:   "batch-restart",
	Short: "This command is used to restart  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time",
	Long:  `This command is used to restart  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchRestartEcs(ProjectID, ecsBatchRestartType, ecsBatchRestartServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsRestartWaitUntilSuccess {
			res, err := ecs.WaitUntilJobSuccess(ProjectID, job.JobID)
			servers, err := ecs.GetListEcsById(ProjectID, res)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(servers, jmesPathQuery)
			}
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsBatchStopServerIds []string
var ecsBatchStopType string
var ecsStopWaitUntilSuccess bool
var ecsBatchStopCmd = &cobra.Command{
	Use:   "batch-stop",
	Short: "This command is used to stop  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time",
	Long:  `This command is used to stop  ECSs in a batch based on specified ECS IDs. A maximum of 1000 ECSs can be started at a time`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.BatchStopEcs(ProjectID, ecsBatchStopType, ecsBatchStopServerIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsStopWaitUntilSuccess {
			res, err := ecs.WaitUntilJobSuccess(ProjectID, job.JobID)
			servers, err := ecs.GetListEcsById(ProjectID, res)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(servers, jmesPathQuery)
			}
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsBatchAddNicsSubnetIds []string
var ecsBatchAddNicsServerId string
var ecsBatchAddNicsCmd = &cobra.Command{
	Use:   "add-nics",
	Short: "This command is used to add nics to ECS",
	Long:  `This command is used to add nics to ECS`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.AddNicsBatchToEcs(ProjectID, ecsBatchAddNicsServerId, ecsBatchAddNicsSubnetIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var ecsBatchDeleteNicsSubnetIds []string
var ecsBatchDeleteNicsServerId string
var ecsBatchDeleteNicsCmd = &cobra.Command{
	Use:   "delete-nics",
	Short: "This command is used to uninstall and delete one or multiple NICs from an ECS.",
	Long:  `This command is used to uninstall and delete one or multiple NICs from an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.DeleteNicsBatchToEcs(ProjectID, ecsBatchDeleteNicsServerId, ecsBatchDeleteNicsSubnetIds)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var ecsJobId string
var jobInfoCmd = &cobra.Command{
	Use:   "job-info",
	Short: "This command is used to get info about job",
	Long:  `This command is used to get info about job`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.GetInfoAboutTask(ProjectID, ecsJobId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsAttachDiskVolumeId string
var ecsAttachDiskEcsId string
var ecsAttachDiskDevice string
var ecsAttachDiskWaitUntilSuccess bool
var ecsAttachDiskCmd = &cobra.Command{
	Use:   "attach-disk",
	Short: "This command is used to attach a disk to an ECS.",
	Long:  `This command is used to attach a disk to an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		job, err := ecs.AttachDiskEcs(ProjectID, ecsAttachDiskEcsId, ecsAttachDiskVolumeId, ecsAttachDiskDevice)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		if ecsAttachDiskWaitUntilSuccess {
			_, err := ecs.WaitUntilJobSuccess(ProjectID, job.JobID)
			if err != nil {
				beautyfulPrints.PrintError(err)
				return
			}
			disks, err := ecs.GetListAttachedDisks(ProjectID, ecsAttachDiskEcsId)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(disks, jmesPathQuery)
			}
		} else {
			beautyfulPrints.PrintStruct(job, jmesPathQuery)
		}
	},
}

var ecsDetachDiskVolumeId string
var ecsDetachDiskEcsId string
var ecsDetachDeleteFlag int
var ecsDetachDiskCmd = &cobra.Command{
	Use:   "detach-disk",
	Short: "This command is used to detach a disk from an ECS.",
	Long:  `This command is used to detach a disk from an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.DetachDiskEcs(ProjectID, ecsDetachDiskEcsId, ecsDetachDiskVolumeId, ecsDetachDeleteFlag)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var ecsGetAttachedDisksEcsId string
var ecsGetAttachedDisksCmd = &cobra.Command{
	Use:   "get-attached-disks",
	Short: "This command is used to query information about disks attached to an ECS.",
	Long:  `This command is used to query information about disks attached to an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		disks, err := ecs.GetListAttachedDisks(ProjectID, ecsGetAttachedDisksEcsId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(disks, jmesPathQuery)
		}
	},
}

var ecsGetDiskInfoEcsId string
var ecsGetDiskInfoVolumeId string
var ecsGetDiskInfoCmd = &cobra.Command{
	Use:   "get-disk-info",
	Short: "This command is used to query information about disk attached to an ECS.",
	Long:  `This command is used to query information about disk attached to an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		disks, err := ecs.GetInfoAboutAttachedDisk(ProjectID, ecsGetDiskInfoEcsId, ecsGetDiskInfoVolumeId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(disks, jmesPathQuery)
		}
	},
}

var ecsGetNicsEcsId string
var ecsGetNicsListCmd = &cobra.Command{
	Use:   "get-nics",
	Short: "This command is used to query NICs of an ECS.",
	Long:  `This command is used to query NICs of an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		disks, err := ecs.GetEcsNics(ProjectID, ecsGetNicsEcsId)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(disks, jmesPathQuery)
		}
	},
}

var ecsBindPrivateIpNicId string
var ecsBindPrivateIpSubnetId string
var ecsBindPrivateIpAddress string
var ecsBindPrivateIpReverseBinding bool
var ecsBindPrivateIpCmd = &cobra.Command{
	Use:   "bind-private-ip",
	Short: "This command is used to configure a virtual IP address for an ECS NIC.",
	Long:  `This command is used to configure a virtual IP address for an ECS NIC.`,
	Run: func(cmd *cobra.Command, args []string) {
		ips, err := ecs.BindPrivateIp(ProjectID, ecsBindPrivateIpNicId, ecsBindPrivateIpSubnetId, ecsBindPrivateIpAddress, ecsBindPrivateIpReverseBinding)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ips, jmesPathQuery)
		}
	},
}

var ecsUnbindPrivateIpNicId string
var ecsUnbindPrivateIpCmd = &cobra.Command{
	Use:   "unbind-private-ip",
	Short: "This command is used to configure a virtual IP address for an ECS NIC",
	Long:  `This command is used to configure a virtual IP address for an ECS NIC`,
	Run: func(cmd *cobra.Command, args []string) {
		ips, err := ecs.BindPrivateIp(ProjectID, ecsBindPrivateIpNicId, "", "", false) //API for unbind IP is same as for bind, but all fields must be empty
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(ips, jmesPathQuery)
		}
	},
}

var createKeypairKeyName string
var ecsCreateKeypairCmd = &cobra.Command{
	Use:   "create-keypair",
	Short: "This command is used to create key pair.",
	Long:  `This command is used to create key pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := ecs.CreateKeyPair(ProjectID, createKeypairKeyName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(key, jmesPathQuery)
		}
	},
}

var importKeypairKeyName string
var importKeypairKeyPublicKey string
var ecsImportKeypairCmd = &cobra.Command{
	Use:   "import-keypair",
	Short: "This command is used to import key pair.",
	Long:  `This command is used to import key pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := ecs.ImportKeyPair(ProjectID, importKeypairKeyName, importKeypairKeyPublicKey)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(key, jmesPathQuery)
		}
	},
}

var ecsListKeypairCmd = &cobra.Command{
	Use:   "list-keypair",
	Short: "This command is used to list existing key pairs.",
	Long:  `This command is used to list existing key pairs.`,
	Run: func(cmd *cobra.Command, args []string) {
		keys, err := ecs.ListKeypairs(ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(keys, jmesPathQuery)
		}
	},
}

var getKeypairKeyName string
var ecsGetKeypairCmd = &cobra.Command{
	Use:   "get-keypair",
	Short: "This command is used to query key pair.",
	Long:  `This command is used to query key pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := ecs.GetKeypair(ProjectID, getKeypairKeyName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(key, jmesPathQuery)
		}
	},
}

var deleteKeypairKeyName string
var ecsDeleteKeypairCmd = &cobra.Command{
	Use:   "delete-keypair",
	Short: "This command is used to delete key pair.",
	Long:  `This command is used to delete key pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.DeleteKeypair(ProjectID, deleteKeypairKeyName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var changeFlavorListEcsID string
var changeFlavorListSortKey string
var changeFlavorListSortDir string
var changeFlavorListMarker string
var changeFlavorListLimit int
var ecsChangeFlavorListCmd = &cobra.Command{
	Use:   "change-flavor-list",
	Short: "This command is used to list possible flavors to change of an ECS.",
	Long:  `This command is used to list possible flavors to change of an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		flavors, err := ecs.ListResizeFlavors(ProjectID, changeFlavorListSortKey, changeFlavorListSortDir, changeFlavorListMarker, changeFlavorListEcsID, changeFlavorListLimit)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(flavors, jmesPathQuery)
		}
	},
}

var changeFlavorEcsID string
var changeFlavorFlavorRef string
var changeFlavorDryRun bool
var ecsChangeFlavorCmd = &cobra.Command{
	Use:   "change-flavor",
	Short: "This command is used to change flavor of an ECS.",
	Long:  `This command is used to change flavor of an ECS.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.ResizeECS(ProjectID, changeFlavorEcsID, changeFlavorFlavorRef, changeFlavorDryRun)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var addSGEcsID string
var addSGName string
var ecsAddSGCmd = &cobra.Command{
	Use:   "add-sg",
	Short: "This API is used to add an ECS to a security group.",
	Long:  `This API is used to add an ECS to a security group.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.AddSecurityGroup(ProjectID, addSGEcsID, addSGName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
		}
	},
}

var removeSGEcsID string
var removeSGName string
var ecsDeleteSGCmd = &cobra.Command{
	Use:   "delete-sg",
	Short: "This API is used to delete an ECS from a security group.",
	Long:  `This API is used to delete an ECS from a security group.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ecs.RemoveSecurityGroup(ProjectID, removeSGEcsID, removeSGName)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			fmt.Println("OK")
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
	ecsCmd.AddCommand(ecsBatchAddNicsCmd)
	ecsCmd.AddCommand(ecsBatchDeleteNicsCmd)
	ecsCmd.AddCommand(jobInfoCmd)
	ecsCmd.AddCommand(ecsAttachDiskCmd)
	ecsCmd.AddCommand(ecsDetachDiskCmd)
	ecsCmd.AddCommand(ecsGetAttachedDisksCmd)
	ecsCmd.AddCommand(ecsGetDiskInfoCmd)
	ecsCmd.AddCommand(ecsGetNicsListCmd)
	ecsCmd.AddCommand(ecsBindPrivateIpCmd)
	ecsCmd.AddCommand(ecsUnbindPrivateIpCmd)
	ecsCmd.AddCommand(ecsCreateKeypairCmd)
	ecsCmd.AddCommand(ecsImportKeypairCmd)
	ecsCmd.AddCommand(ecsListKeypairCmd)
	ecsCmd.AddCommand(ecsGetKeypairCmd)
	ecsCmd.AddCommand(ecsDeleteKeypairCmd)
	ecsCmd.AddCommand(ecsChangeFlavorListCmd)
	ecsCmd.AddCommand(ecsChangeFlavorCmd)
	ecsCmd.AddCommand(ecsAddSGCmd)
	ecsCmd.AddCommand(ecsDeleteSGCmd)
	ecsCmd.AddCommand(ecsGetFlavorCmd)

	ecsListCmd.Flags().IntVarP(&ecsListLimit, "limit", "l", 0, "Specifies the maximum number of ECSs on one page.")
	ecsListCmd.Flags().IntVarP(&ecsListOffset, "offset", "o", 0, "Specifies a page number.")

	ecsFlavorListCmd.Flags().StringVarP(&ecsFlavorListAvailabilityZone, "availability_zone", "a", "", "")
	ecsFlavorListCmd.Flags().IntVarP(&ecsFlavorListVcpus, "vcpus", "c", 0, "")
	ecsFlavorListCmd.Flags().IntVarP(&ecsFlavorListRam, "ram", "r", 0, "")
	ecsFlavorListCmd.Flags().StringVarP(&ecsGetFlavorGen, "gen", "g", "", "")

	ecsGetFlavorCmd.Flags().StringVarP(&ecsGetFlavorAvailabilityZone, "availability_zone", "a", "", "")
	ecsGetFlavorCmd.Flags().IntVarP(&ecsGetFlavorVcpus, "vcpus", "c", 0, "")
	ecsGetFlavorCmd.Flags().IntVarP(&ecsGetFlavorRam, "ram", "r", 0, "")
	ecsGetFlavorCmd.Flags().StringVarP(&ecsGetFlavorType, "type", "t", "normal", "Specifies the ECS flavor type: normal: general computing; cpuv1: computing I; cpuv2: computing II; computingv3: general computing-plus; highmem: memory-optimized; saphana: large-memory HANA ECS; diskintensive: disk-intensive")
	ecsGetFlavorCmd.Flags().StringVarP(&ecsGetFlavorGen, "gen", "g", "", "")

	ecsInfoCmd.Flags().StringVarP(&ecsGetInfoId, "id", "i", "", "Specifies the ECS ID.")

	ecsDeleteCmd.Flags().StringSliceVarP(&ecsDeleteIds, "id", "i", nil, "Specifies the ID of the ECS to be deleted.")
	ecsDeleteCmd.Flags().BoolVar(&ecsDeletePublicIp, "del-ip", false, "Specifies whether to delete the EIP bound to the ECS when deleting the ECS. If you do not want to delete the EIP, the system only unbinds the EIP from the ECS and reserves the IP address.")
	ecsDeleteCmd.Flags().BoolVar(&ecsDeleteVolume, "del-vol", false, "Specifies whether to delete the data disks attached to an ECS when deleting the ECS. If you set the parameter value to false, the system only detaches the disks from the ECS and reserves the disks.")
	ecsDeleteCmd.Flags().BoolVar(&ecsDeleteWaitUntilSuccess, "wait-until-success", true, "")

	ecsCreateCmd.Flags().StringVar(&ecsCreateVpcID, "vpc-id", "", "Specifies the ID of the VPC to which the ECS belongs. The value is in the format of the UUID.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateImageRef, "image-ref", "", "Specifies the ID of the system image used for creating ECSs.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateName, "name", "", "Specifies the ECS name.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateFlavorRef, "flavor-ref", "", "Specifies the flavor ID of the ECS to be created.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateRootVolumeType, "root-volume-type", "SAS", "Specifies the ECS system disk type, which must be one of available disk types.")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSubnetIds, "subnet-ids", nil, "Specifies the subnets of the ECS.")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSecGroupIds, "sg-ids", nil, "Specifies the security groups of the ECS.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateAdminPass, "admin-pass", "", "Specifies the initial login password of the administrator account for logging in to an ECS using password authentication")
	ecsCreateCmd.Flags().StringVar(&ecsCreateEipId, "eip-id", "", "Specifies the EIP ID")
	ecsCreateCmd.Flags().IntVar(&ecsCreateEipBandwidthSize, "eip-size", 0, "Specifies the bandwidth size. Specifies the bandwidth (Mbit/s). The value ranges from 1 to 300.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateEipBandwidthType, "eip-bandwidth", "PER", "Specifies the bandwidth sharing type. Enumerated values: PER (indicates exclusive bandwidth) and WHOLE (indicates sharing)")
	ecsCreateCmd.Flags().StringVar(&ecsCreateEipType, "eip-type", "5_bgp", "Specifies Type of EIP. The value can be 5_bgp, default is 5_bgp")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateVolumeTypes, "data-volume-types", nil, "Specifies the type of the ECS data disk, which must be one of available disk types.")
	ecsCreateCmd.Flags().IntSliceVar(&ecsCreateVolumeSizes, "data-volume-sizes", nil, "Specifies the data disk size, in GB. The value ranges from 10 to 32768.")
	ecsCreateCmd.Flags().IntVar(&ecsCreateRootVolumeSize, "root-volume-size", 40, "Specifies the system disk size, in GB. The value ranges from 1 to 1024.")
	ecsCreateCmd.Flags().IntVar(&ecsCreateCount, "count", 1, "Specifies the number of ECSs to be created.")
	ecsCreateCmd.Flags().StringVar(&ecsCreateAvailabilityZone, "az", "ru-moscow-1a", "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateKeyName, "key-name", "", "Specifies the name of the SSH key used for logging in to the ECS.")
	ecsCreateCmd.Flags().IntVar(&ecsCreateVcpus, "vcpus", 0, "")
	ecsCreateCmd.Flags().IntVar(&ecsCreateRam, "ram", 0, "")
	ecsCreateCmd.Flags().StringVarP(&ecsCreateFlavorType, "flavor-type", "t", "normal", "Specifies the ECS flavor type: normal: general computing; cpuv1: computing I; cpuv2: computing II; computingv3: general computing-plus; highmem: memory-optimized; saphana: large-memory HANA ECS; diskintensive: disk-intensive")
	ecsCreateCmd.Flags().StringVar(&ecsCreateFlavorGen, "flavor-gen", "", "")
	ecsCreateCmd.Flags().BoolVar(&ecsCreateAssignEip, "assign-eip", false, "")
	ecsCreateCmd.Flags().BoolVar(&ecsCreateWaitUntilSuccess, "wait-until-success", true, "")
	ecsCreateCmd.Flags().StringVar(&ecsCreateVpcName, "vpc-name", "", "")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSubnetNames, "subnet-names", nil, "Specifies the subnets of the ECS.")
	ecsCreateCmd.Flags().StringSliceVar(&ecsCreateSGNames, "sg-names", nil, "Specifies the security groups of the ECS.")

	ecsBatchStartCmd.Flags().StringSliceVarP(&ecsBatchStartServerIds, "id", "i", nil, "Specifies ECS IDs")
	ecsBatchStartCmd.Flags().BoolVar(&ecsStartWaitUntilSuccess, "wait-until-success", true, "")

	ecsBatchRestartCmd.Flags().StringSliceVarP(&ecsBatchRestartServerIds, "id", "i", nil, "Specifies ECS IDs")
	ecsBatchRestartCmd.Flags().StringVarP(&ecsBatchRestartType, "type", "t", "SOFT", "Specifies the type of the restart operation.")
	ecsBatchRestartCmd.Flags().BoolVar(&ecsRestartWaitUntilSuccess, "wait-until-success", true, "")

	ecsBatchStopCmd.Flags().StringSliceVarP(&ecsBatchStopServerIds, "id", "i", nil, "Specifies ECS IDs")
	ecsBatchStopCmd.Flags().StringVarP(&ecsBatchStopType, "type", "t", "SOFT", "Specifies an ECS stop type.")
	ecsBatchStopCmd.Flags().BoolVar(&ecsStopWaitUntilSuccess, "wait-until-success", true, "")

	ecsBatchAddNicsCmd.Flags().StringVarP(&ecsBatchAddNicsServerId, "id", "i", "", "Specifies ECS ID")
	ecsBatchAddNicsCmd.Flags().StringSliceVarP(&ecsBatchAddNicsSubnetIds, "subnet-ids", "s", nil, "Specifies subnet IDs")

	ecsBatchDeleteNicsCmd.Flags().StringVarP(&ecsBatchDeleteNicsServerId, "id", "i", "", "Specifies ECS ID")
	ecsBatchDeleteNicsCmd.Flags().StringSliceVarP(&ecsBatchDeleteNicsSubnetIds, "subnet-ids", "s", nil, "Specifies subnet IDs")

	jobInfoCmd.Flags().StringVarP(&ecsJobId, "id", "i", "", "")

	ecsAttachDiskCmd.Flags().StringVar(&ecsAttachDiskEcsId, "ecs-id", "", "Specifies ECS ID")
	ecsAttachDiskCmd.Flags().StringVar(&ecsAttachDiskVolumeId, "vol-id", "", "Specifies volume ID")
	ecsAttachDiskCmd.Flags().StringVar(&ecsAttachDiskDevice, "device", "", "Specifies device")
	ecsAttachDiskCmd.Flags().BoolVar(&ecsAttachDiskWaitUntilSuccess, "wait-until-success", true, "")

	ecsDetachDiskCmd.Flags().StringVarP(&ecsDetachDiskEcsId, "ecs-id", "e", "", "Specifies ECS ID")
	ecsDetachDiskCmd.Flags().StringVarP(&ecsDetachDiskVolumeId, "vol-id", "v", "", "Specifies volume ID")
	ecsDetachDiskCmd.Flags().IntVarP(&ecsDetachDeleteFlag, "delete-flag", "d", 0, "Indicates whether to forcibly detach a data disk.")

	ecsGetAttachedDisksCmd.Flags().StringVarP(&ecsGetAttachedDisksEcsId, "ecs-id", "e", "", "Specifies ECS ID")

	ecsGetDiskInfoCmd.Flags().StringVarP(&ecsGetDiskInfoEcsId, "ecs-id", "e", "", "Specifies ECS ID")
	ecsGetDiskInfoCmd.Flags().StringVarP(&ecsGetDiskInfoVolumeId, "vol-id", "v", "", "Specifies volume ID")

	ecsGetNicsListCmd.Flags().StringVarP(&ecsGetNicsEcsId, "ecs-id", "e", "", "Specifies ECS ID")

	ecsBindPrivateIpCmd.Flags().StringVarP(&ecsBindPrivateIpNicId, "nic-id", "n", "", "")
	ecsBindPrivateIpCmd.Flags().StringVarP(&ecsBindPrivateIpSubnetId, "subnet-id", "s", "", "")
	ecsBindPrivateIpCmd.Flags().StringVarP(&ecsBindPrivateIpAddress, "ip", "i", "", "")
	ecsBindPrivateIpCmd.Flags().BoolVarP(&ecsBindPrivateIpReverseBinding, "reverse-binding", "r", false, "")

	ecsUnbindPrivateIpCmd.Flags().StringVarP(&ecsUnbindPrivateIpNicId, "nic-id", "n", "", "")

	ecsCreateKeypairCmd.Flags().StringVarP(&createKeypairKeyName, "name", "n", "", "Specifies the name of key pair")

	ecsImportKeypairCmd.Flags().StringVarP(&importKeypairKeyName, "name", "n", "", "")
	ecsImportKeypairCmd.Flags().StringVarP(&importKeypairKeyPublicKey, "public-key", "p", "", "")

	ecsGetKeypairCmd.Flags().StringVarP(&getKeypairKeyName, "name", "n", "", "")

	ecsDeleteKeypairCmd.Flags().StringVarP(&deleteKeypairKeyName, "name", "n", "", "")

	ecsChangeFlavorListCmd.Flags().StringVar(&changeFlavorListEcsID, "ecs-id", "", "")
	ecsChangeFlavorListCmd.Flags().StringVar(&changeFlavorListSortKey, "sort-key", "", "")
	ecsChangeFlavorListCmd.Flags().StringVar(&changeFlavorListSortDir, "sort-dir", "", "")
	ecsChangeFlavorListCmd.Flags().StringVar(&changeFlavorListMarker, "marker", "", "")
	ecsChangeFlavorListCmd.Flags().IntVar(&changeFlavorListLimit, "limit", 0, "")

	ecsChangeFlavorCmd.Flags().StringVarP(&changeFlavorEcsID, "ecs-id", "e", "", "")
	ecsChangeFlavorCmd.Flags().StringVarP(&changeFlavorFlavorRef, "flavor-id", "f", "", "")
	ecsChangeFlavorCmd.Flags().BoolVarP(&changeFlavorDryRun, "dry-run", "d", false, "")

	ecsAddSGCmd.Flags().StringVarP(&addSGEcsID, "ecs-id", "e", "", "Specifies the ECS ID.")
	ecsAddSGCmd.Flags().StringVarP(&addSGName, "sg-name", "s", "", "Specifies the UUID or name of the security group to which the ECS is added. The configuration takes effect for the NICs on the ECS.")

	ecsDeleteSGCmd.Flags().StringVarP(&removeSGEcsID, "ecs-id", "e", "", "Specifies the ECS ID.")
	ecsDeleteSGCmd.Flags().StringVarP(&removeSGName, "sg-name", "s", "", "Specifies the UUID or name of the security group from which the ECS is removed. The configuration takes effect for the NICs on the ECS.")
}
