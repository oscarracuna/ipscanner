package main

// TODO: Error handleling
// Infinite loop in case IP is not valid
// Add csv iteration for caps so you can scan ping the entire office with just the cap number
// Possibly add port discovery and output which endpoints may be printers
// Go routines may be the way

import (
	"fmt"
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
	isIP(ip)
	//Here, create another function that is given an IP, in this case a string. Check if it is a proper IP
	//Requirments mean, 4 subsets of a NUMBER(check for only integers here), each in a range between 0-255
	//Return a boolean, false or true.
	//On False, exit

	// ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	// Should Regex be used for this, instead?
	// They're not readable at all

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
			fmt.Println(Green+"Host alive:", newIP+Reset)
		}
		if rcv >= 1 && i == 126 {
			fmt.Println(Green+"Host alive:", newIP+Reset, "<- Fortigate")
		}

		//Here, you can also put another function to post out a Dead Port
		//It seems slower because I am only seeing 1 side, the alive side
		//Angry scanner is also using threads, we are only using 1 thread here so its doing all of the work.

	}
	fmt.Println("\nScan completed.")
	fmt.Print("\n\n\n")
	goto prompt
}

// Testing func to validate IP

func isIP(ip string) bool {

	splitIP := strings.Split(ip, ".")
	oct1, _ := strconv.Atoi(splitIP[0])
	oct2, _ := strconv.Atoi(splitIP[1])
	oct3, _ := strconv.Atoi(splitIP[2])
	oct4, _ := strconv.Atoi(splitIP[3])

	// Gotta figure out how to make this func work

	return true
}

/* Testing port discovery with nmap library for Go
func nmapTest() {
  ctx, cancel := context.WithTimeout(context.Backgroun(), 1*time.Minute)
  defer cancel()


  scanner, err := nmap.NewScanner(
    ctx,
    nmap.WithTargets(&ip),
    nmap.WithPorts(ports)
  )
}
*/
