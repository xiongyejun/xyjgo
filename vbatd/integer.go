package vbatd

import (
	"errors"
	"unsafe"
)

func getInteger(vv VBAVariant) (ret interface{}, err error) {
	return int16(vv.Data[0]) | (int16(vv.Data[1]) << 8), nil
}

func getIntegerFromPtr(vv VBAVariant) (ret interface{}, err error) {
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	return *((*int16)(unsafe.Pointer(vv.getPtr()))), nil
}

func getIntegerArrFromPtr(vv VBAVariant) (ret interface{}, err error) {
	return getIntegerArr(vv.getPtr(), sizeVBADataType[vv.Flags[0]])
}

func getIntegerArrFromPtrPtr(vv VBAVariant) (ret interface{}, err error) {
	pv := *((*int32)(unsafe.Pointer(vv.getPtr())))

	return getIntegerArr(uintptr(pv), sizeVBADataType[vv.Flags[0]])
}

func getIntegerArr(pv uintptr, typeSize uint32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(pv)))

	if safearr.cbElemets != typeSize {
		err = errors.New("SafeArray.cbElemets元素的字节大小与VBA Variant Type不一致.")
		return
	}
	if safearr.cDims == 1 {
		b := make([]int16, safearr.rgsabound[0].cElemets)
		for i := range b {
			b[i] = (*(*int16)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*2)))))
		}
		ret = b
	} else {
		err = errors.New("SafeArray.cDims != 1，未处理的情况.")
		return
	}

	return
}
