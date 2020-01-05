package msxls

import (
	"errors"
)

type rk struct {
	r record

	//2.4.220	RK
	//The RK record specifies the numeric data contained in a single cell.
	rw    uint16 // An Rw structure that specifies a row index.
	col   uint16 //  A Col structure that specifies a column index.
	rkRec        // An RkRec structure that specifies the numeric data for a single cell.

}

func decodeRk(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[rk] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &rk{
		r:   r,
		rw:  byte2uint16(buf[:2]),
		col: byte2uint16(buf[2:]),
		rkRec: rkRec{
			ixfe:     byte2uint16(buf[4:]),
			RkNumber: buf[6:10],
		},
	}, nil
}

func (me *rk) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	c.Row = me.rw
	c.Column = me.col

	if c.Value, err = getValueFromRkNumber(me.RkNumber); err != nil {
		return
	}
	s.MCell[c.RowCol] = len(s.Cells)
	s.Cells = append(s.Cells, c)

	return
}

func (me *rk) Type() RecordType {
	return me.r.Type
}
func (me *rk) Size() uint16 {
	return me.r.Size
}
