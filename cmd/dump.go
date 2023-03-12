package cmd

import (
	"fmt"
	"github.com/cirno42/sbercloud-api/api/ecs"
	"github.com/cirno42/sbercloud-api/api/eip"
	"github.com/cirno42/sbercloud-api/api/evs"
	"github.com/cirno42/sbercloud-api/api/models/dumpModels"
	"github.com/cirno42/sbercloud-api/api/nat"
	"github.com/cirno42/sbercloud-api/api/securityGroup"
	"github.com/cirno42/sbercloud-api/api/subnets"
	"github.com/cirno42/sbercloud-api/api/vpcs"
	"github.com/spf13/cobra"
	"os"
	"sbercloud-cli/internal/beautyfulPrints"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump infrastructure to specified file",
	Long:  `dump infrastructure to specified file in JSON format`,
	Run: func(cmd *cobra.Command, args []string) {
		vpcs, err := vpcs.GetVpcsList(ProjectID, 0, "")
		if err != nil {
			fmt.Print(err)
			return
		}
		subnets, err := subnets.GetSubnetsList(ProjectID, 0, "", "")
		if err != nil {
			fmt.Print(err)
			return
		}
		secGroups, err := securityGroup.GetSecurityGroupsList(ProjectID, 0, "", "")
		nats, err := nat.GetNatList(ProjectID)
		if err != nil {
			fmt.Print(err)
			return
		}
		eips, err := eip.GetEIPsList(ProjectID, 1000, "")
		if err != nil {
			fmt.Print(err)
			return
		}
		if dumpFileName == "" {
			dumpFileName = "dump.json"
		}
		snatRules, err := nat.ListSNATRules(ProjectID, "", "", 0)
		if err != nil {
			fmt.Print(err)
			return
		}
		dnatRules, err := nat.ListDNATRules(ProjectID, "", "", "", "", "", "", "", 0, 0)
		if err != nil {
			fmt.Print(err)
			return
		}
		servers, err := ecs.GetECSList(ProjectID, 0, 0, "")
		if err != nil {
			fmt.Print(err)
			return
		}
		disks, err := evs.GetDisksList(ProjectID, "", 0, 0)
		if err != nil {
			fmt.Print(err)
			return
		}
		keypairs, err := ecs.ListKeypairs(ProjectID)
		dump := dumpModels.DumpModel{
			Vpcs:      vpcs,
			Subnets:   subnets,
			Eips:      eips,
			Nats:      nats,
			SecGroups: secGroups,
			SnatRules: snatRules,
			DnatRules: dnatRules,
			ECSs:      servers,
			Disks:     disks,
			KeyPairs:  keypairs,
		}
		outputFile, err := os.Create(dumpFileName)
		defer outputFile.Close()
		if err != nil {
			fmt.Print(err)
			return
		}
		beautyfulPrints.PrintStructToFile(dump, outputFile)
	},
}

var dumpFileName string

func init() {
	RootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().StringVarP(&dumpFileName, "file", "f", "", "File for dump")
}
