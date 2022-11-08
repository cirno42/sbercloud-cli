package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/natModels"
	"sbercloud-cli/api/nat"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var natCmd = &cobra.Command{
	Use:   "nat",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nat called")
	},
}

var natCreateName string
var natCreateDesc string
var natCreateRouterID string
var natCreateInternalNetworkID string
var natCreateSpec string
var natCreateEnterpriseProjectID string
var natCreateRouterName string
var natCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var createdNat *natModels.NatModel
		var err error
		var vpcID string
		if natCreateRouterID != "" {
			vpcID = natCreateRouterID
		} else if natCreateRouterName != "" {
			vpc, err := vpcs.GetVpcByName(ProjectID, natCreateRouterName)
			if err != nil {
				vpcID = vpc.Id
			}
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			createdNat, err = nat.CreateNAT(ProjectID, natCreateName, natCreateDesc, vpcID, natCreateInternalNetworkID, natCreateSpec, natCreateEnterpriseProjectID)
			if err != nil {
				beautyfulPrints.PrintError(err)
			} else {
				beautyfulPrints.PrintStruct(createdNat, jmesPathQuery)
			}
		}

	},
}

var natListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		nats, err := nat.GetNatList(ProjectID)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(nats, jmesPathQuery)
		}
	},
}

var natGetInfoNatID string
var natGetInfoNatName string
var natGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var natEntity *natModels.NatModel
		if natGetInfoNatID != "" {
			natEntity, err = nat.GetInfoAboutNat(ProjectID, natGetInfoNatID)
		} else if natGetInfoNatName != "" {
			natEntity, err = nat.GetNatByName(ProjectID, natGetInfoNatName)
		} else {
			err = errors.New("NAT name and id both are not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(natEntity, jmesPathQuery)
		}
	},
}

var natDeleteNatID string
var natDeleteNatName string
var natDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if natDeleteNatID != "" {
			err = nat.DeleteNat(ProjectID, natDeleteNatID)
		} else if natDeleteNatName != "" {
			err = nat.DeleteNatByName(ProjectID, natDeleteNatName)
		} else {
			err = errors.New("NAT name and id both are not specified")
		}
		if err != nil {
			beautyfulPrints.PrintError(err)
		}

	},
}

func init() {
	RootCmd.AddCommand(natCmd)
	natCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	natCmd.AddCommand(natCreateCmd)
	natCmd.AddCommand(natListCmd)
	natCmd.AddCommand(natGetInfoCmd)
	natCmd.AddCommand(natDeleteCmd)

	natCreateCmd.Flags().StringVarP(&natCreateName, "name", "n", "", "")
	natCreateCmd.Flags().StringVarP(&natCreateDesc, "description", "d", "", "")
	natCreateCmd.Flags().StringVarP(&natCreateRouterID, "router-id", "i", "", "")
	natCreateCmd.Flags().StringVar(&natCreateInternalNetworkID, "network-id", "", "")
	natCreateCmd.Flags().StringVarP(&natCreateSpec, "spec", "s", "1", "")
	natCreateCmd.Flags().StringVarP(&natCreateEnterpriseProjectID, "ent-project-id", "p", "", "")
	natCreateCmd.Flags().StringVarP(&natCreateRouterName, "router-name", "r", "", "")

	natGetInfoCmd.Flags().StringVarP(&natGetInfoNatID, "id", "i", "", "")
	natGetInfoCmd.Flags().StringVarP(&natGetInfoNatName, "name", "n", "", "")

	natDeleteCmd.Flags().StringVarP(&natDeleteNatID, "id", "i", "", "")
	natDeleteCmd.Flags().StringVarP(&natDeleteNatName, "name", "n", "", "")
}
