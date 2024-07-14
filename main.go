package main

// TODO: Error handleling
// Infinite loop in case IP is not valid 
// Add csv iteration for caps so you can scan ping the entire office with just the cap number

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

	fmt.Print("Enter IP: ")
	fmt.Scanln(&ip)
	ip_s := strings.Split(ip, ".")
	threeOctets := ip_s[0:3]
	preIP := strings.Join(threeOctets, ".")

	for i := 0; i <= 255; i++ {
		temp := strconv.Itoa(i)
		rip := preIP + "." + temp

		pingu, _ := probing.NewPinger(rip)
		pingu.SetPrivileged(true)
		pingu.Count = 1
		pingu.Timeout = 1000
		pingu.Run()

		stats := pingu.Statistics()
		rcv := stats.PacketsRecv
		if rcv >= 1 {
			fmt.Println("Host alive:", rip)
		}
	}
}
