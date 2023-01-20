package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"runtime"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure application",
	Long:  `configure application, set auth, output, endpoint parameters`,
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	RootCmd.AddCommand(configureCmd)
}
