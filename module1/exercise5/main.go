package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (a IPAddr) String() string {
	return strconv.Itoa(int(a[0])) + "." +
		strconv.Itoa(int(a[1])) + "." +
		strconv.Itoa(int(a[2])) + "." +
		strconv.Itoa(int(a[3]))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
