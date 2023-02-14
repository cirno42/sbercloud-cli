package cmd

import (
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/quotaModels"
	"sbercloud-cli/api/quota"
	"sbercloud-cli/internal/beautyfulPrints"
)

var quotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "commands to interact with service quotas",
	Long:  `commands to interact with service quotas`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var quotaInfoType string
var quotaInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "commands to interact with service quotas",
	Long:  `commands to interact with service quotas`,
	Run: func(cmd *cobra.Command, args []string) {
		var q []quotaModels.QuotaModel
		var err error
		if quotaInfoType == "ecs" {
			q, err = quota.GetInfoAboutServerQuota(ProjectID)
		} else if quotaInfoType != "" {
			q, err = quota.GetInfoAboutNetworkQuota(ProjectID, quotaInfoType)
		} else {
			q1, err := quota.GetInfoAboutNetworkQuota(ProjectID, quotaInfoType)
			if err == nil {
				var q2 []quotaModels.QuotaModel
				q2, err = quota.GetInfoAboutServerQuota(ProjectID)
				q = append(q1, q2...)
			}
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(q, jmesPathQuery)
		}
	},
}

func init() {
	RootCmd.AddCommand(quotaCmd)
	quotaCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	quotaCmd.AddCommand(quotaInfoCmd)

	quotaInfoCmd.Flags().StringVarP(&quotaInfoType, "type", "t", "", "Specifies the resource type.\n\nThe value can be vpc, subnet, securityGroup, securityGroupRule, publicIp, vpn, vpngw, vpcPeer, shareBandwidth, shareBandwidthIP, ecs")
}
