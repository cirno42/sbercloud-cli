package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"sbercloud-cli/api/models/dumpModels"
	"sbercloud-cli/api/subnets"
	"sbercloud-cli/api/vpcs"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Long: `add
details
here`,
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
		oldVpcs := dump.Vpcs
		for _, oldVpc := range oldVpcs {
			newVpc, err := vpcs.CreateVpc(ProjectID, oldVpc.Name, oldVpc.Description, oldVpc.Cidr)
			if err != nil {
				fmt.Println(err)
				return
			}
			vpcKeysMapping[oldVpc.Id] = newVpc.Vpc.Id
		}
		oldSubnets := dump.Subnets
		for _, oldSubnet := range oldSubnets {
			_, err := subnets.CreateSubnet(ProjectID, oldSubnet.Name, oldSubnet.Description,
				oldSubnet.Cidr, oldSubnet.GatewayIp, oldSubnet.Ipv6Enable, oldSubnet.DhcpEnable, oldSubnet.PrimaryDns, oldSubnet.SecondaryDns,
				oldSubnet.DnsList, oldSubnet.AvailabilityZone, vpcKeysMapping[oldSubnet.VpcId])
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
