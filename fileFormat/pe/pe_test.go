package pe

import (
	"os"
	"testing"
)

func Test_t(t *testing.T) {
	var err error
	var f *os.File
	/*  ..\testc\Math.dll   C:\Windows\System32\rnr20.dll  */
	if f, err = os.Open(`C:\Windows\System32\rnr20.dll`); err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	var pe *PE = new(PE)
	if err = pe.Parse(f); err != nil {
		t.Log(err)
	} else {
		t.Log(pe.NTHeader.FileHeader.String())

		t.Log(pe.NTHeader.OptionalHeader.String())

		for i := range pe.Sections {
			t.Log(pe.Sections[i].String(i))
		}

		t.Log(pe.ExportDir.String())
		t.Log(pe.ExportDirInfo.String())

	}
}
