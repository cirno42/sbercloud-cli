package cmd

import "github.com/spf13/cobra"

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "commands to interact with Availability Zones",
	Long:  `commands to interact with Availability Zones`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(azCmd)

}
