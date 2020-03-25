// vba type of data vba数据类型
package vbatd

import (
	"errors"
	"unsafe"

	"github.com/xiongyejun/xyjgo/ucs2"
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
	IS_ARR_ADDR      = 0x20 // 是指向Arr的地址
	IS_VAR_ADDR      = 0x40 // 是指向数据的地址
	IS_ARR_ADDR_ADDR = 0x60 // 是指向Arr的地址的地址
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

	switch vv.Flags[0] {
	case VByte:
		return vv.getByte()

	case VInteger:
		return vv.getInteger()

	case VLong:
		return vv.getLong()

	case VString:
		return vv.getString()

	case VSingle:
		return vv.getSingle()

	case VDouble:
		return vv.getDouble()

	case VBoolean:
		return vv.getBoolean()

	default:
		err = errors.New("不能处理的VBA Variant Type.")
	}
	return
}

func (me *VBAVariant) getByte() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = me.Data[0]
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = *((*byte)(unsafe.Pointer(uintptr(pv))))
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]byte, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*byte)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i)))))
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getInteger() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = int16(me.Data[0]) | (int16(me.Data[1]) << 8)
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = *((*int16)(unsafe.Pointer(uintptr(pv))))
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]int16, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*int16)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*2)))))
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getBoolean() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = me.Data[0] != 0
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = (*((*int16)(unsafe.Pointer(uintptr(pv))))) != 0
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]bool, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*int16)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*2))))) != 0
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getLong() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = pv
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = *((*int32)(unsafe.Pointer(uintptr(pv))))
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]int32, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getSingle() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = *((*float32)(unsafe.Pointer(&me.Data)))
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = *((*float32)(unsafe.Pointer(uintptr(pv))))
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]float32, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*float32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getDouble() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		ret = *((*float64)(unsafe.Pointer(&me.Data)))
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		ret = *((*float64)(unsafe.Pointer(uintptr(pv))))
	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
		}
		if safearr.cDims == 1 {
			b := make([]float64, safearr.rgsabound[0].cElements)
			for i := range b {
				b[i] = (*(*float64)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*8)))))
			}
			ret = b
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
	}

	return
}

func (me *VBAVariant) getString() (ret interface{}, err error) {
	var pv int32 = int32(me.Data[0]) | (int32(me.Data[1]) << 8) | (int32(me.Data[2]) << 16) | (int32(me.Data[3]) << 24)

	switch me.Flags[1] {
	case 0x0:
		vs := VBAString{
			lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
			ptr:   uintptr(pv),
		}
		return vs.getGoString()
	// 通过ParamArray传递参数，Variant里保存的是数据地址
	case IS_VAR_ADDR:
		pv = *((*int32)(unsafe.Pointer(uintptr(pv))))

		vs := VBAString{
			lenth: *((*int32)(unsafe.Pointer(uintptr(pv - 4)))),
			ptr:   uintptr(pv),
		}
		return vs.getGoString()

	// 存储的是数组
	case IS_ARR_ADDR, IS_ARR_ADDR_ADDR:
		// 通过ParamArray传递的数组，Variant里保存的地址指向数组的地址
		if me.Flags[1] == IS_ARR_ADDR_ADDR {
			pv = *((*int32)(unsafe.Pointer(uintptr(pv))))
		}
		safearr := *((*SafeArray)(unsafe.Pointer(uintptr(pv))))

		if safearr.cbElements != sizeVBADataType[me.Flags[0]] {
			err = errors.New("SafeArray.cbElements元素的字节大小与VBA Variant Type不一致.")
			return
		}

		if safearr.cDims == 1 {
			b := make([]string, safearr.rgsabound[0].cElements)
			for i := range b {
				pv = *((*int32)(unsafe.Pointer(uintptr(safearr.pvDataas + int32(i*4)))))

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
		}

	default:
		err = errors.New("不能处理的VBA Variant 第2个byte Type.")
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
