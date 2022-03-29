package main

import (
	"fmt"
	"net"
)

const URL = "fontys.nl"

func main() {
	iprecords, _ := net.LookupIP(URL)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}
}
