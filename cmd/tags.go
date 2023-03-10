package cmd

import (
	"github.com/spf13/cobra"
	"sbercloud-cli/api/eip"
	"sbercloud-cli/api/subnets"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var getByTagKeys []string
var getByTagValues []string
var getByTagCmd = &cobra.Command{
	Use:   "get-by-tags",
	Short: "Command to count and filter resources by tags",
	Long:  `Command to count and filter resources by tags`,
	Run: func(cmd *cobra.Command, args []string) {
		vpc, err := vpcs.FilterVpcByTags(ProjectID, getByTagKeys, getByTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		subnet, err := subnets.FilterSubnetByTags(ProjectID, getByTagKeys, getByTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		eips, err := eip.FilterEipByTags(ProjectID, getByTagKeys, getByTagValues)
		if err != nil {
			beautyfulPrints.PrintError(err)
			return
		}
		s1 := append(vpc, subnet...)
		s2 := append(s1, eips...)
		beautyfulPrints.PrintStruct(s2, "")
	},
}

func init() {
	RootCmd.AddCommand(getByTagCmd)

	getByTagCmd.Flags().StringSliceVarP(&getByTagKeys, "keys", "k", nil, "")
	getByTagCmd.Flags().StringSliceVarP(&getByTagValues, "values", "v", nil, "")
	getByTagCmd.Flags().StringVarP(&jmesPathQuery, "query", "q", "", "")
}
