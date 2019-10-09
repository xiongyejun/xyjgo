package pe

import (
	"os"
	"testing"
)

func Test_t(t *testing.T) {
	var err error
	var f *os.File
	if f, err = os.Open(`..\testc\Math.dll`); err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	var pe *PE = new(PE)
	if err = pe.Parse(f); err != nil {
		t.Log(err)
	} else {
		t.Log(pe.NTHeader.FileHeader.GetPrintStr())

		t.Log(pe.NTHeader.OptionalHeader.GetPrintStr())

		for i := range pe.Sections {
			t.Log(pe.Sections[i].GetPrintStr(i))
		}

		t.Logf("%x\n", pe.NTHeader.OptionalHeader.DataDirectory[0].VirtualAddress)
		t.Log("pe.ExportDir:")
		t.Logf("%#v\n", pe.ExportDir)
	}
}
