package main

// TODO: Figure out go routines or threads
// TODO: Pinger function
// UPDATE: Concurreny works really well but we get false negatives
// TODO: Improve accuracy by implementing redundancy
// Redundancy is not needed. Going for higher threshold for timeouts
// nmap works great but Go's net library can do it as well


import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/oscarracuna/ipscanner/pkg/ascii"
	//"github.com/oscarracuna/ipscanner/pkg/nmaper"
	probing "github.com/prometheus-community/pro-bing"
)

var (
	ip          string
	Green       = "\033[32m"
	Reset       = "\033[0m"
	printerOrNot bool
	itsPrinter  string
)

func isIP(ip string) bool {
	var ipRegex = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}$`)
	return ipRegex.MatchString(ip)
}


func hostName(ip string) []string {
	name, _ := net.LookupAddr(ip)

	return name
}


func scanPort(ip string, port string) bool{
  address := fmt.Sprintf("%s:%s", ip, port)
  timeout := 3 * time.Second
  a, err := net.DialTimeout("tcp", address, timeout)
  if err != nil {
    return false
  }
  a.Close()
  return true

}

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

				//ports := nmaper.Nmap(newIP)
				//printer := isPrinter(newIP)
				hostNombre := hostName(newIP)

				pingu, _ := probing.NewPinger(newIP)
				pingu.SetPrivileged(true)
				pingu.Count = 1
				pingu.Timeout = 1000 * time.Millisecond
				pingu.Run()
        o := scanPort(newIP, "80")
        if o == true {
          results <- fmt.Sprintf("=== Port 80 open ===")
        }

				stats := pingu.Statistics()
				rcv := stats.PacketsRecv
				if rcv >= 1 {
					if i == 126 {
						results <- fmt.Sprintf("%sHost alive: %s%s <- Fortigate\nHost name:%s\n", Green, Reset, newIP, hostNombre)
					} else {
						results <- fmt.Sprintf("%sHost alive: %s%s\nHost name:%s\n", Green, Reset, newIP, hostNombre)
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

		fmt.Println("\n-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
		fmt.Printf("\t%sScan completed.%s\n", Green, Reset)
		fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
		fmt.Print("\n\n\n")
		goto prompt
	} else {
		fmt.Println("Invalid IP. Please provide a valid IP address.")
		goto prompt
	}
}

