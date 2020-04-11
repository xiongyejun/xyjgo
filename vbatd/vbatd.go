// vba type of data vba数据类型
package vbatd

import (
	"errors"
	"unsafe"
)

// https://docs.microsoft.com/en-us/office/vba/language/reference/user-interface-help/vartype-function
// VBA Variant第1个字节
const (
	VInteger = 0x2
	VLong    = 0x3
	VSingle  = 0x4
	VDouble  = 0x5
	VDate    = 0x7
	VString  = 0x8
	VObject  = 0x9
	VBoolean = 0xb
	VByte    = 0x11

	// VBAVariant占用字节
	SIZE_VBAVariant = 0x10

	// VBAVariant b1代表的内容
	IS_VAL         = 0x00 // 数据本身，但String还是地址
	IS_ARR_PTR     = 0x20 // 是指向Arr的地址
	IS_VAR_PTR     = 0x40 // 是指向变量的地址
	IS_ARR_PTR_PTR = 0x60 // 是指向Arr的地址的地址
)

type SafeArrayBound struct {
	cElemets uint32 // 该维的长度
	lLbound  uint32 // 该维的数组存取的下限，一般为0
}

type SafeArray struct {
	cDims     uint16 // 数组的维度
	fFeatures uint16
	cbElemets uint32 // 数组元素的字节大小
	cLocksas  uint32
	pvDataas  int32 // 数组的数据指针
	rgsabound [1]SafeArrayBound
}

type VBAVariant struct {
	// b0  存储的数据类型

	/*
		b1
			0x00	数据类型的是数据本身，String是地址
			0x40	8-11存的是数据地址，String是地址的地址
			0x20	8-11存的是数组地址
			0x60	8-11存的是数组地址的地址
	*/

	Flags [8]byte
	Data  [8]byte
}

// VBAVariant 如果保存的是地址，在8-11位
func (me *VBAVariant) getPtr() uintptr {
	return uintptr(int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24))
}

// VBAVariant第1个字节表示的数据类型的字节大小
var sizeVBADataType []uint32

type getValFunc struct {
	b0 byte
	b1 byte
}

// 用map记录b0、b1所对应的func
var mGetVal map[getValFunc]func(vv VBAVariant) (interface{}, error)

func init() {
	sizeVBADataType = make([]uint32, VByte+1)
	sizeVBADataType[VByte] = 1
	sizeVBADataType[VInteger] = 2
	sizeVBADataType[VDouble] = 8
	sizeVBADataType[VSingle] = 4
	sizeVBADataType[VDate] = 8
	sizeVBADataType[VLong] = 4
	sizeVBADataType[VBoolean] = 2
	sizeVBADataType[VString] = 4

	mGetVal = make(map[getValFunc]func(vv VBAVariant) (interface{}, error))
	// 初始化mGetVal
	mGetVal[getValFunc{b0: VByte, b1: IS_VAL}] = getByte
	mGetVal[getValFunc{b0: VInteger, b1: IS_VAL}] = getInteger
	mGetVal[getValFunc{b0: VLong, b1: IS_VAL}] = getLong
	mGetVal[getValFunc{b0: VBoolean, b1: IS_VAL}] = getBoolean
	mGetVal[getValFunc{b0: VSingle, b1: IS_VAL}] = getSingle
	mGetVal[getValFunc{b0: VDouble, b1: IS_VAL}] = getDouble
	mGetVal[getValFunc{b0: VString, b1: IS_VAL}] = getString

	mGetVal[getValFunc{b0: VByte, b1: IS_VAR_PTR}] = getByteFromPtr
	mGetVal[getValFunc{b0: VInteger, b1: IS_VAR_PTR}] = getIntegerFromPtr
	mGetVal[getValFunc{b0: VLong, b1: IS_VAR_PTR}] = getLongFromPtr
	mGetVal[getValFunc{b0: VBoolean, b1: IS_VAR_PTR}] = getBooleanFromPtr
	mGetVal[getValFunc{b0: VSingle, b1: IS_VAR_PTR}] = getSingleFromPtr
	mGetVal[getValFunc{b0: VDouble, b1: IS_VAR_PTR}] = getDoubleFromPtr
	mGetVal[getValFunc{b0: VString, b1: IS_VAR_PTR}] = getStringFromPtr

	mGetVal[getValFunc{b0: VByte, b1: IS_ARR_PTR}] = getIntegerArrFromPtr
	mGetVal[getValFunc{b0: VInteger, b1: IS_ARR_PTR}] = getIntegerArrFromPtr
	mGetVal[getValFunc{b0: VLong, b1: IS_ARR_PTR}] = getLongArrFromPtr
	mGetVal[getValFunc{b0: VBoolean, b1: IS_ARR_PTR}] = getBooleanArrFromPtr
	mGetVal[getValFunc{b0: VSingle, b1: IS_ARR_PTR}] = getSingleArrFromPtr
	mGetVal[getValFunc{b0: VDouble, b1: IS_ARR_PTR}] = getDoubleArrFromPtr
	mGetVal[getValFunc{b0: VString, b1: IS_ARR_PTR}] = getStringArrFromPtr

	mGetVal[getValFunc{b0: VByte, b1: IS_ARR_PTR_PTR}] = getByteArrFromPtrPtr
	mGetVal[getValFunc{b0: VInteger, b1: IS_ARR_PTR_PTR}] = getIntegerArrFromPtrPtr
	mGetVal[getValFunc{b0: VLong, b1: IS_ARR_PTR_PTR}] = getLongArrFromPtrPtr
	mGetVal[getValFunc{b0: VBoolean, b1: IS_ARR_PTR_PTR}] = getBooleanArrFromPtrPtr
	mGetVal[getValFunc{b0: VSingle, b1: IS_ARR_PTR_PTR}] = getSingleArrFromPtrPtr
	mGetVal[getValFunc{b0: VDouble, b1: IS_ARR_PTR_PTR}] = getDoubleArrFromPtrPtr
	mGetVal[getValFunc{b0: VString, b1: IS_ARR_PTR_PTR}] = getStringArrFromPtrPtr
}

// VBAVariant转换为interface
func Variants2interfaces(pParamArray uintptr, nCount int) (ret []interface{}, err error) {
	ret = make([]interface{}, nCount)
	for i := 0; i < nCount; i++ {
		if ret[i], err = Variant2interface(pParamArray); err != nil {
			return
		}
		pParamArray += SIZE_VBAVariant
	}

	return
}

// 将VBA传入的Variant指针转换为go的interface{}
func Variant2interface(pVBAVariant uintptr) (ret interface{}, err error) {
	vv := (*(*VBAVariant)(unsafe.Pointer(pVBAVariant)))

	if f, ok := mGetVal[getValFunc{b0: vv.Flags[0], b1: vv.Flags[1]}]; ok {
		return f(vv)
	}

	err = errors.New("不能处理的VBA Variant Type.")
	return
}
