package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wdhif/nes/nes"
)

const (
	BANNER = `
   _  __________
  / |/ / __/ __/
 /    / _/_\ \  
/_/|_/___/___/  
                
A Nintendo Entertainment System emulator in Go
 Version: %s

`
)

func main() {
	fmt.Fprint(os.Stdout, fmt.Sprintf(BANNER, "0.1.0"))
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
