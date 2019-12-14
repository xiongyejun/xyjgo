package pdf

import (
	"os"
	"testing"

	"github.com/axgle/mahonia"
)

func Test_func(t *testing.T) {
	t.Logf("%s\r\n", []byte{0x25, 0x50, 0x44, 0x46})
	f, err := os.Open("1.pdf")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	if p, err := Parse(f); err != nil {
		t.Error(err)
	} else {

		for i := range p.pagesInfo {
			t.Logf("page %d ContentsObjIndex=%d, ResourcesObjIndex=%d, ResourcesType = %d, ", i, p.pagesInfo[i].ContentsObjIndex, p.pagesInfo[i].ResourcesObjIndex, p.pagesInfo[i].ResourcesType)

			if b, err := p.GetPageByte(i); err != nil {
				t.Error(err)
			} else {
				t.Logf("%s\n", b)
				t.Log("\n")

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
	bbbb := []byte(utf8To("测试pdf", "utf16"))

	t.Logf("% x", bbbb)
	for i := range bbbb {
		bbbb[i] = ^bbbb[i]
	}
	t.Logf("% x", bbbb)

}

func gbkToUtf8(b []byte) string {
	decoder := mahonia.NewDecoder("gbk")
	return string(decoder.ConvertString(string(b)))
}
func utf8ToGbk(src string) string {
	srcCoder := mahonia.NewEncoder("gbk")
	return srcCoder.ConvertString(src)
}
func utf8To(src string, code string) string {
	srcCoder := mahonia.NewEncoder(code)
	return srcCoder.ConvertString(src)
}
