package handler

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/kuli21/gomatic-basic/model"
)

func GetParamset(address string, ccuUrl string) (model.XResponse, error) {
	reader := model.GetParamsetReader(address)
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

func SetValue(address string, param string, value string, ccuUrl string) error {
	reader := model.SetValueReader(address, param, value)
	request, err := http.NewRequest("GET", ccuUrl, reader)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		bodyString = strings.Replace(bodyString, `encoding="iso-8859-1"`, "", 1)
		if IsFaultResponse(bodyString) {
			return errors.New("fault XML Response: " + bodyString)
		} else {
			return nil
		}
	} else {
		return errors.New("Response Error StatusCode: " + strconv.Itoa(resp.StatusCode))
	}
}
