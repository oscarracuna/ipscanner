package nmaper

import (
	"context"
	"time"
  "fmt"
  "log"

	"github.com/Ullaakut/nmap/v3"
)

var (
  ip string
  porto string
)

func Nmap(ip string) string{
  ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
  defer cancel()

  scanner, err := nmap.NewScanner(
    ctx,
    nmap.WithTargets(ip),
    nmap.WithPorts("80,443,8080"),
    )
  if err != nil {
    log.Fatalf("Unable to create nmap scanner")
  }

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", *warnings)
	}
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

  for _, host := range result.Hosts {
    if len(host.Ports) == 0 || len(host.Addresses) == 0 {
      continue
    }
  
	for _, port := range host.Ports {
      //port := (port.ID, port.Protocol, port.State, port.Service.Name)
      porto := fmt.Sprint(port.State)
      fmt.Sprint(porto)
  }
  }	
  return porto
}

