package msxls

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/xiongyejun/xyjgo/ucs2"
)

func byte2struct(b []byte, pStruct interface{}) error {
	buf := bytes.NewBuffer(b)
	return binary.Read(buf, binary.LittleEndian, pStruct)
}

func byte2uint16(b []byte) (ret uint16) {
	if b == nil {
		return 0
	}

	var lenb uint32 = uint32(len(b))
	var i uint32 = 0
	for ; i < 2 && i < lenb; i++ {
		ret |= (uint16(b[i]) << (8 * i))
	}
	return
}

func byte2uint32(b []byte) (ret uint32) {
	if b == nil {
		return 0
	}

	var lenb uint32 = uint32(len(b))
	var i uint32 = 0
	for ; i < 4 && i < lenb; i++ {
		ret |= (uint32(b[i]) << (8 * i))
	}
	return
}
func byte2uint64(b []byte) (ret uint64) {
	if b == nil {
		return 0
	}

	var lenb uint32 = uint32(len(b))
	var i uint32 = 0
	for ; i < 8 && i < lenb; i++ {
		ret |= (uint64(b[i]) << (8 * i))
	}
	return
}
func byte2int32(b []byte) (ret int32) {
	if b == nil {
		return 0
	}

	var lenb uint32 = uint32(len(b))
	var i uint32 = 0
	for ; i < 4 && i < lenb; i++ {
		ret |= (int32(b[i]) << (8 * i))
	}
	return
}

func byte2float64(b []byte) float64 {
	if b == nil {
		return 0
	}

	bits := byte2uint64(b)

	return math.Float64frombits(bits)
}

func getValueFromRkNumber(RkNumber []byte) (ret interface{}, err error) {
	fX100 := RkNumber[0] & 0x1 // 0000 0001
	fInt := RkNumber[0] & 0x2  // 0000 0010

	if fInt != 0x2 {
		numUint64 := byte2uint64(RkNumber)
		numUint64 = numUint64 >> 2
		numUint64 = numUint64 << 34
		numFld := math.Float64frombits(numUint64)
		if fX100 == 0x1 {
			numFld /= 100
		}
		ret = numFld
	} else {
		numInt := int16(byte2int32(RkNumber) >> 2)
		if fX100 == 0x1 {
			numInt /= 100
		}
		ret = numInt
	}
	return
}

//2.5.294	XLUnicodeString
//The XLUnicodeString structure specifies a Unicode string.
type XLUnicodeString struct {
	cch uint16 //specifies the count of characters in the string.

	/*
		A - fHighByte (1 bit):
		A bit that specifies whether the characters in rgb are double-byte characters.
		MUST be a value from the following table:
			0x0	All the characters in the string have a high byte of 0x00 and only the low bytes are in rgb.
			0x1	All the characters in the string are saved as double-byte characters in rgb.

		reserved (7 bits): MUST be zero, and MUST be ignored.
		rgb (variable):
		An array of bytes that specifies the characters.
		If fHighByte is 0x0, the size of the array MUST be equal to cch.
		If fHighByte is 0x1, the size of the array MUST be equal to cch*2.
	*/
}

func parseXLUnicodeString(b []byte) (ret string, err error) {
	cch := byte2uint16(b)

	if cch == 0 {
		return "", nil
	}

	fHighByte := b[2] & 0x1 // 0000 0001 取第1个bit
	if fHighByte == 0x0 {
		return string(b[3:]), nil
	}

	var tmp []byte
	if tmp, err = ucs2.ToUTF8(b[3:]); err != nil {
		return
	}
	ret = string(tmp)
	return
}

//2.5.240	ShortXLUnicodeString
//The ShortXLUnicodeString structure specifies a Unicode string
type ShortXLUnicodeString struct {
	cch byte // An unsigned integer that specifies the count of characters in the string.
	/*
		A - fHighByte (1 bit):
		A bit that specifies whether the characters in rgb are double-byte characters.
		MUST be a value from the following table:
			0x0	All the characters in the string have a high byte of 0x00 and only the low bytes are in rgb.
			0x1	All the characters in the string are saved as double-byte characters in rgb.

		reserved (7 bits): MUST be zero, and MUST be ignored.
		rgb (variable):
		An array of bytes that specifies the characters.
		If fHighByte is 0x0, the size of the array MUST be equal to the value of cch.
		If fHighByte is 0x1, the size of the array MUST be equal to the value of cch*2.
	*/
}

func parseShortXLUnicodeString(b []byte) (ret string, err error) {
	cch := byte2uint16([]byte{b[0]})

	if cch == 0 {
		return "", nil
	}

	fHighByte := b[1] & 0x1 // 0000 0001 取第1个bit
	if fHighByte == 0x0 {
		return string(b[2:]), nil
	}

	var tmp []byte
	if tmp, err = ucs2.ToUTF8(b[2:]); err != nil {
		return
	}
	ret = string(tmp)
	return
}

/*
2.5.293	XLUnicodeRichExtendedString
The XLUnicodeRichExtendedString structure specifies a Unicode string,
which can contain formatting information and phonetic string data.
This structure’s non-variable fields MUST be specified in the same record.
This structure’s variable fields can be extended with Continue records.
A value from the table for fHighByte MUST be specified in the first byte of the continue field
of the Continue record followed by the remaining portions of this structure’s variable fields.
*/
type XLUnicodeRichExtendedString struct {
	cch uint16 // specifies the count of characters in the string.

	/*
		A - fHighByte (1 bit): A bit that specifies whether the characters in rgb are double-byte characters. MUST be a value from the following table:
			0x0	All the characters in the string have a high byte of 0x00 and only the low bytes are in rgb.
			0x1	All the characters in the string are saved as double-byte characters in rgb.
		B - reserved1 (1 bit): MUST be zero, and MUST be ignored.
		C - fExtSt (1 bit): A bit that specifies whether the string contains phonetic string data.
		D - fRichSt (1 bit): A bit that specifies whether the string is a rich string and the string has at least two character formats applied.
		reserved2 (4 bits): MUST be zero, and MUST be ignored.
	*/
	ADr2 byte

	// An optional unsigned integer that specifies the number of elements in rgRun. MUST exist if and only if fRichSt is 0x1.
	cRun uint16
	// An optional signed integer that specifies the byte count of ExtRst. MUST exist if and only if fExtSt is 0x1. MUST be zero or greater.
	cbExtRst int32

	/*
		rgb (variable):
			An array of bytes that specifies the characters in the string.
			If fHighByte is 0x0, the size of the array is cch. If fHighByte is 0x1,
			the size of the array is cch*2.
			If fHighByte is 0x1 and rgb is extended with a Continue record the break MUST occur at the
			double-byte character boundary.

		rgRun (variable):
			An optional array of FormatRun structures that specifies the formatting for each text run.
			The number of elements in the array is cRun. MUST exist if and only if fRichSt is 0x1.

		ExtRst (variable):
			An optional ExtRst that specifies the phonetic string data.
			The size of this field is cbExtRst. MUST exist if and only if fExtSt is 0x1.
	*/
}

func parseXLUnicodeRichExtendedString(b []byte) (ret string, retByteLen int, err error) {
	cch := byte2uint16(b)
	ioffset := 2

	if cch == 0 {
		// TODO cch==0的时候是什么原因？
		ioffset += 1
		cch = byte2uint16(b[ioffset:])
		ioffset += 2
	}

	fByte := b[ioffset]
	ioffset += 1
	fHighByte := fByte & 0x1 // 0000 0001 取第1个bit
	if fHighByte == 0x1 {
		cch *= 2
	}
	// D - fRichSt
	var cRun uint16 = 0
	fRichSt := fByte & 0x8 // 0000 1000 取第4个bit
	if fRichSt == 0x8 {
		cRun = byte2uint16(b[ioffset:])
		ioffset += 2
		print("fRichSt\n")
	}

	// C - fExtSt
	var cbExtRst int32 = 0
	fExtSt := fByte & 0x4 // 0000 0100 取第3个bit
	if fExtSt == 0x4 {
		cbExtRst = byte2int32(b[ioffset:])
		ioffset += 4
		print("fExtSt\n")
	}

	var tmp []byte = b[ioffset : ioffset+int(cch)]
	if fHighByte == 0x1 {
		if tmp, err = ucs2.ToUTF8(tmp); err != nil {
			return
		}
	}

	ioffset += int(cRun)
	ioffset += int(cbExtRst)
	ret = string(tmp)
	retByteLen = ioffset + int(cch)

	return
}
