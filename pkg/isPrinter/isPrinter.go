package isPrinter

import (
  "fmt"
  "net"
)

func isPrinter(ip string) string {
  port := ":80"
  ipAndPort := ip + port
  _, err := net.DialTimeout("tcp", ipAndPort, 400)
  if err != nil {

  } else {

  }

  return
}
