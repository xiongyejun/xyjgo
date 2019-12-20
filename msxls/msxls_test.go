package msxls

import (
	"io/ioutil"
	"testing"
)

func Test_func(t *testing.T) {
	var b []byte
	var err error
	if b, err = ioutil.ReadFile(`test.xls`); err != nil {
		t.Error(err)
		return
	}
	var ms *MSXLS
	if ms, err = New(b); err != nil {
		t.Error(err)
		return
	}
	var wk *Workbook
	if wk, err = ms.ParseWorkbook(); err != nil {
		t.Error(err)
		return
	}
	for i := range wk.Worksheets {
		t.Logf("Worksheet %d, %s:", i, wk.Worksheets[i].Name)
		for j := range wk.Worksheets[i].Cells {
			t.Logf("%d %s=%v\n", j, wk.Worksheets[i].Cells[j].Address(), wk.Worksheets[i].Cells[j].Value)
		}
	}
}
