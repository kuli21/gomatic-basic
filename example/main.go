package main

import (
	"fmt"
	"os"

	"github.com/kuli21/gomatic-basic/device"
)

func main() {
	//Get these by your env-File e.g.
	dc := device.HomematicConfig{
		Url:  os.Getenv("HOMEMATIC_CCU_URL"),
		Port: 2001,
	}
	ccuUrl := device.GetCcuUrl(dc)

	//Init device
	x := device.PowerPlugSwitch{Address: "RE12345678:1"}
	fmt.Printf("p: %#v\n", x)
	//Get data
	x.GetData(ccuUrl)
	fmt.Printf("p: %#v\n", x)
	//Switch off
	x.ToggleOff(ccuUrl)
	//Switch on
	x.ToggleOn(ccuUrl)

	//Init device
	w := device.DoorWindowSensorOptical{Address: "ME12345678:1"}
	//Get data
	w.GetData(ccuUrl)
	fmt.Printf("p: %#v\n", w)
}
