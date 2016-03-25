package glew

/*
#include "glew.h"

#cgo LDFLAGS: -L${SRCDIR}/../ -lglew32
*/
import "C"

import (
	"github.com/AlkBur/Tutorial-Golang-GLWF/Tutorial_1/gl"
)

const (
	OK = C.GLEW_OK
)

func InitCore() bool {
	C.glewExperimental = gl.TRUE // Needed in core profile
	if C.glewInit() != OK {
		return false
	}
	return true
}
