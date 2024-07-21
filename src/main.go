package main

// TODO: Figure out go routines or threads
// TODO: Pinger function
// UPDATE: Concurreny works really well but we get false negatives
// TODO: Improve accuracy by implementing redundancy
// Redundancy is not needed. Going for higher threshold with timeouts

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/oscarracuna/ipscanner/pkg/ascii"
  "github.com/oscarracuna/ipscanner/pkg/nmaper"
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

        ports := nmaper.Nmap(newIP)

				pingu, _ := probing.NewPinger(newIP)
				pingu.SetPrivileged(true)
				pingu.Count = 1
				pingu.Timeout = 1000 * time.Millisecond
				pingu.Run()

				stats := pingu.Statistics()
				rcv := stats.PacketsRecv
				if rcv >= 1 {
					if i == 126 {
            results <- fmt.Sprintf("%sHost alive: %s%s <- Fortigate - Ports: %s", Green, Reset, newIP, ports)
					} else {
            results <- fmt.Sprintf("%sHost alive: %s%s - Ports: %s", Green, Reset, newIP, ports)
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

		fmt.Println("\nScan completed.")
		fmt.Print("\n\n\n")
		goto prompt
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
