package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cirno42/sbercloud-api/api/eip"
	"github.com/cirno42/sbercloud-api/api/models/dumpModels"
	"github.com/cirno42/sbercloud-api/api/nat"
	"github.com/cirno42/sbercloud-api/api/subnets"
	"github.com/cirno42/sbercloud-api/api/vpcs"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore infrastructure from specified file",
	Long:  `Restore infrastructure from specified file`,
	Run: func(cmd *cobra.Command, args []string) {
		var dump dumpModels.DumpModel
		dumpFile, err := os.Open(restoreFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dumpFile.Close()
		dumpBytes, _ := ioutil.ReadAll(dumpFile)
		err = json.Unmarshal(dumpBytes, &dump)
		vpcKeysMapping := make(map[string]string)
		subnetKeysMapping := make(map[string]string)
		oldVpcs := dump.Vpcs
		for _, oldVpc := range oldVpcs {
			newVpc, err := vpcs.CreateVpc(ProjectID, oldVpc.Name, oldVpc.Description, oldVpc.Cidr)
			if err != nil {
				fmt.Println(err)
				return
			}
			vpcKeysMapping[oldVpc.Id] = newVpc.Id
		}
		oldSubnets := dump.Subnets
		for _, oldSubnet := range oldSubnets {
			newSubnet, err := subnets.CreateSubnet(ProjectID, oldSubnet.Name, oldSubnet.Description,
				oldSubnet.Cidr, oldSubnet.GatewayIp, oldSubnet.Ipv6Enable, oldSubnet.DhcpEnable, oldSubnet.PrimaryDns, oldSubnet.SecondaryDns,
				oldSubnet.DnsList, oldSubnet.AvailabilityZone, vpcKeysMapping[oldSubnet.VpcId])
			if err != nil {
				fmt.Println(err)
				return
			}
			subnetKeysMapping[oldSubnet.Id] = newSubnet.Id
		}

		oldEips := dump.Eips
		for _, oldEip := range oldEips {
			_, err := eip.AssignEIP(ProjectID, oldEip.EipType, oldEip.IPVersion, oldEip.BandwidthName, oldEip.BandwidthSize, oldEip.BandwidthShareType)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		oldNats := dump.Nats
		for _, oldNat := range oldNats {
			_, err := nat.CreateNAT(ProjectID, oldNat.Name, oldNat.Description, vpcKeysMapping[oldNat.RouterID], subnetKeysMapping[oldNat.InternalNetworkID], oldNat.Spec, oldNat.EnterpriseProjectID)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

var restoreFileName string

func init() {
	RootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVarP(&restoreFileName, "file", "f", "", "File with data to restore")
}
