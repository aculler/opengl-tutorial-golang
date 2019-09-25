package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	if err := glfw.Init(); err != nil {
		panic("Failed to initialize GLFW")
	}

	glfw.WindowHint(glfw.Samples, 4)             // 4x antialiasing
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // We want OpenGL 4.1
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // Not sure if this is still required, but here it is.
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // New OpenGL questionmark

	// Open a window and create an OpenGL context
	window, err := glfw.CreateWindow(1024, 768, "Tutorial 1", nil, nil)
	if err != nil {
		panic("Failed to open GLFW window.")
	}
	window.MakeContextCurrent()
	defer glfw.Terminate()

	// Create OpenGL Program
	// Differs a bit from the source material since this doesn't seem needed in C++
	if err := gl.Init(); err != nil {
		panic("Failed to initialize OpenGL")
	}
	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	// end different

	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(prog)

		glfw.PollEvents()
		window.SwapBuffers()

		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			window.SetShouldClose(true)
		}
	}
}
