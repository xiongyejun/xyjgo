package fileHash

import (
	"testing"
)

func Test_func(t *testing.T) {
	d := New(`E:\00-学习资料\15-go\src\github.com\xiongyejun\xyjgo\fileHash\1`)

	if err := d.GetFilesInfo(); err != nil {
		t.Error(err)
	} else {
		for i := range d.FilesInfo {
			t.Logf("%d %s %s\n", i, d.FilesInfo[i].Name, d.FilesInfo[i].Hash)
		}
	}
}
