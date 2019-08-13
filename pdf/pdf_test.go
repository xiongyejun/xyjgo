package pdf

import (
	"testing"
)

func Test_func(t *testing.T) {
	t.Logf("%s\r\n", []byte{0x25, 0x50, 0x44, 0x46})

	if p, err := New(`2.pdf`); err != nil {
		t.Error(err)
	} else {
		if err := p.Parse(); err != nil {
			t.Error(err)
		} else {
			// if b, err := p.GetObjByte("1"); err != nil {
			// 	t.Error(err)
			// } else {
			// 	if r, err := findR(b, "Pages"); err != nil {
			// 		t.Error(err)
			// 	} else {
			// 		t.Log("r=", r)
			// 	}
			// }

			t.Logf("Src Len = %d, Root obj = %s \r\n", len(p.Src), p.RootR)
			for i := range p.objs {
				t.Logf("index%d\tobj index %s\tindex src = %d\r\n\r\n", i, p.objs[i].strIndex, p.objs[i].indexSrc)
			}
		}
	}
}
