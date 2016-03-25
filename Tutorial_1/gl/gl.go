package gl

/*
#include <gl\gl.h>

#cgo LDFLAGS: -lopengl32
*/
import "C"

const (
	TRUE  = C.GL_TRUE
	FALSE = C.GL_FALSE
)
