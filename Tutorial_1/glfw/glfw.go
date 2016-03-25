package glfw

/*
#include <stdlib.h>

#include "glfw3.h"
#cgo LDFLAGS: -L${SRCDIR}/../ -lglfw3
*/
import "C"

import (
	"github.com/AlkBur/Tutorial-Golang-GLWF/Tutorial_1/gl"
	"sync"
	"unsafe"
)

const (
	GLFW_SAMPLES               = C.GLFW_SAMPLES
	GLFW_CONTEXT_VERSION_MAJOR = C.GLFW_CONTEXT_VERSION_MAJOR
	GLFW_CONTEXT_VERSION_MINOR = C.GLFW_CONTEXT_VERSION_MINOR
	GLFW_OPENGL_FORWARD_COMPAT = C.GLFW_OPENGL_FORWARD_COMPAT
	GLFW_OPENGL_PROFILE        = C.GLFW_OPENGL_PROFILE
	GLFW_OPENGL_CORE_PROFILE   = C.GLFW_OPENGL_CORE_PROFILE
	GLFW_STICKY_KEYS           = C.GLFW_STICKY_KEYS
	GLFW_KEY_ESCAPE            = C.GLFW_KEY_ESCAPE
	GLFW_PRESS                 = C.GLFW_PRESS
)

type windowGLF struct {
	glfwWindow *C.GLFWwindow
}

// Инстанс, который будет содержать единственный экземпляр
var instance *windowGLF

// Объект, который позволяет выполнять некоторое действие только один раз
var once sync.Once

func glfwbool(b C.int) bool {
	if b == gl.TRUE {
		return true
	}
	return false
}

func InitCore() bool {
	if !glfwbool(C.glfwInit()) {

		return false
	}
	return true
}

func WindowHint(hint, value C.int) {
	C.glfwWindowHint(hint, value)
}

func CreateWindow(width, height int, title string, monitor *C.GLFWmonitor, share *C.GLFWwindow) {
	once.Do(func() {
		title_char := C.CString(title)
		defer C.free(unsafe.Pointer(title_char))
		w := C.glfwCreateWindow(C.int(width), C.int(height), title_char, monitor, share)
		instance = &windowGLF{glfwWindow: w}
	})
}

func IsCreateWindow() bool {
	if instance != nil && instance.glfwWindow != nil {
		return true
	}
	return false
}

func Terminate() {
	C.glfwTerminate()
}

func MakeContextCurrent() {
	C.glfwMakeContextCurrent(instance.glfwWindow)
}

func SetInputMode(mode, value C.int) {
	C.glfwSetInputMode(instance.glfwWindow, mode, value)
}

func WindowShouldClose() bool {
	return glfwbool(C.glfwWindowShouldClose(instance.glfwWindow))
}

func GetKey(key C.int) C.int {
	return C.glfwGetKey(instance.glfwWindow, key)
}

func SwapBuffers() {
	C.glfwSwapBuffers(instance.glfwWindow)
}

func PollEvents() {
	C.glfwPollEvents()
}
