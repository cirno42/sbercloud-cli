package cmd

import (
	"github.com/cirno42/sbercloud-api/api/iam"
	"github.com/spf13/cobra"
	"sbercloud-cli/internal/beautyfulPrints"
)

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "commands to interact with Identity and Access Management Service",
	Long:  `commands to interact with Identity and Access Management Service`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var iamListProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "This command is used to query Projects.",
	Long:  `This command is used to query Projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := iam.ListProjects()
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(projects, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(iamCmd)
	iamCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	iamCmd.AddCommand(iamListProjectsCmd)
}
