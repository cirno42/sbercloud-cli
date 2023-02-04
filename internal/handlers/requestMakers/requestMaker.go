package requestMakers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sbercloud-cli/api/auth/authKeys"
	"sbercloud-cli/api/auth/core"
)

type httpMethodType int

const (
	HTTP_METHOD_GET httpMethodType = iota
	HTTP_METHOD_POST
	HTTP_METHOD_DELETE
	HTTP_METHOD_PUT
)

func makeRESTRequest(URL string, method string, body *interface{}) (*http.Request, error) {

	var requestBody []byte
	var err error
	requestBody = nil
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	keys := authKeys.GetAkSkKeys()
	signer := core.Signer{
		Key:    keys.AccessKey,
		Secret: keys.SecretKey,
	}

	req, err := http.NewRequest(method, URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("content-type", "application/json")
	}

	err = signer.Sign(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func doRequest(request *http.Request) (*http.Response, error) {
	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func parseResponseBody(response *http.Response, parsedBody *interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(body))
	//fmt.Println()
	err = json.Unmarshal([]byte(body), parsedBody)
	return err
}

func isHttpRequestSuccess(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode <= 299
}

func httpMethodToString(methodType httpMethodType) string {
	switch methodType {
	case HTTP_METHOD_GET:
		return "GET"
	case HTTP_METHOD_DELETE:
		return "DELETE"
	case HTTP_METHOD_POST:
		return "POST"
	case HTTP_METHOD_PUT:
		return "PUT"
	}
	return ""
}

func CreateAndDoRequest(URL string, method httpMethodType, requestBody interface{},
	parsedResponseBodyPointer interface{}, parsedErrorPointer interface{}) error {
	methodString := httpMethodToString(method)
	if methodString == "" {
		return errors.New("Wrong http method")
	}
	req, err := makeRESTRequest(URL, methodString, &requestBody)
	if err != nil {
		return err
	}
	resp, err := doRequest(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	if err != nil {
		return err
	}
	if isHttpRequestSuccess(resp) {
		if parsedResponseBodyPointer != nil && resp.StatusCode != 204 {
			err = parseResponseBody(resp, &parsedResponseBodyPointer)
		}
		return err
	} else {
		errorData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(errorData))

	}
}

//todo: rename package to httpUtils or smth like
