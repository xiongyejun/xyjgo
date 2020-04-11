package vbatd

import (
	"errors"
	"unsafe"

	"github.com/xiongyejun/xyjgo/ucs2"
)

func getString(vv VBAVariant) (ret interface{}, err error) {
	vs := VBAString{
		lenth: *((*int32)(unsafe.Pointer(vv.getPtr() - 4))),
		ptr:   vv.getPtr(),
	}
	return vs.getGoString()

}

func getStringFromPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := uintptr(*((*int32)(unsafe.Pointer(vv.getPtr()))))

	vs := VBAString{
		lenth: *((*int32)(unsafe.Pointer(pv - 4))),
		ptr:   pv,
	}
	return vs.getGoString()

}

func getStringArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getStringArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getStringArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getStringArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getStringArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]string, safearr.rgsabound[0].cElemets)
		for i := range b {
			pv = uintptr(*((*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4))))))

			if pv == 0 {
				b[i] = ""
				continue
			}

			vs := VBAString{
				lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
				ptr:   uintptr(pv),
			}

			if b[i], err = vs.getGoString(); err != nil {
				return
			}
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}

type VBAString struct {
	lenth int32   // ptr - 4
	ptr   uintptr // Variant的8-11
}

func (me *VBAString) getGoString() (ret string, err error) {
	b := make([]byte, me.lenth)
	for i := range b {
		b[i] = (*(*byte)(unsafe.Pointer(me.ptr + uintptr(i))))
	}

	if b, err = ucs2.ToUTF8(b); err != nil {
		return
	}

	ret = string(b)

	return
}
