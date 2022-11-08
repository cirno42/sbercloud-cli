package beautyfulPrints

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
	"strings"
)

func PrintError(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
}

//todo: add check for print type (json/yaml/json-c)
func PrintStruct(s interface{}, jmesPathQuery string) {
	if len(strings.TrimSpace(jmesPathQuery)) > 0 {
		printStructUsingQuery(s, jmesPathQuery)
		return
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}

func printStructUsingQuery(s interface{}, query string) {
	queryResult, err := jmespath.Search(query, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(queryResult)
}
