package weavedns

import (
	"fmt"
	"net"
	"time"
)

func EnsureInterface(ifaceName string, wait int) (iface *net.Interface, err error) {
	if iface, err = findInterface(ifaceName); err == nil || wait == 0 {
		return
	}
	Info.Println("Waiting for interface", ifaceName, "to come up")
	for ; err != nil && wait > 0; wait -= 1 {
		time.Sleep(1 * time.Second)
		iface, err = findInterface(ifaceName)
	}
	if err == nil {
		Info.Println("Interface", ifaceName, "is up")
	}
	return
}

func findInterface(ifaceName string) (iface *net.Interface, err error) {
	if iface, err = net.InterfaceByName(ifaceName); err != nil {
		return iface, fmt.Errorf("Unable to find interface %s", ifaceName)
	}
	if 0 == (net.FlagUp & iface.Flags) {
		return iface, fmt.Errorf("Interface %s is not up", ifaceName)
	}
	return
}
