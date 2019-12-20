package msxls

import (
	"errors"
	"sort"
	"strconv"

	"github.com/xiongyejun/xyjgo/compoundFile"
)

type MSXLS struct {
	Src []byte
	cf  *compoundFile.CompoundFile
}

type Workbook struct {
	b             []byte
	pointer       int
	mDecodeRecord map[RecordType]func(record, *Workbook) (IRecord, error)

	Sheets     []*Sheet
	Worksheets []*Worksheet
}

type Sheet struct {
	Index int
	Name  string
	SheetType

	Cells []Cell

	// 根据行列来获取Cell，记录的是Cells下标
	MCell       map[RowCol]int
	first, last RowCol
}

type Worksheet struct {
	*Sheet
}

type RowCol struct {
	Row, Column uint16
}

type Cell struct {
	Value interface{}
	RowCol
}

type IRecord interface {
	Type() RecordType
	Size() uint16
	parse(wk *Workbook) (err error)
}

// 2.1.4
type record struct {
	Type uint16
	Size uint16 // the total size of the record data. >= 0 and <= 8224.
	//	Data []byte
}

type SheetType = int

const (
	HEADSIZE uint16 = 4

	SheetDialog SheetType = iota
	SheetWorksheet
	SheetChart
	SheetMacro
	SheetUndefine
)

var errorValue []string = make([]string, 0x2B+1)
var sstString []string
var shtName []string // 临时保存shtName，因为boundSheet8是在sheet解析前出现的

func init() {
	errorValue[0x00] = "#NULL!"
	errorValue[0x07] = "#DIV/0!"
	errorValue[0x0F] = "#VALUE!"
	errorValue[0x17] = "#REF!"
	errorValue[0x1D] = "#NAME?"
	errorValue[0x24] = "#NUM!"
	errorValue[0x2A] = "#N/A"
	errorValue[0x2B] = "#GETTING_DATA"
}

func New(b []byte) (ret *MSXLS, err error) {
	ret = new(MSXLS)
	ret.Src = b

	if ret.cf, err = compoundFile.NewCompoundFile(b); err != nil {
		return
	}
	if err = ret.cf.Parse(); err != nil {
		return
	}

	return
}

func (me *MSXLS) ParseWorkbook() (ret *Workbook, err error) {
	ret = new(Workbook)

	if ret.b, err = me.cf.GetStream(`Workbook`); err != nil {
		return
	}

	ret.mDecodeRecord = map[RecordType]func(record, *Workbook) (IRecord, error){
		RecordBOF:         decodeBOF,
		RecordWsBool:      decodeWsBool,
		RecordBoundSheet8: decodeBoundSheet8,
		RecordSST:         decodeSST,
		RecordDimensions:  decodeDimensions,
		//	RK record specifies the numeric data contained in a single cell.
		//	MulRk record specifies a series of cells with numeric data in a sheet row.
		//	Number record specifies a cell that contains a floating-point number.
		//	LabelSst record specifies a cell that contains a string.
		//	FormulaValue record specifies the current value of a formula
		RecordMulRk:    decodeMulRk,
		RecordFormula:  decodeFormula,
		RecordRK:       decodeRk,
		RecordNumber:   decodeNumber,
		RecordBoolErr:  decodeBoolErr,
		RecordLabelSst: decodeLabelSst,
	}

	for ret.pointer < len(ret.b) {
		var r record
		if r, err = decodeRecordHead(ret); err != nil {
			return
		}
		if f, ok := ret.mDecodeRecord[r.Type]; ok {
			var ir IRecord
			if ir, err = f(r, ret); err != nil {
				return
			}
			if err = ir.parse(ret); err != nil {
				return
			}

			//			switch ir.(type) {
			//			case *bof:
			//				print("bof\n")
			//			default:

			//			}
		} else {
			ret.pointer += int(r.Size)
		}
	}

	return
}

func decodeRecordHead(wk *Workbook) (r record, err error) {
	buf := wk.b[wk.pointer:]
	if len(wk.b[wk.pointer:]) < 4 {
		errors.New("不足 [record] Size，至少4字节.")
		return
	}

	r = record{
		Type: byte2uint16(buf),
		Size: byte2uint16(buf[2:]),
	}
	wk.pointer += 4

	return
}

func (me *RowCol) Address() (ret string) {
	var step int = 'Z' - 'A' + 1
	var tmp int = int(me.Column) + 1

	for tmp > step {
		tmptmp := tmp % step
		tmp /= step
		ret += string([]byte{byte(tmp + 'A' - 1)})
		tmp = tmptmp
	}
	if tmp > 0 {
		ret += string([]byte{byte(tmp + 'A' - 1)})
	}

	ret = ret + strconv.Itoa(int(me.Row+1))
	return
}

func (me *Worksheet) UsedRange() (ret [][]interface{}) {
	rows := me.last.Row - me.first.Row + 1
	cols := me.last.Column - me.first.Column + 1

	ret = make([][]interface{}, rows)
	for i := range ret {
		ret[i] = make([]interface{}, cols)
		for j := range ret[i] {
			if index, ok := me.MCell[RowCol{uint16(i), uint16(j)}]; ok {
				ret[i][j] = me.Cells[index].Value
			} else {
				ret[i][j] = ""
			}
		}
	}

	return
}

func (me *Worksheet) GetCell(rowZeroBase, colZeroBase uint16) (c Cell) {
	if index, ok := me.MCell[RowCol{rowZeroBase, colZeroBase}]; ok {
		c = me.Cells[index]
	} else {
		c = Cell{"", RowCol{rowZeroBase, colZeroBase}}
	}
	return
}

func (me *Worksheet) SortCell() {
	sort.Sort(me)

	me.MCell = make(map[RowCol]int)
	for i := range me.Cells {
		me.MCell[me.Cells[i].RowCol] = i
	}

	return
}
