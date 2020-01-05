package msxls

import (
	"errors"
)

// specifies string constants
type sst struct {
	r record
	// specifies the total number of references in the workbook to the strings in the shared string table.
	// MUST be greater than or equal to 0.
	cstTotal int32
	// specifies the number of unique strings in the shared string table. MUST be greater than or equal to 0.
	cstUnique int32
	//rgb (variable): An array of XLUnicodeRichExtendedString structures.  Records in this array are unique.
	rgb []byte
}

func decodeSST(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[sst] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &sst{
		r:         r,
		cstTotal:  byte2int32(buf[:]),
		cstUnique: byte2int32(buf[4:]),
		rgb:       buf[8:],
	}, nil
}

func (me *sst) parse(wk *Workbook) (err error) {
	sstString = make([]string, me.cstUnique)
	XLUnicodeLen := 0
	index := 0
	for i := range sstString {
		if sstString[i], XLUnicodeLen, err = parseXLUnicodeRichExtendedString(me.rgb[index:]); err != nil {
			return
		}
		index += XLUnicodeLen
	}
	return
}

func (me *sst) Type() RecordType {
	return me.r.Type
}
func (me *sst) Size() uint16 {
	return me.r.Size
}
