package main

// TODO: Figure out go routines or threads
// TODO: Pinger function
// UPDATE: Concurreny works really well but we get false negatives
// TODO: Improve accuracy by implementing redundancy
// Redundancy is not needed. Going for higher threshold for timeouts
// nmap works great but Go's net library can do it as well

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/oscarracuna/ipscanner/pkg/adios"
	"github.com/oscarracuna/ipscanner/pkg/ascii"
	"github.com/oscarracuna/ipscanner/pkg/portScan"
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

	confirmIP := portScan.IsIP(ip)

	if confirmIP {
		splitIP := strings.Split(ip, ".")
		arrayOfIP := splitIP[0:3]
		joinedArrayofIP := strings.Join(arrayOfIP, ".")

		var wg sync.WaitGroup
		results := make(chan string, 256)
		for i := 0; i <= 255; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				temp := strconv.Itoa(i)
				newIP := joinedArrayofIP + "." + temp

				hostNombre := portScan.HostName(newIP)

				pingu, _ := probing.NewPinger(newIP)
				pingu.SetPrivileged(true)
				pingu.Count = 1
				pingu.Timeout = 1000 * time.Millisecond
				pingu.Run()

				stats := pingu.Statistics()
				rcv := stats.PacketsRecv
				openPorts := portScan.PortScan(newIP)
				if rcv >= 1 {
					if i == 100 {
						results <- fmt.Sprintf("%sHost alive: %s%s <- Printer\nHost name:%s\nOpen ports: %v", Green, Reset, newIP, hostNombre, openPorts)
					}
					if i == 103 {
						results <- fmt.Sprintf("%sHost alive: %s%s <- Scanner\nHost name:%s\nOpen ports: %v", Green, Reset, newIP, hostNombre, openPorts)
					}
					if i == 126 {
						results <- fmt.Sprintf("%sHost alive: %s%s <- Fortigate\nHost name:%s\nOpen ports: %v", Green, Reset, newIP, hostNombre, openPorts)
					} else {
						results <- fmt.Sprintf("%sHost alive: %s%s\nHost name:%s\nOpen ports: %v", Green, Reset, newIP, hostNombre, openPorts)
					}
				}
			}(i)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		for result := range results {
			fmt.Println(result)
		}

		adios.Adios()
		goto prompt
	} else {
		fmt.Println("Invalid IP. Please provide a valid IP address.")
		goto prompt
	}
}
