package main

// TODO: Figure out go routines or threads

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/oscarracuna/ipscanner/pkg/ascii"
	probing "github.com/prometheus-community/pro-bing"
)

var (
	ip    string
	Green = "\033[32m"
	Reset = "\033[0m"
)

func main() {

	fmt.Println(ascii.Ascii_saludo())

prompt:
	fmt.Print("Enter LAN IP or VNC: ")
	fmt.Scanln(&ip)
  
	confirmIP := isIP(ip)
//nmapTest(ip)
  
	if confirmIP == true {
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
			pingu.Timeout = 200 * time.Millisecond
			pingu.Run()

			stats := pingu.Statistics()
			rcv := stats.PacketsRecv
			if rcv >= 1 {
				fmt.Println(Green+"Host alive:", Reset+newIP)
			}
			if rcv >= 1 && i == 126 {
				fmt.Println(Green+"Host alive:", Reset+newIP, "<- Fortigate")
			}
			if i == 255 {
				fmt.Println("\nScan completed.")
				fmt.Print("\n\n\n")
				goto prompt
			}
		}
	} else {
		fmt.Println("Invalid IP. Please provide a valid IP address.")
		goto prompt
	}

}

func isIP(ip string) bool {
	var ipRegex = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}$`)
	return ipRegex.MatchString(ip)
}

//Testing port discovery with nmap library for Go
/*
func nmapTest(test string) {
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
  defer cancel()


  scanner, err := nmap.NewScanner(
    ctx,
    nmap.WithTargets(&ip),
    nmap.WithPorts(ports)
  )
  return 
}
*/

