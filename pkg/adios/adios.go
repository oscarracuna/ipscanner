package adios

import (
	"fmt"
)

var (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Adios() {
	fmt.Println("\n-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	fmt.Printf("\t%sScan completed.%s\n", Green, Reset)
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	fmt.Print("\n\n\n")
}
