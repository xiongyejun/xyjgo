package msxls

import (
	"errors"
)

type mulRk struct {
	r record
	// 2.4.175	MulRk
	// The MulRk record specifies a series of cells with numeric data in a sheet row.
	// This record can store up to 256 RkRec structures.
	rw       uint16 //  An Rw structure that specifies the row containing the cells with numeric data.
	colFirst uint16 //   A Col structure that specifies the first column in the series of numeric cells within the sheet. The value of colFirst.col MUST be less than or equal to 254.
	/*
	   An array of RkRec structures.
	   Each element in the array specifies an RkRec in the row.
	   The number of entries in the array MUST be equal to the value given by the following formula:
	   	Number of entries in rgrkrec = (colLast.col – colFirst.col +1)
	*/
	rgrkrec []byte
	colLast uint16 //   A Col structure that specifies the last column in the set of numeric cells within the sheet. This colLast.col value MUST be greater than the colFirst.col value.

}

func decodeMulRk(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[mulRk] Size 2.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &mulRk{
		r:        r,
		rw:       byte2uint16(buf[:]),
		colFirst: byte2uint16(buf[2:]),
		rgrkrec:  buf[r.Size-4 : r.Size-2],
		colLast:  byte2uint16(buf[r.Size-2:]),
	}, nil
}

func (me *mulRk) parse(wk *Workbook) (err error) {
	rgrkrecSize := len(me.rgrkrec)
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	for j := 0; j < int(rgrkrecSize/6); j++ {
		c.Row = me.rw
		c.Column = me.colFirst + uint16(j)

		RkNumber := me.rgrkrec[j*6 : j*6+2+4]
		if c.Value, err = getValueFromRkNumber(RkNumber); err != nil {
			return
		}
		s.MCell[c.RowCol] = len(s.Cells)
		s.Cells = append(s.Cells, c)
	}

	return
}

func (me *mulRk) Type() RecordType {
	return me.r.Type
}
func (me *mulRk) Size() uint16 {
	return me.r.Size
}
