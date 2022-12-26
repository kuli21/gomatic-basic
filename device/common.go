package device

import (
	"fmt"
	"strconv"

	"github.com/kuli21/gomatic-basic/handler"
	"github.com/kuli21/gomatic-basic/model"
)

type HomematicConfig struct {
	Url  string
	Port int
}

func GetCcuUrl(hmConfig HomematicConfig) string {
	hmPort := strconv.Itoa(hmConfig.Port)
	ccuUrl := hmConfig.Url + ":" + hmPort
	return ccuUrl
}

func PrintAllDevices(ccuUrl string) {
	d := GetAllDevices(ccuUrl)
	fmt.Println("----------------------------------")
	fmt.Printf("d: %#v\n", d)
	fmt.Println("----------------------------------")
}

func PrintParamset(address string, ccuUrl string) {
	r, err := handler.GetParamset(address, ccuUrl)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("----------------------------------")
		fmt.Printf("d: %#v\n", r)
		fmt.Println("----------------------------------")
	}
}

func GetAllDevices(ccuUrl string) model.XResponse {
	r, err := handler.ListAllDevices(ccuUrl)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func Int2Bool(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}
