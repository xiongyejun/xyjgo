package screen

import (
	"errors"
	"image"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/gdi32"
	"github.com/xiongyejun/xyjgo/winAPI/kernel32"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type ScreenData struct {
	hwnd, hdc, memory_device, bitmap, hmem, old, mehmem uint32
	width, height                                       int32
	header                                              gdi32.BITMAPINFOHEADER
	bitmapDataSize                                      uintptr
	memptr                                              unsafe.Pointer
}

func New(hwnd uint32, width, height int32) (s *ScreenData) {
	s = new(ScreenData)
	s.hwnd = hwnd
	s.width = width
	s.height = height

	s.initData()

	return
}

func (me *ScreenData) initData() {
	me.hdc = user32.GetDC(me.hwnd)
	if me.hdc == 0 {
		panic("GetDC failed")
	}

	me.memory_device = gdi32.CreateCompatibleDC(me.hdc)
	if me.memory_device == 0 {
		panic("CreateCompatibleDC failed")
	}

	me.bitmap = gdi32.CreateCompatibleBitmap(me.hdc, me.width, me.height)
	if me.bitmap == 0 {
		panic("CreateCompatibleBitmap failed")
	}

	me.header.BiSize = uint32(unsafe.Sizeof(me.header))
	me.header.BiPlanes = 1
	me.header.BiBitCount = 32
	me.header.BiWidth = me.width
	me.header.BiHeight = -me.height
	me.header.BiCompression = gdi32.BI_RGB
	me.header.BiSizeImage = 0

	me.bitmapDataSize = uintptr(((int64(me.width)*int64(me.header.BiBitCount) + 31) / 32) * 4 * int64(me.height))

	// GetDIBits balks at using Go memory on some systems. The MSDN example uses
	// GlobalAlloc, so we'll do that too. See:
	// https://docs.microsoft.com/en-gb/windows/desktop/gdi/capturing-an-image
	me.hmem = kernel32.GlobalAlloc(kernel32.GMEM_MOVEABLE, me.bitmapDataSize)
	me.memptr = kernel32.GlobalLock(me.hmem)

	me.old = gdi32.SelectObject(me.memory_device, me.bitmap)
	if me.old == 0 {
		panic("SelectObject failed")
	}
}

func (me *ScreenData) Free() {
	//	gdi32.SelectObject(me.memory_device, me.old)
	kernel32.GlobalUnlock(me.mehmem)
	kernel32.GlobalFree(me.hmem)
	gdi32.DeleteDC(me.bitmap)
	gdi32.DeleteDC(me.memory_device)
	user32.ReleaseDC(me.hwnd, me.hdc)

	user32.Free()
	gdi32.Free()
	kernel32.Free()

}

// https://github.com/kbinani/screenshot
func (me *ScreenData) Screen(nXSrc, nYSrc int32) (img *image.RGBA, err error) {
	img = image.NewRGBA(image.Rect(0, 0, int(me.width), int(me.height)))

	if !gdi32.BitBlt(me.memory_device, 0, 0, me.width, me.height, me.hdc, nXSrc, nYSrc, gdi32.SRCCOPY) {
		return nil, errors.New("BitBlt failed")
	}

	if gdi32.GetDIBits(me.hdc, me.bitmap, 0, uint32(me.height), (*uint8)(me.memptr), (*gdi32.BITMAPINFO)(unsafe.Pointer(&me.header)), gdi32.DIB_RGB_COLORS) == 0 {
		return nil, errors.New("GetDIBits failed")
	}

	i := 0
	src := uintptr(me.memptr)
	var x, y int32
	for y = 0; y < me.height; y++ {
		for x = 0; x < me.width; x++ {
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
