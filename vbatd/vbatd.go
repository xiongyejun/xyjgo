// vba type of data vba数据类型
package vbatd

import (
	"errors"
	"unsafe"

	"github.com/xiongyejun/xyjgo/ucs2"
)

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

	SIZE_VBAVariant  = 0x10
	IS_ARR_ADDR      = 0x20
	IS_VAR_ADDR      = 0x40
	IS_ARR_ADDR_ADDR = 0x60
)

type SafeArrayBound struct {
	cElements uint32 // 该维的长度
	lLbound   uint32 // 该维的数组存取的下限，一般为0
}

type SafeArray struct {
	cDims      uint16 // 数组的维度
	fFeatures  uint16
	cbElements uint32 // 数组元素的字节大小
	cLocksas   uint32
	pvDataas   int32 // 数组的数据指针
	rgsabound  [2]SafeArrayBound
}

type VBAString struct {
	lenth int32   // ptr - 4
	ptr   uintptr // Variant的8-11
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

var sizeVBADataType []uint32

func init() {
	// VBAVariant第1个字节表示的数据类型的字节大小
	sizeVBADataType = make([]uint32, VByte+1)
	sizeVBADataType[VByte] = 1
	sizeVBADataType[VInteger] = 2
	sizeVBADataType[VDouble] = 8
	sizeVBADataType[VSingle] = 4
	sizeVBADataType[VDate] = 8
	sizeVBADataType[VLong] = 4
	sizeVBADataType[VBoolean] = 2
	sizeVBADataType[VString] = 4
}
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

	switch vv.Flags[1] {
	case 0x0:
		return vv.getValue()

	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		return vv.getValueAddr()

	// 存储的是数组
	case IS_ARR_ADDR:
		var pv int32 = int32(vv.Data[0]) | (int32(vv.Data[1]) << 8) | (int32(vv.Data[2]) << 16) | (int32(vv.Data[3]) << 24)
		return vv.getArr(pv)

	// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
	case IS_ARR_ADDR_ADDR:
		var pv int32 = int32(vv.Data[0]) | (int32(vv.Data[1]) << 8) | (int32(vv.Data[2]) << 16) | (int32(vv.Data[3]) << 24)
		if vv.Flags[0] != VString {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		return vv.getArr(pv)

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}
	return
}

func (me *VBAVariant) getValue() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[0] {
	case VByte:
		ret = me.Data[0]

	case VInteger:
		ret = int16(me.Data[0]) | (int16(me.Data[1]) << 8)

	case VLong:
		ret = pv

	case VString:
		vs := VBAString{
			lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
			ptr:   uintptr(pv),
		}
		return vs.getGoString()

	case VSingle:
		ret = *((*float32)(unsafe.Pointer(&me.Data)))
	case VDouble:
		ret = *((*float64)(unsafe.Pointer(&me.Data)))
	case VBoolean:
		ret = (me.Data[0] != 0)
	default:
		err = errors.New("不能处理的VBA Variant Type.")
	}
	return
}

func (me *VBAVariant) getValueAddr() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[0] {
	case VByte:
		ret = *((*byte)(unsafe.Pointer(uintptr(pv))))

	case VInteger:
		ret = *((*int16)(unsafe.Pointer(uintptr(pv))))

	case VLong:
		ret = *((*int32)(unsafe.Pointer(uintptr(pv))))

	case VString:
		pv = *((*int32)(unsafe.Pointer(uintptr(pv))))

		vs := VBAString{
			lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
			ptr:   uintptr(pv),
		}

		b := make([]byte, vs.lenth)
		for i := range b {
			b[i] = (*(*byte)(unsafe.Pointer(vs.ptr + uintptr(i))))
		}

		if b, err = ucs2.ToUTF8(b); err != nil {
			return
		}
		ret = string(b)
	case VBoolean:
		tmp := *((*int16)(unsafe.Pointer(uintptr(pv))))
		ret = (tmp != 0)

	default:
		err = errors.New("不能处理的VBA Variant Type.")
	}

	return
}

func (me *VBAVariant) getArr(pv int32) (ret interface{}, err error) {
	safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

	if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
		err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
	}

	switch me.Flags[0] {
	case VByte:
		if safearr.cDims == 1 {
			b := make([]byte, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*byte)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i)))))
			}
			ret = b
		}

	case VInteger:
		if safearr.cDims == 1 {
			b := make([]int16, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*int16)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*2)))))
			}
			ret = b
		}

	case VLong:
		if safearr.cDims == 1 {
			b := make([]int32, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))
			}
			ret = b
		}
	case VString:
		if safearr.cDims == 1 {
			b := make([]string, safearr.rgsabound[0].cElements)
			for i := range b {
				pv = *((*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))

				vs := VBAString{
					lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
					ptr:   uintptr(pv),
				}

				if b[i], err = vs.getGoString(); err != nil {
					return
				}

			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant Type.")
	}

	return
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
