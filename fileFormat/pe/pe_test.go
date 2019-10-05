package pe

import (
	"os"
	"testing"
)

func Test_t(t *testing.T) {
	var err error
	var f *os.File
	if f, err = os.Open(`..\testc\SimpleSection.exe`); err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	var pe *PE = new(PE)
	if err = pe.Parse(f); err != nil {
		t.Log(err)
	} else {
		t.Logf("%x\n", pe.DosHeader.E_lfanew)

		t.Logf("%x\n", pe.NTHeader.Signature)

		for i := range pe.NTHeader.OptionalHeader.DataDirectory {
			t.Logf("%2d %x\t%x\n", i, pe.NTHeader.OptionalHeader.DataDirectory[i].VirtualAddress, pe.NTHeader.OptionalHeader.DataDirectory[i].Size)
		}
	}
}
