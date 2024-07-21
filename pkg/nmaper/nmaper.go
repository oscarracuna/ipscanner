package nmaper

import (
	"context"
	"time"
  "fmt"
  "strings"

	"github.com/Ullaakut/nmap/v3"
)

var (
  ip string
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

  var porto []string
  for _, host := range result.Hosts {
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      continue
    }

	  for _, port := range host.Ports {
      porto = append(porto, fmt.Sprintf("%d -> %s",port.ID, port.State))
    }
  }
  return strings.Join(porto, "\n")
}

