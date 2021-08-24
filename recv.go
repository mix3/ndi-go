package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"

type RecvBandwidthE C.NDIlib_recv_bandwidth_e

const (
	RecvBandwidthMetadataOnly = RecvBandwidthE(C.NDIlib_recv_bandwidth_metadata_only)
	RecvBandwidthAudioOnly    = RecvBandwidthE(C.NDIlib_recv_bandwidth_audio_only)
	RecvBandwidthLowest       = RecvBandwidthE(C.NDIlib_recv_bandwidth_lowest)
	RecvBandwidthHighest      = RecvBandwidthE(C.NDIlib_recv_bandwidth_highest)
	RecvBandwidthMax          = RecvBandwidthE(C.NDIlib_recv_bandwidth_max)
)

type RecvColorFormatE C.NDIlib_recv_color_format_e

const (
	RecvColorFormat_BGRX_BGRA = RecvColorFormatE(C.NDIlib_recv_color_format_BGRX_BGRA)
	RecvColorFormat_UYVY_BGRA = RecvColorFormatE(C.NDIlib_recv_color_format_UYVY_BGRA)
	RecvColorFormat_RGBX_RGBA = RecvColorFormatE(C.NDIlib_recv_color_format_RGBX_RGBA)
	RecvColorFormat_UYVY_RGBA = RecvColorFormatE(C.NDIlib_recv_color_format_UYVY_RGBA)
	RecvColorFormat_fastest   = RecvColorFormatE(C.NDIlib_recv_color_format_fastest)
	RecvColorFormat_best      = RecvColorFormatE(C.NDIlib_recv_color_format_best)
	RecvColorFormat_max       = RecvColorFormatE(C.NDIlib_recv_color_format_max)
)

type RecvInstanceT C.NDIlib_recv_instance_t

type RecvCreateV3T C.NDIlib_recv_create_v3_t

func NewRecvCreateV3T(opts ...NewRecvCreateV3TOption) RecvCreateV3T {
	v := RecvCreateV3T{
		color_format:       C.NDIlib_recv_color_format_e(RecvColorFormat_UYVY_BGRA),
		bandwidth:          C.NDIlib_recv_bandwidth_e(RecvBandwidthHighest),
		allow_video_fields: true,
	}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

type NewRecvCreateV3TOption func(*RecvCreateV3T)

func NewRecvCreateV3TOptionSourceToConnectTo(v SourceT) NewRecvCreateV3TOption {
	return func(t *RecvCreateV3T) {
		t.source_to_connect_to = C.NDIlib_source_t(v)
	}
}

func NewRecvCreateV3TOptionColorFormat(v RecvColorFormatE) NewRecvCreateV3TOption {
	return func(t *RecvCreateV3T) {
		t.color_format = C.NDIlib_recv_color_format_e(v)
	}
}

func NewRecvCreateV3TOptionBandwidth(v RecvBandwidthE) NewRecvCreateV3TOption {
	return func(t *RecvCreateV3T) {
		t.bandwidth = C.NDIlib_recv_bandwidth_e(v)
	}
}

func NewRecvCreateV3TOptionAllowVideoFields(v bool) NewRecvCreateV3TOption {
	return func(t *RecvCreateV3T) {
		t.allow_video_fields = C.bool(v)
	}
}

func NewRecvCreateV3TOptionName(v *string) NewRecvCreateV3TOption {
	return func(t *RecvCreateV3T) {
		if v != nil {
			t.p_ndi_recv_name = C.CString(*v)
		} else {
			t.p_ndi_recv_name = nil
		}
	}
}

func RecvCreateV3(v RecvCreateV3T) RecvInstanceT {
	return RecvInstanceT(C.NDIlib_recv_create_v3((*C.NDIlib_recv_create_v3_t)(&v)))
}

func RecvDestory(t RecvInstanceT) {
	C.NDIlib_recv_destroy(C.NDIlib_recv_instance_t(t))
}

func RecvConnect(instance RecvInstanceT, sourceT *SourceT) {
	C.NDIlib_recv_connect(C.NDIlib_recv_instance_t(instance), (*C.NDIlib_source_t)(sourceT))
}

func RecvCaptureV2(instance RecvInstanceT, videoFrame *VideoFrameV2T, audioFrame *AudioFrameV2T, metadataFrame *MetadataFrameT, timeoutInMs uint32) FrameTypeE {
	var (
		cVideoFrame    C.NDIlib_video_frame_v2_t
		cAudioFrame    C.NDIlib_audio_frame_v2_t
		cMetadataFrame C.NDIlib_metadata_frame_t
	)
	ret := C.NDIlib_recv_capture_v2(C.NDIlib_recv_instance_t(instance), &cVideoFrame, &cAudioFrame, &cMetadataFrame, C.uint32_t(timeoutInMs))
	if videoFrame != nil {
		*videoFrame = VideoFrameV2T(cVideoFrame)
	}
	if audioFrame != nil {
		*audioFrame = AudioFrameV2T(cAudioFrame)
	}
	if metadataFrame != nil {
		*metadataFrame = MetadataFrameT(cMetadataFrame)
	}
	return FrameTypeE(ret)
}

func RecvFreeVideoV2(instance RecvInstanceT, videoFrame *VideoFrameV2T) {
	C.NDIlib_recv_free_video_v2(C.NDIlib_recv_instance_t(instance), (*C.NDIlib_video_frame_v2_t)(videoFrame))
}

func RecvFreeAudioV2(instance RecvInstanceT, audioFrame *AudioFrameV2T) {
	C.NDIlib_recv_free_audio_v2(C.NDIlib_recv_instance_t(instance), (*C.NDIlib_audio_frame_v2_t)(audioFrame))
}

func RecvFreeMetadataV2(instance RecvInstanceT, metadataFrame *MetadataFrameT) {
	C.NDIlib_recv_free_metadata(C.NDIlib_recv_instance_t(instance), (*C.NDIlib_metadata_frame_t)(metadataFrame))
}
