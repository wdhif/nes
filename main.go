package main

import (
	"fmt"
	"os"

	"github.com/wdhif/nes/nes"
	"log"
)

func main() {
	if len(os.Args) > 1 {
		path := os.Args[1]
		fmt.Println("Path to the rom:", path)
		_, err := nes.Loader(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Usage go run main.go roms/nestest.nes")
	}
}
