package cmd

import (
	"github.com/spf13/cobra"
	"sbercloud-cli/api/ims"
	"sbercloud-cli/internal/beautyfulPrints"
)

var imsCmd = &cobra.Command{
	Use:   "ims",
	Short: "Commands to interact with IMS",
	Long:  `Commands to interact with IMS`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var imsGetListOfImagesPlatform string
var imsGetListOfImagesCmd = &cobra.Command{
	Use:   "images-list",
	Short: "Get list of images",
	Long:  `Get list of images`,
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
	imsGetListOfImagesCmd.Flags().StringVar(&imsGetListOfImagesPlatform, "platform", "", "Specifies the image platform type. The value can be Windows, Ubuntu, RedHat, SUSE, CentOS, Debian, OpenSUSE, Oracle Linux, Fedora, Other, CoreOS, or EulerOS.")
}
