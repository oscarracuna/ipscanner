package nmaper

import (
	"context"
	"time"
  "fmt"

	"github.com/Ullaakut/nmap/v3"
)

var (
  ip string
  porto string
)

func Nmap(ip string) string {
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
  defer cancel()

  scanner, _ := nmap.NewScanner(
    ctx,
    nmap.WithTargets(ip),
    nmap.WithPorts("80,443,8080"),
    )

	result, _, _ := scanner.Run()

  for _, host := range result.Hosts {
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      continue
    }
  
	for _, port := range host.Ports {
      fmt.Print(port.State)
    }
  }
  return porto
}

