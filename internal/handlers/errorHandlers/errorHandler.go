package errorHandlers

import "fmt"

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type NeutronError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Detail  string `json:"detail"`
}

func HandleError(errorResp ErrorResponse) {
	fmt.Printf("%v+\n", errorResp)
}
