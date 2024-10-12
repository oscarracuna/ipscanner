package portScan

import (
	"fmt"
	"net"
	"regexp"
	"time"
)

func PortScan(ip string) []string {
	var port string
	noPorts := []string{"No open ports"}
	ports := []string{"80", "443", "8080"}
	openPorts := []string{}
	timeout := 1500 * time.Millisecond
	for _, port = range ports {
		address := fmt.Sprintf("%s:%s", ip, port)
		a, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			continue
		}
		openPorts = append(openPorts, port)
		a.Close()
	}
	if len(openPorts) > 0 {
		return openPorts
	} else {
		return noPorts
	}
}

func IsIP(ip string) bool {
	var ipRegex = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}$`)
	return ipRegex.MatchString(ip)
}

func HostName(ip string) []string {
	noName := []string{"None found"}
	name, _ := net.LookupAddr(ip)

	if len(name) > 0 {
		return name
	} else {
		return noName
	}
}
