package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/ecs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("esc called")
	},
}

var ecsFlavorListAvailabilityZone string
var ecsFlavorListCmd = &cobra.Command{
	Use:   "flavor-list",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		flavors, err := ecs.GetESCFlavorList(ProjectID, ecsFlavorListAvailabilityZone)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(flavors, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(ecsCmd)
	ecsCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	ecsCmd.AddCommand(ecsFlavorListCmd)

	ecsFlavorListCmd.Flags().StringVarP(&ecsFlavorListAvailabilityZone, "availability_zone", "a", "", "")
}
