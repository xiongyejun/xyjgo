package msxls

import (
	"errors"
)

type boolErr struct {
	r record
	//2.4.24	BoolErr
	//The BoolErr record specifies a cell that contains either a Boolean value or an error value.
	cell //  A Cell structure that specifies the cell.

	/*
		2.5.10	Bes
		The Bes structure specifies either a Boolean (section 2.5.14)  value or an error value.
		bBoolErr specifies the value and fError specifies the value’s type.

		bBoolErr (1 byte): An unsigned integer that specifies either a Boolean value or an error value, depending on the value of fError.
		A Boolean value MUST be a value from the following table:
			0x00	False
			0x01	True

		An error value MUST be a value from the following table:
			0x00	#NULL!
			0x07	#DIV/0!
			0x0F	#VALUE!
			0x17	#REF!
			0x1D	#NAME?
			0x24	#NUM!
			0x2A	#N/A
			0x2B	#GETTING_DATA

		fError (1 byte):  A Boolean that specifies whether bBoolErr contains an error code or a Boolean value. MUST be a value from the following table:
			0x00	bBoolErr SHOULD<150> contain a Boolean value.
			0x01	bBoolErr contains an error value.
	*/
	bes []byte // A Bes structure that specifies a Boolean or an error value. size 2
}

func decodeBoolErr(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[number] Size.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &boolErr{
		r: r,
		cell: cell{
			rw:   byte2uint16(buf[:2]),
			col:  byte2uint16(buf[2:4]),
			ixfe: byte2uint16(buf[4:6]),
		},
		bes: buf[6:8],
	}, nil
}

func (me *boolErr) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	c.Row = me.rw
	c.Column = me.col

	if me.bes[1] == 0x00 {
		// Boolean value
		c.Value = (me.bes[0] == 0x01)
	} else {
		// error value
		c.Value = errorValue[me.bes[0]]
	}
	s.MCell[c.RowCol] = len(s.Cells)
	s.Cells = append(s.Cells, c)

	return
}

func (me *boolErr) Type() RecordType {
	return me.r.Type
}
func (me *boolErr) Size() uint16 {
	return me.r.Size
}
