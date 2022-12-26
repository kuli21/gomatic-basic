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
	//Only necessary for further development of library
	//------------------------------------------------------
	fmt.Println("------- How to implement new device -------")
	fmt.Println("--- 1) Print all devices and search for your device (number) -------")
	device.PrintAllDevices()
	fmt.Println("--- 2) Get the device number (id) and put it in the function e.g. devices.PrintParamset(OE9999999:1) -------")
	device.PrintParamset("QE1111112:1")
}
```

## Usage
```
See examples in example folder!
```

## License
Open Source

## Project status
Just began to implement this - so status is ongoing
