package main

// TODO: Error handleling
// Infinite loop in case IP is not valid 
// Add csv iteration for caps so you can scan ping the entire office with just the cap number
// Possibly add port discovery and output which endpoints may be printers

import (
	"fmt"
	"strconv"
	"strings"

	probing "github.com/prometheus-community/pro-bing"
)

var (
	ip string
)

func main() {

	fmt.Print("Enter LAN IP or VNC: ")
	fmt.Scanln(&ip)
	splitIP := strings.Split(ip, ".")
	arrayOfIP := splitIP[0:3]
	joinedArrayofIP := strings.Join(arrayOfIP, ".")

	for i := 0; i <= 255; i++ {
    // This converts ints to ascii
		temp := strconv.Itoa(i)
		newIP := joinedArrayofIP + "." + temp

		pingu, _ := probing.NewPinger(newIP)
		pingu.SetPrivileged(true)
		pingu.Count = 1
		pingu.Timeout = 10*time.Second
		pingu.Run()

		stats := pingu.Statistics()
		rcv := stats.PacketsRecv
		if rcv >= 1 {
			fmt.Println("Host alive:", newIP)
		}
	}
}
