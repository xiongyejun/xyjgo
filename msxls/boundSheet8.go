package msxls

import (
	"errors"
)

type boundSheet8 struct {
	r record

	// A FilePointer as specified in [MS-OSHARED] section 2.2.1.5
	// that specifies the stream position of the start of the BOF record for the sheet.
	LbPlyPos uint32
	/*
		0x00	Visible
		0x01	Hidden
		0x02	Very Hidden; the sheet (1) is hidden and cannot be displayed using the user interface.
	*/
	// 2 bits
	HsState byte
	// Unused (6 bits): Undefined and MUST be ignored
	/*
		0x00	Worksheet or dialog sheet
				The sheet substream that starts with the BOF record specified
				in lbPlyPos MUST contain one WsBool record.
				If the fDialog field in that WsBool is 1 then the sheet is dialog sheet.
				Otherwise, the sheet is a worksheet.
		0x01	Macro sheet
		0x02	Chart sheet
		0x06	VBA module
	*/
	//the sheet type.
	Dt     uint8
	StName []byte // A ShortXLUnicodeString structure that specifies the unique case-insensitive name of the sheet
}

func decodeBoundSheet8(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[BoundSheet8] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &boundSheet8{
		r:        r,
		LbPlyPos: byte2uint32(buf[:]),
		HsState:  buf[4],
		Dt:       buf[5],
		StName:   buf[6:],
	}, nil
}

func (me *boundSheet8) parse(wk *Workbook) (err error) {
	var tmp string
	if tmp, err = parseShortXLUnicodeString(me.StName); err != nil {
		return
	}
	shtName = append(shtName, tmp)

	return
}

func (me *boundSheet8) Type() RecordType {
	return me.r.Type
}
func (me *boundSheet8) Size() uint16 {
	return me.r.Size
}
