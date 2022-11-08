/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/models/vpcModels"
	"sbercloud-cli/api/vpcs"
	"sbercloud-cli/internal/beautyfulPrints"
)

var listMarker string
var listLimit int

var createName string
var createCidr string
var createDescription string

var infoName string
var infoVpcID string

var deleteName string
var deleteID string
var vpcDeleteIsRecursive bool

// vpcCmd represents the vpc command
var vpcCmd = &cobra.Command{
	Use:   "vpc",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpc called")
	},
}

var vpcCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpc create called")
		vpc, err := vpcs.CreateVpc(ProjectID, createName, createDescription, createCidr)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(vpc, jmesPathQuery)
		}
	},
}

var vpcGetListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpc list called")
		vpcs, err := vpcs.GetVpcsList(ProjectID, listLimit, listMarker)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		beautyfulPrints.PrintStruct(vpcs, jmesPathQuery)
	},
}

var vpcGetInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vpc info called")
		var vpc *vpcModels.VpcModel
		var err error
		if infoVpcID != "" {
			vpc, err = vpcs.GetInfoAboutVpc(ProjectID, infoVpcID)
		} else {
			vpc, err = vpcs.GetVpcByName(ProjectID, infoName)
		}
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		} else {
			beautyfulPrints.PrintStruct(*vpc, jmesPathQuery)
		}
	},
}

var vpcDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if vpcDeleteIsRecursive {
			var vpcID string
			if deleteID == "" {
				vpc, err := vpcs.GetVpcByName(ProjectID, deleteName)
				if err != nil {
					fmt.Println("ERROR: " + err.Error())
					return
				}
				vpcID = vpc.Id
			} else {
				vpcID = deleteID
			}
			err := vpcs.DeleteVpcRecursive(ProjectID, vpcID)
			if err != nil {
				fmt.Println("ERROR: " + err.Error())

			}
			return
		}
		var id string
		if deleteID != "" {
			id = deleteID
		} else {
			vpc, err := vpcs.GetVpcByName(ProjectID, deleteName)
			if err != nil {
				fmt.Printf("ERROR: %s\n", err.Error())
				return
			}
			id = vpc.Id
		}
		err := vpcs.DeleteVpc(ProjectID, id)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(vpcCmd)

	vpcCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	vpcCmd.AddCommand(vpcCreateCmd)
	vpcCreateCmd.Flags().StringVarP(&createName, "name", "n", "", "add details here")
	vpcCreateCmd.Flags().StringVarP(&createCidr, "cidr", "c", "", "add details here")
	vpcCreateCmd.Flags().StringVarP(&createDescription, "description", "d", "", "add details here")

	vpcCmd.AddCommand(vpcGetListCmd)
	vpcGetListCmd.Flags().IntVarP(&listLimit, "limit", "l", 0, "add details here")
	vpcGetListCmd.Flags().StringVarP(&listMarker, "marker", "m", "", "add details here")

	vpcCmd.AddCommand(vpcGetInfoCmd)
	vpcGetInfoCmd.Flags().StringVarP(&infoName, "name", "n", "", "add details here")
	vpcGetInfoCmd.Flags().StringVarP(&infoVpcID, "id", "i", "", "add details here")

	vpcCmd.AddCommand(vpcDeleteCmd)
	vpcDeleteCmd.Flags().StringVarP(&deleteName, "name", "n", "", "add details here")
	vpcDeleteCmd.Flags().StringVarP(&deleteID, "id", "i", "", "add details here")
	vpcDeleteCmd.Flags().BoolVarP(&vpcDeleteIsRecursive, "rec", "r", false, "add details here")
}

//todo: Add "vpc-" prefix to every flag variable
