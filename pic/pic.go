// 图片相关操作，Resize、Cos对比相似度等
package pic

import (
	"errors"
	"image"
	"math"
	"os"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/gdi32"
	"github.com/xiongyejun/xyjgo/winAPI/kernel32"
	"github.com/xiongyejun/xyjgo/winAPI/user32"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Pic struct {
	path string
	image.Image
	ext  string
	sqrt float64
}

// 新建1个Pic结构
func New(path string) (p *Pic, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	p = new(Pic)
	p.path = path
	if p.Image, p.ext, err = image.Decode(f); err != nil {
		return
	}

	return
}

// 计算2个图片的余弦cos
func (me *Pic) Cos(p *Pic) (ret float64, err error) {
	var numerator float64 // 分子
	if numerator, err = me.multiply(p); err != nil {
		return
	}

	me.Sqrt()
	p.Sqrt()
	var denominator float64 = me.sqrt * p.sqrt // 分母

	return numerator / denominator, nil
}

// 分子的乘积
// x1*y1 + x2*y2 ……
func (me *Pic) multiply(p *Pic) (ret float64, err error) {
	width := me.Bounds().Max.X
	height := me.Bounds().Max.Y

	if width != p.Bounds().Max.X || height != p.Bounds().Max.Y {
		return 0, errors.New("2个图片的长宽不一致。")
	}

	var sum uint32 = 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c1 := me.Image.At(i, j)
			r1, g1, b1, a1 := c1.RGBA()

			c2 := p.Image.At(i, j)
			r2, g2, b2, a2 := c2.RGBA()

			sum += (r1 >> 8) * (r2 >> 8)
			sum += (g1 >> 8) * (g2 >> 8)
			sum += (b1 >> 8) * (b2 >> 8)
			sum += (a1 >> 8) * (a2 >> 8)
			//			ret += (float64(uint8(r1)) * float64(uint8(r2)))
			//			ret += (float64(uint8(g1)) * float64(uint8(g2)))
			//			ret += (float64(uint8(b1)) * float64(uint8(b2)))
			//			ret += (float64(uint8(a1)) * float64(uint8(a2)))
		}
	}
	ret = float64(sum)
	return
}

// 这个一般是分母的
func (me *Pic) Sqrt() {
	width := me.Bounds().Max.X
	height := me.Bounds().Max.Y

	me.sqrt = 0
	var sum uint32 = 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := me.Image.At(i, j)
			r, g, b, a := c.RGBA()

			r >>= 8
			sum += r * r

			g >>= 8
			sum += g * g

			b >>= 8
			sum += b * b

			a >>= 8
			sum += a * a
			//			me.sqrt += (float64(uint8(r)) * float64(uint8(r)))
			//			me.sqrt += (float64(uint8(g)) * float64(uint8(g)))
			//			me.sqrt += (float64(uint8(b)) * float64(uint8(b)))
			//			me.sqrt += (float64(uint8(a)) * float64(uint8(a)))
		}
	}
	me.sqrt = math.Sqrt(float64(sum))
}

// 改变大小
func (me *Pic) Resize(width, height int) (err error) {
	return nil
}

// https://github.com/kbinani/screenshot
func Screen(hwnd uint32, width, height int) (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rect(0, 0, width, height))

	hdc := user32.GetDC(hwnd)
	defer user32.Free()

	if hdc == 0 {
		return nil, errors.New("GetDC failed")
	}
	defer user32.ReleaseDC(hwnd, hdc)

	memory_device := gdi32.CreateCompatibleDC(hdc)
	if memory_device == 0 {
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer gdi32.Free()
	defer gdi32.DeleteDC(memory_device)

	bitmap := gdi32.CreateCompatibleBitmap(hdc, int32(width), int32(height))
	if bitmap == 0 {
		return nil, errors.New("CreateCompatibleBitmap failed")
	}
	defer gdi32.DeleteDC(bitmap)

	var header gdi32.BITMAPINFOHEADER
	header.BiSize = uint32(unsafe.Sizeof(header))
	header.BiPlanes = 1
	header.BiBitCount = 32
	header.BiWidth = int32(width)
	header.BiHeight = int32(-height)
	header.BiCompression = gdi32.BI_RGB
	header.BiSizeImage = 0

	// GetDIBits balks at using Go memory on some systems. The MSDN example uses
	// GlobalAlloc, so we'll do that too. See:
	// https://docs.microsoft.com/en-gb/windows/desktop/gdi/capturing-an-image
	bitmapDataSize := uintptr(((int64(width)*int64(header.BiBitCount) + 31) / 32) * 4 * int64(height))
	hmem := kernel32.GlobalAlloc(kernel32.GMEM_MOVEABLE, bitmapDataSize)
	defer kernel32.Free()
	defer kernel32.GlobalFree(hmem)
	memptr := kernel32.GlobalLock(hmem)
	defer kernel32.GlobalUnlock(hmem)

	old := gdi32.SelectObject(memory_device, bitmap)
	if old == 0 {
		return nil, errors.New("SelectObject failed")
	}
	defer gdi32.SelectObject(memory_device, old)

	if !gdi32.BitBlt(memory_device, 0, 0, int32(width), int32(height), hdc, 0, 0, gdi32.SRCCOPY) {
		return nil, errors.New("BitBlt failed")
	}

	if gdi32.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*gdi32.BITMAPINFO)(unsafe.Pointer(&header)), gdi32.DIB_RGB_COLORS) == 0 {
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(memptr)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v0 := *(*uint8)(unsafe.Pointer(src))
			v1 := *(*uint8)(unsafe.Pointer(src + 1))
			v2 := *(*uint8)(unsafe.Pointer(src + 2))

			// BGRA => RGBA, and set A to 255
			img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = v2, v1, v0, 255

			i += 4
			src += 4
		}
	}

	return img, nil
}
