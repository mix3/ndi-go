package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"

func Initialize() bool {
	return (bool)(C.NDIlib_initialize())
}

func Destroy() {
	C.NDIlib_destroy()
}

func Version() string {
	return C.GoString(C.NDIlib_version())
}

func IsSupportedCPU() bool {
	return (bool)(C.NDIlib_is_supported_CPU())
}
