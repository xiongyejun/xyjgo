package vbatd

import (
	"errors"
	"unsafe"
)

func getBoolean(vv VBAVariant) (ret interface{}, err error) {
	return vv.Data[0] != 0, nil
}

func getBooleanFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return (*((*int16)(unsafe.Pointer(vv.getPtr())))) != 0, nil
}

func getBooleanArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getBooleanArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getBooleanArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getBooleanArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getBooleanArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]bool, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*int16)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*2))))) != 0
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
