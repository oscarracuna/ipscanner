package main

import (
  "fmt"
  
  probing "github.com/prometheus-community/pro-bing"
)

func main() {
  fmt.Println("awa")

  pinger, err := probing.NewPinger()
    if err != nil {
      fmt.Println("Something went wrong with the pinger. ", err)
    }
}
