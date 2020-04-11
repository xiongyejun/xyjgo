package vbatd

import (
	"errors"
	"unsafe"
)

func getDouble(vv VBAVariant) (ret interface{}, err error) {
	return *((*float64)(unsafe.Pointer(&vv.Data))), nil
}

func getDoubleFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return *((*float64)(unsafe.Pointer(vv.getPtr()))), nil
}

func getDoubleArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getDoubleArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getDoubleArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getDoubleArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getDoubleArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]float64, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*float64)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*8)))))
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
