package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func GetJsonEncodedResponse(data interface{}) (int, string) {
	var response []byte
	var err error

	response, err = json.Marshal(data)

	if err != nil {
		fmt.Println("ERR", err)
		return 500, err.Error()
	}

	return 200, string(response)
}

func GetDataFromBody(body io.ReadCloser, result interface{}) error {
	defer body.Close()
	var data, _ = ioutil.ReadAll(body)
	var err = json.Unmarshal(data, &result)

	return err
}
