# gomatic-basic

## Name
Go-Homematic-Basic library

## Description
Basic library for connecting to a local CCU/CCU2 and getting values from several devices and setting values on those devices

## Example how to add new devices
```
package main

import (
	"fmt"
	"github.com/kuli21/gomatic-basic/device"
)

func main() {
	fmt.Println("------- How to implement new device -------")
	fmt.Println("--- 1) Print all devices and search for your device (number) -------")
	device.PrintAllDevices()
	fmt.Println("--- 2) Get the device number (id) and put it in the function e.g. devices.PrintParamset(OE9999999:1) -------")
	device.PrintParamset("QE1111112:1")

	//Special case for rail-mounting-switch - that contains 2 addresses - you can request combined!
	p := device.GetRailMountingSwitchDataAll("OE3333333")
	fmt.Printf("p: %#v\n", p)
}
```

## Usage
```
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

	fmt.Println("------- Getting values from Thermostat Valve Device MEQ000000:4 -------")
	t := device.GetThermostatValveData("MEQ0000000:4", ccuUrl)
	fmt.Printf("t: %#v\n", t)

	fmt.Println("------- Getting values from Power Plug Switch Device REQ000000:1 -------")
	p := device.GetPowerPlugSwitchData("REQ0000000:1", ccuUrl)
	fmt.Printf("p: %#v\n", p)

	fmt.Println("---------- TURN ON Power Plug Switch -----------")
	device.ToggleOnPowerPlugSwitch("REQ0000000:1", ccuUrl)

	time.Sleep(4 * time.Second)

	fmt.Println("---------- TURN OFF Power Plug Switch -----------")
	device.ToggleOffPowerPlugSwitch("REQ0000000:1", ccuUrl)

	fmt.Println("---------- Getting Values from Rail Mounting Switch (:1 and :2) -----------")
	p := device.GetRailMountingSwitchDataAll("O33333333", ccuUrl)
	fmt.Printf("p: %#v\n", p)
}
```
OUTPUT
```
t: devices.ThermostatValve {
    Address:"MEQ0806594:4", 
    ActualTemperature:23.4, 
    BatteryState:2.7, 
    BoostState:0, 
    ControlMode:0, 
    FaultReporting:0, 
    PartyStartDay:1, 
    PartyStartMonth:1, 
    PartyStartTime:0, 
    PartyStartYear:0, 
    PartyStopDay:1, 
    PartyStopMonth:1, 
    PartyStopTime:0, 
    PartyStopYear:0, 
    PartyTemperature:5, 
    SetTemperature:20.5, 
    ValveState:0
    }
```

## License
Open Source

## Project status
Just began to implement this - so status is ongoing
