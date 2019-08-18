package pdf

import (
	"testing"

	"github.com/axgle/mahonia"
)

func Test_func(t *testing.T) {
	t.Logf("%s\r\n", []byte{0x25, 0x50, 0x44, 0x46})

	if p, err := Parse("1.pdf"); err != nil {
		t.Error(err)
	} else {
		defer p.f.Close()
		for i := range p.pagesInfo {
			t.Logf("page %d ContentsObjIndex=%d, ResourcesObjIndex=%d, ResourcesType = %d, ", i, p.pagesInfo[i].ContentsObjIndex, p.pagesInfo[i].ResourcesObjIndex, p.pagesInfo[i].ResourcesType)

			if b, err := p.GetPageByte(i); err != nil {
				t.Error(err)
			} else {
				if bb, err := p.ParsePageByte(b); err != nil {
					t.Error(err)
				} else {
					t.Logf("% x\n", bb)
					t.Log("\n")
					t.Logf("%s\n", bb)
				}
			}
			t.Log("\n")
		}

	}

	t.Logf("% x", []byte("测试pdf"))
	t.Logf("% x", []byte(utf8ToGbk("测试pdf")))
}

func gbkToUtf8(b []byte) string {
	decoder := mahonia.NewDecoder("gbk")
	return string(decoder.ConvertString(string(b)))
}
func utf8ToGbk(src string) string {
	srcCoder := mahonia.NewEncoder("gbk")
	return srcCoder.ConvertString(src)
}
