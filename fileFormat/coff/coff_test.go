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
		t.Log(coff.Header.GetCoffHeaderPrintStr())

		t.Log("\nSections:")
		for i := range coff.Sections {
			t.Logf("%d %s\t%x\t%x\n", i, coff.Sections[i].Name, coff.Sections[i].Misc, coff.Sections[i].SizeOfRawData)
		}
		t.Log("\nSymbols:")
		for i := range coff.Symbols {
			t.Logf("%d %s\t%x\t%x\n", i, coff.Symbols[i].ShortName, coff.Symbols[i].Value, coff.Symbols[i].SectionNumber)
		}
	}
}
