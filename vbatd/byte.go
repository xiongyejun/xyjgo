package vbatd

import (
	"errors"
	"unsafe"
)

func getByte(vv VBAVariant) (ret interface{}, err error) {
	return vv.Data[0], nil
}

func getByteFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return *((*byte)(unsafe.Pointer(vv.getPtr()))), nil
}

func getByteArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getByteArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getByteArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getByteArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getByteArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]byte, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*byte)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i)))))
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
