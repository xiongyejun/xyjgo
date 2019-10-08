package coff

import (
	"os"
	"testing"
)

func Test_f(t *testing.T) {
	var err error
	var f *os.File
	if f, err = os.Open(`..\testc\SimpleSection.obj`); err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	var coff *COFF = new(COFF)
	if err := coff.Parse(f); err != nil {
		t.Error(err)
		return
	} else {
		t.Log(coff.Header.GetPrintStr())

		t.Log("\nSections:")
		for i := range coff.Sections {
			t.Log(coff.Sections[i].GetPrintStr(i))
		}
		t.Log("\nSymbols:")
		for i := range coff.Symbols {
			t.Logf("%d %s\t%x\t%x\n", i, coff.Symbols[i].ShortName, coff.Symbols[i].Value, coff.Symbols[i].SectionNumber)
		}
	}
}
