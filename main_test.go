package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/wdhif/nes/nes"
)

func TestMain(m *testing.M) {
	fmt.Fprint(os.Stdout, fmt.Sprintf(BANNER, "0.1.0"))
	if len(os.Args) > 1 {
		path := "roms/nestest.nes"
		fmt.Println("Path to the rom:", path)
		_, err := nes.Loader(path)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Usage go run main.go roms/nestest.nes")
	}
}
