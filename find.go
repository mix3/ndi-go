package ndi

/*
#cgo LDFLAGS: -L/usr/local/lib -lndi
#include "include/Processing.NDI.Lib.h"
*/
import "C"
import (
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

func FindGetCurrentSources(instance FindInstanceT) []*SourceT {
	var pNoSources C.uint32_t
	pSources := C.NDIlib_find_get_current_sources(C.NDIlib_find_instance_t(instance), &pNoSources)
	if pNoSources == 0 {
		return nil
	}
	sources := (*[1 << 28]C.NDIlib_source_t)(unsafe.Pointer(pSources))[:pNoSources:pNoSources]
	result := make([]*SourceT, pNoSources)
	for i, source := range sources {
		result[i] = (*SourceT)(&source)
	}
	return result
}

func FindWaitForSources(instance FindInstanceT, timeoutInMS uint32) bool {
	return (bool)(C.NDIlib_find_wait_for_sources(C.NDIlib_find_instance_t(instance), C.uint32_t(timeoutInMS)))
}
