package main

import (
	"fmt"
	"time"

	idevice "github.com/kai5263499/go-idevice"
)

func main() {
	idevice.Start()
	devices, num_devices := idevice.GetDeviceList()
	fmt.Printf("device count %d [%+v]\n", num_devices, devices)

	select {
	case <-time.After(time.Second * 30):
		fmt.Println("timeout 10")
	}
}
