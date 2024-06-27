package main

import (
  "fmt"
  "net"
  //"strings"
 // "bufio"
//  probing "github.com/prometheus-community/pro-bing"
)


func main() {
  var oct1, oct2, oct3, oct4 byte

  fmt.Print("Enter IP: ")
  fmt.Scanln(&oct1, &oct2, &oct3, &oct4)

  ip := net.IPv4(oct1, oct2, oct3, oct4)

  fmt.Println(ip)
  
  

}
