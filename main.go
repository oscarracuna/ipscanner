package main

import (
  "fmt"
  "strings"
  probing "github.com/prometheus-community/pro-bing"
)


func main() {
  var ip string
  fmt.Print("Enter IP: ")
  fmt.Scanln(&ip)

  // This returns an array of the values
  // [192 168 0 1]
  // Which we can later call by its index
  ip_s := strings.Split(ip,".")


  // Original ip
  fmt.Println(ip)
  // Array of the values
  fmt.Println(ip_s)  

  // When dealing with ranges of arrays[0:2], you get index 0 and 1, but not 2.
  // For example, in [0:3], you get 0,1 AND 2.
  oct1 := ip_s[0:2]
  fmt.Println(oct1)

  pinger, err := probing.NewPinger("google.com")
  if err != nil {
    panic(err)
  }

  // How many packets
  pinger.Count = 3
  //You need this to run the pinger
  pinger.Run()

  // This just prints all of the stats.
  // You can do pinger.[the specific stat you want]
  stats := pinger.Statistics()
  fmt.Println(stats)

}
