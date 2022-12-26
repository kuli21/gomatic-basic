package handler

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/kuli21/gomatic-basic/model"
)

func ListAllDevices(ccuUrl string) (model.XResponse, error) {
	reader := model.ListDevicesReader()
	request, err := http.NewRequest("GET", ccuUrl, reader)
	if err != nil {
		return model.XResponse{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return model.XResponse{}, err
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return model.XResponse{}, err
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		bodyString = strings.Replace(bodyString, `encoding="iso-8859-1"`, "", 1)
		var response model.XResponse
		err = xml.Unmarshal([]byte(bodyString), &response)
		if err != nil {
			return model.XResponse{}, err
		}
		return response, nil
	} else {
		return model.XResponse{}, errors.New("Response Error StatusCode: " + strconv.Itoa(resp.StatusCode))
	}
}

func IsFaultResponse(bodyString string) bool {
	i1 := strings.Contains(bodyString, "<name>faultCode</name>")
	i2 := strings.Contains(bodyString, "<fault>")
	i3 := strings.Contains(bodyString, "<methodResponse>")
	if i1 && i2 && i3 {
		return true
	} else {
		return false
	}
}
