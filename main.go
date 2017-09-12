package main

import (
	"fmt"
	"os"

	"github.com/wdhif/nes/nes"
)

func main() {
	if len(os.Args) > 1 {
		path := os.Args[1]
		fmt.Println(path)
		nes.Loader(path)
	} else {
		fmt.Println("Usage go run main.go nestest.nes")
	}
}
