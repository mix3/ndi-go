package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"
import (
	"log"
	"unsafe"
)

type FindInstanceT C.NDIlib_find_instance_t

type FindCreateT C.NDIlib_find_create_t

func NewFindCreateT(opts ...NewFindCreateTOption) FindCreateT {
	v := FindCreateT{
		show_local_sources: true,
	}
	for _, opt := range opts {
		opt(&v)
	}
	return v
}

type NewFindCreateTOption func(*FindCreateT)

func NewFindCreateTOptionShowLocalSources(v bool) NewFindCreateTOption {
	return func(t *FindCreateT) {
		t.show_local_sources = C.bool(v)
	}
}

func NewFindCreateTOptionGroups(v *string) NewFindCreateTOption {
	return func(t *FindCreateT) {
		if v == nil {
			t.p_groups = nil
		} else {
			t.p_groups = C.CString(*v)
		}
	}
}

func NewFindCreateTOptionExtraIPs(v *string) NewFindCreateTOption {
	return func(t *FindCreateT) {
		if v == nil {
			t.p_extra_ips = nil
		} else {
			t.p_extra_ips = C.CString(*v)
		}
	}
}

func FindCreateV2(v FindCreateT) FindInstanceT {
	t := C.NDIlib_find_create_t(v)
	return FindInstanceT(C.NDIlib_find_create_v2(&t))
}

func FindDestroy(t FindInstanceT) {
	C.NDIlib_find_destroy(C.NDIlib_find_instance_t(t))
}

//E>const NDIlib_source_t* NDIlib_find_get_current_sources(NDIlib_find_instance_t p_instance, uint32_t* p_no_sources);

type SourceT struct {
	ref *C.NDIlib_source_t
}

func (t SourceT) Name() string {
	return C.GoString(t.ref.p_ndi_name)
}

func FindGetCurrentSources(instance FindInstanceT) []*SourceT {
	var pNoSources C.uint32_t
	ret := C.NDIlib_find_get_current_sources(C.NDIlib_find_instance_t(instance), &pNoSources)
	if pNoSources == 0 {
		return nil
	}
	slices := (*[1 << 28]*C.NDIlib_source_t)(unsafe.Pointer(ret))[:pNoSources:pNoSources]
	result := make([]*SourceT, pNoSources)
	for i, v := range slices {
		vv := *v
		log.Println(C.GoString(vv.p_ndi_name))
		result[i] = &SourceT{ref: v}
	}
	return result
}

func FindWaitForSources(instance FindInstanceT, timeoutInMS uint32) bool {
	return (bool)(C.NDIlib_find_wait_for_sources(C.NDIlib_find_instance_t(instance), C.uint32_t(timeoutInMS)))
}
