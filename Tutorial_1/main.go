package main

import (
	"fmt"
	gl "github.com/AlkBur/Tutorial-Golang-GLWF/Tutorial_1/gl"
	glew "github.com/AlkBur/Tutorial-Golang-GLWF/Tutorial_1/glew"
	glfw "github.com/AlkBur/Tutorial-Golang-GLWF/Tutorial_1/glfw"
)

func main() {
	if !glfw.InitCore() {
		fmt.Println("Failed to initialize GLFW")
		return
	}
	glfw.WindowHint(glfw.GLFW_SAMPLES, 4)               // 4x antialiasing
	glfw.WindowHint(glfw.GLFW_CONTEXT_VERSION_MAJOR, 3) // We want OpenGL 3.3
	glfw.WindowHint(glfw.GLFW_CONTEXT_VERSION_MINOR, 3)
	glfw.WindowHint(glfw.GLFW_OPENGL_FORWARD_COMPAT, gl.TRUE)                // To make MacOS happy; should not be needed
	glfw.WindowHint(glfw.GLFW_OPENGL_PROFILE, glfw.GLFW_OPENGL_CORE_PROFILE) //We don't want the old OpenGL

	glfw.CreateWindow(800, 600, "Tutorial 01", nil, nil)
	if !glfw.IsCreateWindow() {
		fmt.Println("Failed to open GLFW window. If you have an Intel GPU, they are not 3.3 compatible. Try the 2.1 version of the tutorials.")
		glfw.Terminate()
		return
	}

	glfw.MakeContextCurrent() // Initialize GLEW
	if !glew.InitCore() {
		glfw.Terminate()
		fmt.Println("Failed to initialize GLEW")
		return
	}

	// Ensure we can capture the escape key being pressed below
	glfw.SetInputMode(glfw.GLFW_STICKY_KEYS, gl.TRUE)

	for !glfw.WindowShouldClose() && glfw.GetKey(glfw.GLFW_KEY_ESCAPE) != glfw.GLFW_PRESS {
		glfw.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}
