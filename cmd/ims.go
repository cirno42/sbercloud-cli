package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sbercloud-cli/api/ims"
	"sbercloud-cli/internal/beautyfulPrints"
)

var imsCmd = &cobra.Command{
	Use:   "ims",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ims called")
	},
}

var imsGetListOfImagesPlatform string
var imsGetListOfImagesCmd = &cobra.Command{
	Use:   "images-list",
	Short: "A brief description of your command",
	Long: `add
details
here`,
	Run: func(cmd *cobra.Command, args []string) {
		images, err := ims.GetImagesList(imsGetListOfImagesPlatform)
		if err != nil {
			beautyfulPrints.PrintError(err)
		} else {
			beautyfulPrints.PrintStruct(images, jmesPathQuery)
			//println(images[0].ID)
		}
	},
}

func init() {
	RootCmd.AddCommand(imsCmd)
	imsCmd.PersistentFlags().StringVarP(&jmesPathQuery, "query", "q", "", "JMES Path query")

	imsCmd.AddCommand(imsGetListOfImagesCmd)
	imsGetListOfImagesCmd.Flags().StringVar(&imsGetListOfImagesPlatform, "platform", "", "")
}
