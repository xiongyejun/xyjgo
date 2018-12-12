package pic

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_new(t *testing.T) {
	if p, err := New("test.png"); err != nil {
		t.Error(err)
	} else {
		if img, _, err := NewImage("find.png"); err != nil {
			t.Error(err)
		} else {
			pic := Image2RGBA(img)
			if retx, rety, retSimilar, err := p.FindSimilar(pic, 0.85); err != nil {
				t.Error(err)
			} else {
				t.Log(retx, rety, retSimilar)

				minp := image.Point{retx, rety}
				maxp := image.Point{retx + p.Bounds().Max.X, rety + p.Bounds().Max.Y}
				subimg := pic.SubImage(image.Rectangle{minp, maxp})

				if f, err := os.Create("similar0.85.png"); err != nil {
					t.Error(err)
				} else {
					defer f.Close()
					if err := png.Encode(f, subimg); err != nil {
						t.Error(err)
					}
				}
			}
		}
	}
}
