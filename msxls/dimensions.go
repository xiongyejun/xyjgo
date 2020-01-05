package msxls

import (
	"errors"
)

type dimensions struct {
	r record
	/*
		The Dimensions record specifies the used range of the sheet. It specifies the row and column bounds of used cells in the sheet. Used cells include all cells with formulas (section 2.2.2) or data. Used cells also include all cells with formatting applied directly to the cell. Cells can also be formatted by default row or column formatting. If a row has default formatting then the used range includes that row in its row bounds, but does not affect the used range column bounds, unless the used range would otherwise be empty, in which case the column bounds are set to include the first column. If a column has default formatting then the used range includes that column in its column bounds, but does not affect the used range row bounds, unless the used range would otherwise be empty, in which case the row bounds are set to include the first row.
	*/
	rwMic    uint32 //  A RwLongU structure that specifies the first row in the sheet that contains a used cell.
	rwMac    uint32 //  An unsigned integer that specifies the zero-based index of the row after the last row in the sheet that contains a used cell. MUST be less than or equal to 0x00010000. If this value is 0x00000000, no cells on the sheet are used cells.
	colMic   uint16 //  A ColU structure that specifies the first column in the sheet that contains a used cell.
	colMac   uint16 // An unsigned integer that specifies the zero-based index of the column after the last column in the sheet that contains a used cell. MUST be less than or equal to 0x0100. If this value is 0x0000, no cells on the sheet are used cells.
	reserved uint16 // MUST be zero, and MUST be ignored.
}

func decodeDimensions(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[dimensions] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &dimensions{
		r:        r,
		rwMic:    byte2uint32(buf[:]),
		rwMac:    byte2uint32(buf[4:]),
		colMic:   byte2uint16(buf[8:]),
		colMac:   byte2uint16(buf[10:]),
		reserved: byte2uint16(buf[12:]),
	}, nil
}

func (me *dimensions) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]

	s.first.Row = uint16(me.rwMic)
	s.last.Row = uint16(me.rwMac) - 1
	s.first.Column = me.colMic
	s.last.Column = me.colMac - 1

	return
}

func (me *dimensions) Type() RecordType {
	return me.r.Type
}
func (me *dimensions) Size() uint16 {
	return me.r.Size
}
