package o

import (
	"testing"
	"time"
)

func Test_f(t *testing.T) {
	var coff *COFF = new(COFF)

	if err := coff.Parse(`c\SimpleSection.obj`); err != nil {
		t.Error(err)
		return
	} else {
		t.Logf("FILE HEADER VALUES\n")
		t.Logf("%16X machine (x86)\n", coff.Header.Machine)
		t.Logf("%16X number of sections\n", coff.Header.NumberOfSections)
		t.Logf("%16X  time date stamp %s\n", coff.Header.TimeDateStamp, time.Unix(int64(coff.Header.TimeDateStamp), 0).Format("2006-01-02 15:04:05"))
		t.Logf("%16X file pointer to symbol table\n", coff.Header.PointerToSymbolTable)
		t.Logf("%16X number of symbols\n", coff.Header.NumberOfSymbols)
		t.Logf("%16X size of optional header\n", coff.Header.SizeOfOptionalHeader)
		t.Logf("%16X characteristics\n", coff.Header.Characteristics)

		for i := range coff.Sections {
			t.Logf("%d %s\t%x\t%x\n", i, coff.Sections[i].Name, coff.Sections[i].Misc, coff.Sections[i].SizeOfRawData)
		}

		for i := range coff.Symbols {
			t.Logf("%d %s\t%x\t%x\n", i, coff.Symbols[i].ShortName, coff.Symbols[i].Value, coff.Symbols[i].SectionNumber)
		}
	}
}
