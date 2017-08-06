package main

import (
	"fmt"
	"os"

	"github.com/wdhif/nes/nes"
)

func main() {
	path := os.Args[1]
	fmt.Println(path)
	nes.Loader(path)
}
