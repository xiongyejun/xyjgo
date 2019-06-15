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

type rgba struct {
	r, g, b, a uint64 // 用uint64避免溢出溢出
}
type Pic struct {
	path string
	image.Image
	ext   string
	sqrt  float64
	rgbas [][]rgba
}

// 新建1个Pic结构
func New(path string) (p *Pic, err error) {
	p = new(Pic)
	p.path = path
	if p.Image, p.ext, err = NewImage(path); err != nil {
		return
	}

	return
}

func NewImage(path string) (img image.Image, ext string, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	if img, ext, err = image.Decode(f); err != nil {
		return
	}

	return
}

func Image2RGBA(img image.Image) (ret *image.RGBA) {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	ret = image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := img.At(i, j)
			ret.Set(i, j, c)
		}
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
	var sum uint64 = 0
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			sum += me.rgbas[i][j].r * me.rgbas[i][j].r
			sum += me.rgbas[i][j].g * me.rgbas[i][j].g
			sum += me.rgbas[i][j].b * me.rgbas[i][j].b
			//			sum += me.rgbas[i][j].a * me.rgbas[i][j].a // 图片透明度一般不用
		}
	}
	me.sqrt = math.Sqrt(float64(sum))
}

// 改变大小
func (me *Pic) Resize(width, height int) (err error) {
	return nil
}

// 把RGBA获取到，是uint8类型的，转化为uint64用来计算
func (me *Pic) GetRGBA() {
	width := me.Bounds().Max.X
	height := me.Bounds().Max.Y

	me.rgbas = New2DSlice(width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := me.Image.At(i, j)
			r, g, b, a := c.RGBA()

			me.rgbas[i][j].a = uint64(a >> 8)
			me.rgbas[i][j].b = uint64(b >> 8)
			me.rgbas[i][j].g = uint64(g >> 8)
			me.rgbas[i][j].r = uint64(r >> 8)
		}
	}
}

func New2DSlice(x int, y int) (theSlice [][]rgba) {
	theSlice = make([][]rgba, x, x)
	for i := 0; i < x; i++ {
		s2 := make([]rgba, y)
		theSlice[i] = s2
	}
	return
}

// 在img里，查找最相似的，返回在img中的开始坐标
// similar	如果similar大于0，找到>=similar的就可以停止
func (me *Pic) FindSimilar(img *image.RGBA, similar float64) (retx, rety int, retSimilar float64, err error) {
	width1 := me.Bounds().Max.X
	height1 := me.Bounds().Max.Y

	width2 := img.Bounds().Max.X
	height2 := img.Bounds().Max.Y

	if width1 > width2 {
		err = errors.New("width1 > width2")
		return
	}

	if height1 > height2 {
		err = errors.New("height1 > height2")
		return
	}

	endx := width2 - width1
	endy := height2 - height1

	// 先计算me的Sqrt----放到调用的程序中用
	//	me.GetRGBA()
	//	me.Sqrt()

	tmp := new(Pic)
	tmp.Image = img
	tmp.GetRGBA()

	var cos float64 = 0
	// 逐个计算cos
	var i, j int
	for i = 0; i < endx; i++ { // 起点x
		for j = 0; j < endy; j++ { // 起点y

			var numerator uint64 = 0    // 分子
			var denominator2 uint64 = 0 // tmp的分母

			for x := 0; x < width1; x++ {
				for y := 0; y < height1; y++ {
					// 两两相乘
					numerator += me.rgbas[x][y].r * tmp.rgbas[x+i][y+j].r
					numerator += me.rgbas[x][y].g * tmp.rgbas[x+i][y+j].g
					numerator += me.rgbas[x][y].b * tmp.rgbas[x+i][y+j].b

					denominator2 += tmp.rgbas[x+i][y+j].r * tmp.rgbas[x+i][y+j].r
					denominator2 += tmp.rgbas[x+i][y+j].g * tmp.rgbas[x+i][y+j].g
					denominator2 += tmp.rgbas[x+i][y+j].b * tmp.rgbas[x+i][y+j].b
				}
			}
			// 计算cos，并获取大的那个
			cos = float64(numerator) / (math.Sqrt(float64(denominator2)) * me.sqrt)
			if cos > retSimilar {
				retSimilar = cos
				retx = i
				rety = j
			}

			if similar > 0 {
				if retSimilar > similar {
					return
				}
			}
		}
	}

	return
}

// https://github.com/kbinani/screenshot
func Screen(hwnd uint32, nXSrc, nYSrc, width, height int32) (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	var errCount int = 0

restart:
	hdc := user32.GetDC(hwnd)
	defer user32.Free()

	if hdc == 0 {
		errCount++
		if errCount < 3 {
			goto restart
		}
		return nil, errors.New("GetDC failed")
	}
	defer user32.ReleaseDC(hwnd, hdc)

	memory_device := gdi32.CreateCompatibleDC(hdc)
	if memory_device == 0 {
		errCount++
		if errCount < 3 {
			goto restart
		}
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer gdi32.Free()
	defer gdi32.DeleteDC(memory_device)

	bitmap := gdi32.CreateCompatibleBitmap(hdc, int32(width), int32(height))
	if bitmap == 0 {
		errCount++
		if errCount < 3 {
			goto restart
		}
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
		errCount++
		if errCount < 3 {
			goto restart
		}
		return nil, errors.New("SelectObject failed")
	}
	defer gdi32.SelectObject(memory_device, old)

	if !gdi32.BitBlt(memory_device, 0, 0, int32(width), int32(height), hdc, nXSrc, nYSrc, gdi32.SRCCOPY) {
		errCount++
		if errCount < 3 {
			goto restart
		}
		return nil, errors.New("BitBlt failed")
	}

	if gdi32.GetDIBits(hdc, bitmap, 0, uint32(height), (*uint8)(memptr), (*gdi32.BITMAPINFO)(unsafe.Pointer(&header)), gdi32.DIB_RGB_COLORS) == 0 {
		errCount++
		if errCount < 3 {
			goto restart
		}
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(memptr)
	var x, y int32
	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
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
