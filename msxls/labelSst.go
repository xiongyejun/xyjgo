package msxls

import (
	"errors"
)

type labelSst struct {
	r record
	// 2.4.149
	// specifies a cell that contains a string
	cell // A Cell structure that specifies the cell containing the string from the shared string table.

	/*
		specifies the zero-based index of an element in the array of XLUnicodeRichExtendedString structure
		in the rgb field of the SST record in this Workbook Stream ABNF that specifies the string contained in the cell.
		MUST be greater than or equal to zero and less than the number of elements in the rgb field of the SST record.
	*/
	isst uint32
}

func decodeLabelSst(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[labelSst] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &labelSst{
		r: r,
		cell: cell{
			rw:   byte2uint16(buf[:2]),
			col:  byte2uint16(buf[2:]),
			ixfe: byte2uint16(buf[4:6]),
		},
		isst: byte2uint32(buf[6:10]),
	}, nil
}

func (me *labelSst) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	c.Row = me.rw
	c.Column = me.col

	// 在sst里的下标
	c.Value = sstString[me.isst]
	s.MCell[c.RowCol] = len(s.Cells)
	s.Cells = append(s.Cells, c)

	return
}

func (me *labelSst) Type() RecordType {
	return me.r.Type
}
func (me *labelSst) Size() uint16 {
	return me.r.Size
}
