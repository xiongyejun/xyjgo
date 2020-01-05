package msxls

import (
	"errors"
)

type wsBool struct {
	r record
	/*
		A - fShowAutoBreaks (1 bit):  A bit that specifies whether page breaks inserted automatically are visible on the sheet.
		B - reserved1 (3 bits):  MUST be zero, and MUST be ignored.
		C - fDialog (1 bit):  A bit that specifies whether the sheet is a dialog sheet.
		D - fApplyStyles (1 bit):  A bit that specifies whether to apply styles in an outline when an outline is applied.
		E - fRowSumsBelow (1 bit):  A bit that specifies whether summary rows appear below an outline's detail rows.
		F - fColSumsRight (1 bit):  A bit that specifies whether summary columns appear to the right or left of an outline's detail columns. Valid values are specified in the following table:
			0	The summary columns appear to the right, if the sheet is displayed left-to-right, or appear to the left, if the sheet is displayed right-to-left.
			1	The summary columns appear to the left, if the sheet is displayed left-to-right, or appear to the right, if the sheet is displayed right-to-left.
		G - fFitToPage (1 bit):  A bit that specifies whether to fit the printable contents to a single page when printing this sheet.
		H - reserved2 (1 bit):  MUST be zero, and MUST be ignored.
		I - unused (2 bits): Undefined and MUST be ignored.
		J - fSyncHoriz (1 bit):  A bit that specifies whether horizontal scrolling is synchronized across multiple windows displaying this sheet.
		K - fSyncVert (1 bit):  A bit that specifies whether vertical scrolling is synchronized across multiple windows displaying this sheet.
		L - fAltExprEval (1 bit):  A bit that specifies whether the sheet uses transition formula evaluation.
		M - fAltFormulaEntry (1 bit):  A bit that specifies whether the sheet uses transition formula entry
	*/
	A2M []byte // size 2
}

func decodeWsBool(r record, wk *Workbook) (retI IRecord, err error) {
	buf := wk.b[wk.pointer:]
	if len(buf) < int(r.Size) {
		err = errors.New("不足[WsBool] Size 2.")
		return
	}
	buf = buf[:r.Size]
	wk.pointer += int(r.Size)

	return &wsBool{
		r:   r,
		A2M: buf[:2],
	}, nil
}

func (me *wsBool) parse(wk *Workbook) (err error) {
	fDialog := me.A2M[0] & 0x10 // 0x0001 0000

	s := wk.Sheets[len(wk.Sheets)-1]
	if fDialog == 0x10 {
		s.SheetType = SheetDialog
	} else {
		s.SheetType = SheetWorksheet
		ws := new(Worksheet)
		ws.Sheet = s
		wk.Worksheets = append(wk.Worksheets, ws)
	}

	return
}

func (me *wsBool) Type() RecordType {
	return me.r.Type
}
func (me *wsBool) Size() uint16 {
	return me.r.Size
}
