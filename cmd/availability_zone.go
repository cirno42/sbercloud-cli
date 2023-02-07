package cmd

import (
	"github.com/spf13/cobra"
	"sbercloud-cli/api/availabilityZone"
	"sbercloud-cli/internal/beautyfulPrints"
)

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "commands to interact with Availability Zones",
	Long:  `commands to interact with Availability Zones`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var azGetZonesCmd = &cobra.Command{
	Use:   "get-zones",
	Short: "This command is used to query AZs.",
	Long:  `This command is used to query AZs.`,
	Run: func(cmd *cobra.Command, args []string) {
		zones, err := availabilityZone.GetZonesList(ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(zones, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(azCmd)

	azCmd.AddCommand(azGetZonesCmd)
}
