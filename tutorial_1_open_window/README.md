# Tutorial 1: Open a Window
In this tutorial we're going to simply open a window that can be closed when you hit the escape key.

I'm going to step through the code below, but the full source can be found [here](tutorial_1_open_window/main.go)

## Imports
In order to work with OpenGL we're going to use the [go-gl](https://github.com/go-gl/) library. We'll need to import the core GL module as wel as the Graphics Library Framework module.

```go
import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)
```

## GLFW Initialization
The Graphics Library Framework (GLEW) is a utility library for OpenGL. In this tutorial we'll be using it to create our window that will then run our OpenGL program.

```go
if err := glfw.Init(); err != nil {
    panic("Failed to initialize GLFW")
}

glfw.WindowHint(glfw.Samples, 4) // 4x antialiasing
glfw.WindowHint(glfw.ContextVersionMajor, 4) // We want OpenGL 4.1
glfw.WindowHint(glfw.ContextVersionMinor, 1)
glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) // Not sure if this is still required, but here it is.
glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // New OpenGL questionmark

// Open a window and create an OpenGL context
window, err := glfw.CreateWindow(1024, 768, "Tutorial 1", nil, nil)
if err != nil {
    panic("Failed to open GLFW window.")
}
window.MakeContextCurrent()
defer glfw.Terminate()
```

First we initialize the glfw library and panic if it fails.

Next we set a number of configuration values for the window that we're creating.

And finally we actually create the window and panic on error. After that we have to tell the window that it's the currently active window and defer termination of the initialized glfw library.

## Create an OpenGL Program
This section deviates the most from the original C++ source. I'm not sure if/why C++ doesn't require OpenGL to be initialized in order for GLFW to display a window, but for go-gl it seems to be required.

```go
if err := gl.Init(); err != nil {
    panic("Failed to initialize OpenGL")
}
prog := gl.CreateProgram()
gl.LinkProgram(prog)
```

It's fairly straightforward in Go. We just initialize the gl library, panic on an error, and then tell the gl library to create a new program. Later this program will store things like shaders, but for now it just needs to exist. Finally, we link the program, which is compiling the shaders that are linked to the program.

## Actually Draw the Window
Finally we get to the code that is utilizing all of the above setup code to actually render our window and listen for the events that allow us to eventually close the window with the escape key.

```go
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
```

First we're telling our window to turn on `StickyKeysMode`. This isn't the infamous Windows stickykeys that has annoyed so many of us over the years, but is instead telling the window to keep track of pressed keys until the next `GetKey` function call. If we did not enable this mode but still used the `GetKey` function (instead of a button press callback), then we would run the risk that a key would be pressed and released all before we polled for the key press event, which would cause us to miss the key.

Next we enter the main loop. There are a number of reasons that a window might need to close, so we want to use the `window.ShouldClose()` as our conditional. We then clear the screen, tell OpenGL to use our (currently empty) program, poll for events (used to detect the keypress in this instance), swap our window buffer (doesn't do anything because we aren't rendering anything yet), and finally check to see if the escape key was pressed. If it was pressed, we let GLFW know that it should close the window.

And that's it! If you build and execute your code you should see a blank window appear, which sets the stage nicely for [Tutorial 2](tutorial_2_first_triangle/README.md).