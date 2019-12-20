package msxls

import (
	"errors"
)

// 2.4.127	Formula
// The Formula record specifies a formula (section 2.2.2) for a cell.
type formula struct {
	r record

	cell         //  A Cell structure that specifies a cell on the sheet.
	formulaValue //val (8 bytes):  A FormulaValue structure that specifies the value of the formula.

	/*
		A - fAlwaysCalc (1 bit):
			A bit that specifies whether the formula needs to be calculated during the next recalculation.
		B - reserved1 (1 bit):
			MUST be zero, and MUST be ignored.
		C - fFill (1 bit):
			A bit that specifies whether the cell has a fill alignment or a center-across-selection alignment.
				0	Cell does not have a fill alignment or a center-across-selection alignment.
				1	Cell has either a fill alignment or a center-across-selection alignment.
		D - fShrFmla (1 bit):
			A bit that specifies whether the formula is part of a shared formula as defined in ShrFmla.
			If this formula is part of a shared formula, formula.rgce MUST begin with a PtgExp structure.
		E - reserved2 (1 bit):
			MUST be zero, and MUST be ignored.
		F - fClearErrors (1 bit):
			A bit that specifies whether the formula is excluded from formula error checking.
		reserved3 (10 bits):
			MUST be zero, and MUST be ignored.
	*/
	AFr3 []byte // size 2

	/*
		 A field that specifies an application-specific cache of information.
		This cache exists for performance reasons only,
		and can be rebuilt based on information stored elsewhere in the file without affecting calculation results.
	*/
	chn []byte // size 4
	//	formula (variable): A CellParsedFormula structure that specifies the formula.
}

func decodeFormula(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[formula] Size 2.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &formula{
		cell: cell{
			rw:   byte2uint16(buf[:2]),
			col:  byte2uint16(buf[2:]),
			ixfe: byte2uint16(buf[4:]),
		},
		formulaValue: formulaValue{
			byte1:  buf[6],
			byte2:  buf[7],
			byte3:  buf[8],
			byte4:  buf[9],
			byte5:  buf[10],
			byte6:  buf[11],
			fExprO: buf[12:14],
		},
		AFr3: buf[14:16],
		chn:  buf[16:20],
	}, nil
}

func (me *formula) parse(wk *Workbook) (err error) {
	s := wk.Sheets[len(wk.Sheets)-1]
	c := Cell{}
	c.Row = me.rw
	c.Column = me.col

	if me.fExprO[0] == 0xff && me.fExprO[1] == 0xff {
		switch me.byte1 {
		case 0x00:
			//  The string value is stored in a String record that immediately follows this record
			var r record
			if r, err = decodeRecordHead(wk); err != nil {
				return
			}

			if r.Type != uint16(RecordString) {
				return errors.New(s.Name + " [Formula] 后面，immediately record is not [String].")
			}
			bStringSize := r.Size
			bString := wk.b[wk.pointer : wk.pointer+int(bStringSize)]
			wk.pointer += int(r.Size)

			var str string
			if str, err = parseXLUnicodeString(bString); err != nil {
				return
			}
			c.Value = str
		case 0x01:
			c.Value = (me.byte3 != 0x0) // bool
		case 0x02:
			c.Value = errorValue[me.byte3] // byte3
		case 0x03:
			c.Value = ""
		default:
			c.Value = "未定义的formulaValue"
		}
	} else {
		c.Value = byte2float64([]byte{me.byte1, me.byte2, me.byte3, me.byte4, me.byte5, me.byte6, me.fExprO[0], me.fExprO[1]})
	}

	s.MCell[c.RowCol] = len(s.Cells)
	s.Cells = append(s.Cells, c)
	return nil
}

func (me *formula) Type() RecordType {
	return me.r.Type
}
func (me *formula) Size() uint16 {
	return me.r.Size
}
