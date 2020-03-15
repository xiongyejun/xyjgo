package vbatd

import (
	"testing"
	"unsafe"
)

func Test_func(t *testing.T) {
	var vv VBAVariant = VBAVariant{}

	vv.Flags[0] = 0x11
	vv.Data[0] = 0xb9
	vv.Data[1] = 0xfc
	vv.Data[2] = 0x87
	vv.Data[3] = 0xf4
	vv.Data[4] = 0x3b
	vv.Data[5] = 0xf3
	vv.Data[6] = 0xb4
	vv.Data[7] = 0x40
	t.Log(vv.Flags[0], vv.Data[0])

	t.Log(Variant2interface(uintptr(unsafe.Pointer(&vv))))

	var f float64 = 5363.2342
	b := *((*[8]byte)(unsafe.Pointer(&f)))
	t.Logf("% x\n", b)

	f = 1.0
	f = *((*float64)(unsafe.Pointer(&vv.Data)))
	t.Log(f)

}
