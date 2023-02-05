package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var setOutputFormat string
var setSecretKey string
var setAccessKey string
var setProjectId string
var setRegion string
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure application",
	Long:  `configure application, set auth, output, endpoint parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		if setOutputFormat == "" && setSecretKey == "" && setAccessKey == "" && setProjectId == "" {
			setConfigManual()
		} else {
			sk := os.Getenv("SECRET_KEY")
			if setSecretKey != "" {
				sk = setSecretKey
			}
			ak := os.Getenv("ACCESS_KEY")
			if setAccessKey != "" {
				ak = setAccessKey
			}
			output := os.Getenv("OUTPUT_FORMAT")
			if setOutputFormat != "" {
				output = setOutputFormat
			}
			projectId := os.Getenv("PROJECT_ID")
			if setProjectId != "" {
				projectId = setProjectId
			}
			region := os.Getenv("REGION")
			if setRegion != "" {
				region = setRegion
			}
			config := make(map[string]string, 5)
			config["ACCESS_KEY"] = ak
			config["SECRET_KEY"] = sk
			config["PROJECT_ID"] = projectId
			config["OUTPUT_FORMAT"] = output
			config["REGION"] = region
			err := godotenv.Write(config, ".env")
			if err != nil {
				fmt.Println("Error while writing file: ", err)
				return
			}
		}
	},
}

func setConfigManual() {
	keys := []string{"ACCESS_KEY", "SECRET_KEY", "PROJECT_ID", "OUTPUT_FORMAT", "REGION"}
	config := make(map[string]string, len(keys))
	var value string
	if runtime.GOOS == "windows" {
		fmt.Println("Available values for OUTPUT_FORMAT: YAML/JSON/TABLE")
	} else if runtime.GOOS == "linux" {
		fmt.Println("Available values for OUTPUT_FORMAT: YAML/JSON/JSON-C/TABLE")
	}
	fmt.Println("Available values for OUTPUT_FORMAT: YAML/JSON/TABLE")
	fmt.Println("Available values for REGION: Ru-Moscow")
	for _, key := range keys {
		fmt.Print(key, "=")
		_, err := fmt.Scanln(&value)
		if err != nil {
			fmt.Println("Error while reading stdin: ", err)
			return
		}
		config[key] = value
	}
	err := godotenv.Write(config, ".env")
	if err != nil {
		fmt.Println("Error while writing file: ", err)
		return
	}
}

func editConfig() {

}

func init() {
	RootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVar(&setOutputFormat, "set-output-format", "", "")
	configureCmd.Flags().StringVar(&setSecretKey, "set-secret-key", "", "")
	configureCmd.Flags().StringVar(&setAccessKey, "set-access-key", "", "")
	configureCmd.Flags().StringVar(&setProjectId, "set-project-id", "", "")
	configureCmd.Flags().StringVar(&setRegion, "set-region", "", "")
}
