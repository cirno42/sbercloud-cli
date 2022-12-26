package beautyfulPrints

import (
	"encoding/json"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"github.com/jmespath/go-jmespath"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func PrintError(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
}

func PrintStruct(s interface{}, jmesPathQuery string) {
	if len(strings.TrimSpace(jmesPathQuery)) > 0 {
		printStructUsingQuery(s, jmesPathQuery)
		return
	}
	format := os.Getenv("OUTPUT_FORMAT")
	if format == "YAML" {
		printStructAsYaml(s)
	} else if format == "JSON" {
		printStructAsJson(s)
	} else if format == "JSON-C" {
		printStructAsJsonColor(s)
	} else {
		fmt.Println("Output format is not configured!")
	}
}

func printStructAsYaml(s interface{}) {
	output, err := yaml.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(output))
}

func printStructAsJson(s interface{}) {
	output, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(output))
}

func printStructAsJsonColor(s interface{}) {
	output, err := prettyjson.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(output))
}

func PrintStructToFile(s interface{}, file *os.File) {
	output, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(string(output))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printStructUsingQuery(s interface{}, query string) {
	queryResult, err := jmespath.Search(query, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	printStructAsJson(queryResult)
}
