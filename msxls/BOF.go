package msxls

import (
	"errors"
)

type bof struct {
	r record

	Vers uint16 //	BIFF version of the file. The value MUST be 0x0600.

	/*
		0x0005	Specifies the workbook substream.
		0x0010	Specifies the dialog sheet substream or the worksheet substream.
				The sheet substream that starts with this BOF record MUST contain one WsBool record.
				If the fDialog field in that WsBool is 1 then the sheet is dialog sheet
				otherwise the sheet is a worksheet.
		0x0020	Specifies the chart sheet substream.
		0x0040	Specifies the macro sheet substream.
	*/
	Dt       uint16
	RupBuild uint16 // build identifier
	RupYear  uint16 // year when this BIFF version was first created. The value MUST be 0x07CC or 0x07CD

	/*
		A - fWin (1 bit): A bit that specifies whether this file was last edited on a Windows platform. The value MUST be 1.
		B - fRisc (1 bit): A bit that specifies whether the file was last edited on a RISC platform. The value MUST be 0.
		C - fBeta (1 bit): A bit that specifies whether this file was last edited by a beta version of the application. The value MUST be 0.
		D - fWinAny (1 bit): A bit that specifies whether this file has ever been edited on a Windows platform. The value SHOULD<28> be 1.
		E - fMacAny (1 bit): A bit that specifies whether this file has ever been edited on a Macintosh platform. The value MUST be 0.
		F - fBetaAny (1 bit): A bit that specifies whether this file has ever been edited by a beta version of the application. The value MUST be 0.
		G - unused1 (2 bits): Undefined and MUST be ignored.
		H - fRiscAny (1 bit): A bit that specifies whether this file has ever been edited on a RISC platform. The value MUST be 0.
		I - fOOM (1 bit): A bit that specifies whether this file had an out-of-memory failure.
		J - fGlJmp (1 bit): A bit that specifies whether this file had an out-of-memory failure during rendering.
		K - unused2 (2 bits): Undefined, and MUST be ignored.
		L - fFontLimit (1 bit): A bit that specified that whether this file hit the 255 font limit<29>.
		M - verXLHigh (4 bits): An unsigned integer that specifies the highest version of the application that once saved this file. MUST be a value from the following table:
			0x0	Specifies the highest version of the application that has ever saved this file. <30>
			0x1	Specifies the highest version of the application that has ever saved this file. <31>
			0x2	Specifies the highest version of the application that has ever saved this file. <32>
			0x3	Specifies the highest version of the application that has ever saved this file. <33>
			0x4	Specifies the highest version of the application that has ever saved this file. <34>
			0x6	Specifies the highest version of the application that has ever saved this file. <35>
			0x7	Specifies the highest version of the application that has ever saved this file. <36>
		N - unused3 (1 bit): Undefined, and MUST be ignored.
	*/
	A2NReserved1  []byte // size 4
	VerLowestBiff uint8  // BIFF version saved. The value MUST be 6.
	/*
		O - verLastXLSaved (4 bits):An unsigned integer that specifies the application that saved this file most recently. The value MUST be the value of field verXLHigh or less. MUST be a value from the following table:
			0x0	Specifies the highest version of the application that has ever saved this file. <37>
			0x1	Specifies the highest version of the application that has ever saved this file.<38>
			0x2	Specifies the highest version of the application that has ever saved this file.<39>
			0x3	Specifies the highest version of the application that has ever saved this file. <40>
			0x4	Specifies the highest version of the application that has ever saved this file.<41>
			0x6	Specifies the highest version of the application that has ever saved this file.<42>
			0x7	Specifies the highest version of the application that has ever saved this file.<43>
		reserved2 (20 bits): MUST be zero, and MUST be ignored.
	*/
	OReserved2 []byte //size 3
}

func decodeBOF(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[BOF] Size 16.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &bof{
		r:             r,
		Vers:          byte2uint16(buf[0:]),
		Dt:            byte2uint16(buf[2:]),
		RupBuild:      byte2uint16(buf[4:]),
		RupYear:       byte2uint16(buf[6:]),
		A2NReserved1:  buf[8:12],
		VerLowestBiff: buf[12],
		OReserved2:    buf[13:16],
	}, nil
}

func (me *bof) parse(wk *Workbook) (err error) {
	if me.Vers != 0x0600 {
		return errors.New("BIFF version of the file. The value MUST be 0x0600.")
	}
	if me.VerLowestBiff != 6 {
		return errors.New("BIFF version saved. The value MUST be 6.")
	}

	if me.Dt != 0x0005 {
		s := new(Sheet)

		switch me.Dt {
		case 0x0020:
			s.SheetType = SheetChart
		case 0x0040:
			s.SheetType = SheetMacro
		case 0x0010:
			s.SheetType = SheetUndefine
		default:
			err = errors.New("不存在的SheetType.")
			return
		}

		s.Name = shtName[len(wk.Sheets)]
		s.MCell = make(map[RowCol]int, 0)
		wk.Sheets = append(wk.Sheets, s)
	}

	return
}

func (me *bof) Type() RecordType {
	return me.r.Type
}
func (me *bof) Size() uint16 {
	return me.r.Size
}
