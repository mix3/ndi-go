package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"

// NDIlib_recv_instance_t NDIlib_recv_create_v3(const NDIlib_recv_create_v3_t* p_create_settings NDILIB_CPP_DEFAULT_VALUE(NULL));

type RecvInstanceT C.NDIlib_recv_instance_t

type RecvCreateV3T C.NDIlib_recv_create_v3_t

func NewRecvCreateV3T(opts ...NewRecvCreateV3TOption) RecvCreateV3T {
	v := RecvCreateV3T{
		//source_to_connect_to: nil,
	}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

//typedef struct NDIlib_recv_create_v3_t
//{	// The source that you wish to connect to.
//	NDIlib_source_t source_to_connect_to;
//
//	// Your preference of color space. See above.
//	NDIlib_recv_color_format_e color_format;
//
//	// The bandwidth setting that you wish to use for this video source. Bandwidth controlled by changing
//	// both the compression level and the resolution of the source. A good use for low bandwidth is working
//	// on WIFI connections.
//	NDIlib_recv_bandwidth_e bandwidth;
//
//	// When this flag is FALSE, all video that you receive will be progressive. For sources that provide
//	// fields, this is de-interlaced on the receiving side (because we cannot change what the up-stream
//	//  source was actually rendering. This is provided as a convenience to down-stream sources that do not
//	// wish to understand fielded video. There is almost no  performance impact of using this function.
//	bool allow_video_fields;
//
//	// The name of the NDI receiver to create. This is a NULL terminated UTF8 string and should be the name
//	// of receive channel that you have. This is in many ways symmetric with the name of senders, so this
//	// might be "Channel 1" on your system. If this is NULL then it will use the filename of your application
//	// indexed with the number of the instance number of this receiver.
//	const char* p_ndi_recv_name;
//
//#if NDILIB_CPP_DEFAULT_CONSTRUCTORS
//	NDIlib_recv_create_v3_t(const NDIlib_source_t source_to_connect_to_ = NDIlib_source_t(), NDIlib_recv_color_format_e color_format_ = NDIlib_recv_color_format_UYVY_BGRA,
//	                        NDIlib_recv_bandwidth_e bandwidth_ = NDIlib_recv_bandwidth_highest, bool allow_video_fields_ = true, const char* p_ndi_name_ = NULL);
//#endif // NDILIB_CPP_DEFAULT_CONSTRUCTORS
//
//} NDIlib_recv_create_v3_t;

type NewRecvCreateV3TOption func(*RecvCreateV3T)

func RecvCreateV3(v RecvCreateV3T) RecvInstanceT {
	t := C.NDIlib_recv_create_v3_t(v)
	return RecvInstanceT(&t)
}
