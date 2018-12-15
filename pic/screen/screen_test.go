package screen

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_screen(t *testing.T) {
	s := New(0, 100, 100)

	defer s.Free()

	img, err := s.Screen(100, 100)
	if err != nil {
		t.Error(err)
	} else {
		if err := savePic(img, "1.png"); err != nil {
			t.Error(err)
		}
	}

	img2, err := s.Screen(200, 200)
	if err != nil {
		t.Error(err)
	} else {
		if err := savePic(img2, "2.png"); err != nil {
			t.Error(err)
		}
	}

}

func savePic(img *image.RGBA, name string) error {
	if f, err := os.Create(name); err != nil {
		return err
	} else {
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			return err
		}
	}
	return nil
}
