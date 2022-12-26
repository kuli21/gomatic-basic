package main

import (
	"fmt"

	"github.com/kuli21/gomatic-basic/device"
)

func main() {
	//Get these by your env-File e.g.
	dc := device.HomematicConfig{
		Url:  "http://192.168.188.32",
		Port: 2001,
	}
	ccuUrl := device.GetCcuUrl(dc)

	y := device.RailMountingSwitch{Address: "OEQ0071814"}
	y.GetData(ccuUrl)
	fmt.Printf("p: %#v\n", y)

	//var x device.PowerPlugSwitch
	x := device.PowerPlugSwitch{Address: "REQ1797859:1"}
	fmt.Printf("p: %#v\n", x)
	x.GetData(ccuUrl)
	fmt.Printf("p: %#v\n", x)

	//x.ToggleOff(ccuUrl)
	//x.GetData(ccuUrl)
	//fmt.Printf("p: %#v\n", x)
}
