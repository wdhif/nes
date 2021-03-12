package ui

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func Run() {
	// GLFW Init
	err := glfw.Init()
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
}
