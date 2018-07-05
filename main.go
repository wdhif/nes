package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"

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

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	fmt.Fprint(os.Stdout, fmt.Sprintf(BANNER, "0.1.0"))

	if len(os.Args) > 1 {
		path := os.Args[1]
		fmt.Println("Path to the rom:", path)

		rom, err := nes.Loader(path)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Printing ROM data")
		fmt.Println(rom)

		// GLFW Init
		err = glfw.Init()
		if err != nil {
			panic(err)
		}
		defer glfw.Terminate()

		// Window Creation
		window, err := glfw.CreateWindow(640, 480, "NES", nil, nil)
		if err != nil {
			panic(err)
		}

		window.MakeContextCurrent()

		// GLFW Loop
		for !window.ShouldClose() {
			// Do OpenGL stuff.
			window.SwapBuffers()
			glfw.PollEvents()
		}
	} else {
		fmt.Println("Usage go run main.go roms/nestest.nes")
	}
}
