package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"
import "unsafe"

type SourceT C.NDIlib_source_t

func (t SourceT) Name() string {
	return C.GoString(t.p_ndi_name)
}

func (t SourceT) URLAddress() string {
	return C.GoString(*(**C.char)(unsafe.Pointer(&t.anon0)))
}

type FrameTypeE C.NDIlib_frame_type_e

const (
	FrameTypeNone     = FrameTypeE(C.NDIlib_frame_type_none)
	FrameTypeVideo    = FrameTypeE(C.NDIlib_frame_type_video)
	FrameTypeAudio    = FrameTypeE(C.NDIlib_frame_type_audio)
	FrameTypeMetadata = FrameTypeE(C.NDIlib_frame_type_metadata)
	FrameTypeError    = FrameTypeE(C.NDIlib_frame_type_error)
	FrameTypeMax      = FrameTypeE(C.NDIlib_frame_type_max)
)

type VideoFrameV2T C.NDIlib_video_frame_v2_t

func (t VideoFrameV2T) Xres() int {
	return int(t.xres)
}

func (t VideoFrameV2T) Yres() int {
	return int(t.yres)
}

type AudioFrameV2T C.NDIlib_audio_frame_v2_t

func (t AudioFrameV2T) NoSamples() int {
	return int(t.no_samples)
}

type MetadataFrameT C.NDIlib_metadata_frame_t
