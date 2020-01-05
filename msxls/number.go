package msxls

import (
	"errors"
)

type number struct {
	r record
	//2.4.180	Number
	//The Number record specifies a cell that contains a floating-point number.
	cell //  A Cell structure that specifies the cell.

	/*
		2.5.342	Xnum
		Xnum is a 64-bit binary floating-point number as specified in [IEEE754].
		This value MUST NOT<191> be infinity, denormalized, not-a-number (NaN), nor negative zero.

		An Xnum (section 2.5.342) value that specifies the cell value.
		If this record appears in a SERIESDATA record collection,
		and this record specifies a cell in the chart data cache that specifies data for an error bar series,
		then this field is a ChartNumNillable value.
		If a ChartNumNillable is used, a blank cell is specified by a NilChartNum structure that
		has a type field with a value of 0x0000,
		and a cell with a #N/A error is specified by a NilChartNum that has a type field with a value of 0x0100.
	*/
	num []byte //size 8
}

func decodeNumber(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[number] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &number{
		r: r,
		cell: cell{
			rw:   byte2uint16(buf[:2]),
			col:  byte2uint16(buf[2:]),
			ixfe: byte2uint16(buf[4:6]),
		},
		num: buf[6:14],
	}, nil
}

func (me *number) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	c.Row = me.rw
	c.Column = me.col

	c.Value = byte2float64(me.num)
	s.MCell[c.RowCol] = len(s.Cells)
	s.Cells = append(s.Cells, c)

	return
}

func (me *number) Type() RecordType {
	return me.r.Type
}
func (me *number) Size() uint16 {
	return me.r.Size
}
