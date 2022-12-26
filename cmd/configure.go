package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure application",
	Long:  `configure application, set auth, output, endpoint parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		keys := []string{"ACCESS_KEY", "SECRET_KEY", "PROJECT_ID", "OUTPUT_FORMAT"}
		config := make(map[string]string, len(keys))
		var value string
		fmt.Println("Available values for OUTPUT_FORMAT: YAML/JSON")
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
