package main

import (
    "fmt"
    "time"
)

/*
#cgo darwin CFLAGS: -Wmacro-redefined -Wincompatible-pointer-types-discards-qualifiers -I/usr/local/include/libimobiledevice
#cgo darwin LDFLAGS: -limobiledevice
#include <libimobiledevice/libimobiledevice.h>
#include <libimobiledevice/syslog_relay.h>
#include <libimobiledevice/lockdown.h>

#define NULL 0

extern void DeviceEventCbGo(idevice_event_t* event);

static inline void device_event_cb(const idevice_event_t* event, void* userdata) {
    DeviceEventCbGo(event);
}

static inline void subscribe_to_device_events() {
    idevice_event_subscribe(device_event_cb, NULL);
}

*/
import "C"

func getDeviceList() (string, int) {
    num := C.int(0)
    var buf **C.char

    C.idevice_get_device_list(&buf, &num)
    defer C.idevice_device_list_free(buf)
    return C.GoString(*buf), int(num)
}

//export DeviceEventCbGo
func DeviceEventCbGo(event *C.idevice_event_t) {
    fmt.Printf("got device event: [%v]\n", event)
    switch event.event {
        case C.IDEVICE_DEVICE_ADD:
        fmt.Printf("device %s added!\n", C.GoString(event.udid))
    }
}

func main() {
    devices, num_devices := getDeviceList()
    fmt.Printf("device count %d [%+v]\n", num_devices, devices)

    C.subscribe_to_device_events()

    select {
        case <-time.After(time.Second * 10):
            fmt.Println("timeout 10")
    }
}
