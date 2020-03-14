// vba type of data vba数据类型
package vbatd

import (
	"errors"
	"unsafe"

	"github.com/xiongyejun/xyjgo/ucs2"
)

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
)

type VBAString struct {
	lenth int32 // ptr - 4
	ptr   int32 // Variant的8-11
}
type VBAVariant struct {
	VType int64
	Data  []byte // 16 - 8
}

func Variants2interfaces(pParamArray uintptr, nCount int) (ret []interface{}, err error) {
	ret = make([]interface{}, nCount)
	for i := 0; i < nCount; i++ {
		if ret[i], err = Variant2interface(pParamArray); err != nil {
			return
		}
		pParamArray += 16
	}

	return
}

// 将VBA传入的Variant指针转换为go的interface{}
func Variant2interface(pVBAVariant uintptr) (ret interface{}, err error) {
	b := (*(*[16]byte)(unsafe.Pointer(pVBAVariant)))[:]

	switch b[0] {
	case VInteger, VLong, VByte:
		ret = cint32(b)
	case VString:
		pVBAString := int32(b[8]) |
			(int32(b[9]) << 8) |
			(int32(b[10]) << 16) |
			(int32(b[11]) << 24)

		vs := VBAString{
			lenth: *((*int32)(unsafe.Pointer(uintptr(pVBAString - 4)))),
			ptr:   pVBAString,
		}

		if vs.lenth < 1024 {
			b = (*(*[1024]byte)(unsafe.Pointer(uintptr(vs.ptr))))[:vs.lenth]
		} else if vs.lenth < 1024*3 {
			b = (*(*[1024 * 3]byte)(unsafe.Pointer(uintptr(vs.ptr))))[:vs.lenth]
		} else {
			err = errors.New("字符串太长了.")
		}
		if b, err = ucs2.ToUTF8(b); err != nil {
			return
		}
		ret = string(b)
	case VSingle:
		b = b[8:12]
		ret = *((*float32)(unsafe.Pointer(&b)))
	case VDouble:
		b = b[8:16]
		ret = *((*float64)(unsafe.Pointer(&b)))
	case VBoolean:
		ret = (b[8] != 0)
	default:
		err = errors.New("不能处理的VBA Variant Type.")
	}
	return
}

func cint32(b []byte) (ret interface{}) {
	i := int32(b[8]) |
		(int32(b[9]) << 8) |
		(int32(b[10]) << 16) |
		(int32(b[11]) << 24)

	ret = i

	return
}
