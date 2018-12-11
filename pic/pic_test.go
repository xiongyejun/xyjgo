package pic

import (
	"image/jpeg"
	"os"
	"testing"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

func Test_new(t *testing.T) {
	hwnd := user32.FindWindow("", "MapleStory")

	var rect user32.RECT
	user32.GetWindowRect(hwnd, &rect)

	img, err := Screen(hwnd, int(rect.Right-rect.Left), int(rect.Bottom-rect.Top))
	if err != nil {
		t.Log(err)
		return
	}

	f, err := os.Create("3.jpg")
	if err != nil {
		t.Log(err)
		return
	}
	defer f.Close()
	jpeg.Encode(f, img, nil)

	t.Log(img.At(100, 100))
}
