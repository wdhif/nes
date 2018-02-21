package main

import (
	"fmt"
	"os"

	"github.com/wdhif/nes/nes"
)

func main() {
	if len(os.Args) > 1 {
		path := os.Args[1]
		fmt.Println("Path to the rom:", path)
		nes.Loader(path)
	} else {
		fmt.Println("Usage go run main.go roms/nestest.nes")
	}
}
