package isPrinter

// TODO figure out if array or boolean approach
// Figure out net.JoinHostPort(host, port)
// Figure out conn
// wtf is that

import (
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
