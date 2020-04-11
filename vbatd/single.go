package vbatd

import (
	"errors"
	"unsafe"
)

func getSingle(vv VBAVariant) (ret interface{}, err error) {
	return *((*float32)(unsafe.Pointer(&vv.Data))), nil
}

func getSingleFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return *((*float32)(unsafe.Pointer(vv.getPtr()))), nil
}

func getSingleArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getSingleArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getSingleArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getSingleArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getSingleArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]float32, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*float32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
