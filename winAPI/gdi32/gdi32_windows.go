package gdi32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                    uintptr
	getPixel               uintptr
	createCompatibleDC     uintptr
	deleteDC               uintptr
	createCompatibleBitmap uintptr
	deleteObject           uintptr
	selectObject           uintptr
	bitBlt                 uintptr
	getDIBits              uintptr
)

// Bitmap compression constants
const (
	BI_RGB       = 0
	BI_RLE8      = 1
	BI_RLE4      = 2
	BI_BITFIELDS = 3
	BI_JPEG      = 4
	BI_PNG       = 5
)

// Bitmap color table usage
const (
	DIB_RGB_COLORS = 0
	DIB_PAL_COLORS = 1
)

// Ternary raster operations
const (
	SRCCOPY        = 0x00CC0020
	SRCPAINT       = 0x00EE0086
	SRCAND         = 0x008800C6
	SRCINVERT      = 0x00660046
	SRCERASE       = 0x00440328
	NOTSRCCOPY     = 0x00330008
	NOTSRCERASE    = 0x001100A6
	MERGECOPY      = 0x00C000CA
	MERGEPAINT     = 0x00BB0226
	PATCOPY        = 0x00F00021
	PATPAINT       = 0x00FB0A09
	PATINVERT      = 0x005A0049
	DSTINVERT      = 0x00550009
	BLACKNESS      = 0x00000042
	WHITENESS      = 0x00FF0062
	NOMIRRORBITMAP = 0x80000000
	CAPTUREBLT     = 0x40000000
)

type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

func init() {
	// Library
	lib = win.MustLoadLibrary("gdi32.dll")

	// Functions
	getPixel = win.MustGetProcAddress(lib, "GetPixel")
	createCompatibleDC = win.MustGetProcAddress(lib, "CreateCompatibleDC")
	deleteDC = win.MustGetProcAddress(lib, "DeleteDC")
	createCompatibleBitmap = win.MustGetProcAddress(lib, "CreateCompatibleBitmap")
	deleteObject = win.MustGetProcAddress(lib, "DeleteObject")
	selectObject = win.MustGetProcAddress(lib, "SelectObject")
	bitBlt = win.MustGetProcAddress(lib, "BitBlt")
	getDIBits = win.MustGetProcAddress(lib, "GetDIBits")
}

// COLORREF GetPixel(HDC hdc, int nXPos, int nYPos)
func GetPixel(HDC uint32, x, y int32) int32 {
	ret, _, _ := syscall.Syscall(getPixel, 3,
		uintptr(HDC),
		uintptr(x),
		uintptr(y))

	return int32(ret)
}

func CreateCompatibleDC(hdc uint32) uint32 {
	ret, _, _ := syscall.Syscall(createCompatibleDC, 1,
		uintptr(hdc),
		0,
		0)

	return uint32(ret)
}

func DeleteDC(hdc uint32) bool {
	ret, _, _ := syscall.Syscall(deleteDC, 1,
		uintptr(hdc),
		0,
		0)

	return ret != 0
}

func CreateCompatibleBitmap(hdc uint32, nWidth, nHeight int32) uint32 {
	ret, _, _ := syscall.Syscall(createCompatibleBitmap, 3,
		uintptr(hdc),
		uintptr(nWidth),
		uintptr(nHeight))

	return uint32(ret)
}

func DeleteObject(hObject uint32) bool {
	ret, _, _ := syscall.Syscall(deleteObject, 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}

func SelectObject(hdc, hgdiobj uint32) uint32 {
	ret, _, _ := syscall.Syscall(selectObject, 2,
		uintptr(hdc),
		uintptr(hgdiobj),
		0)

	return uint32(ret)
}

func BitBlt(hdcDest uint32, nXDest, nYDest, nWidth, nHeight int32, hdcSrc uint32, nXSrc, nYSrc int32, dwRop uint32) bool {
	ret, _, _ := syscall.Syscall9(bitBlt, 9,
		uintptr(hdcDest),
		uintptr(nXDest),
		uintptr(nYDest),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hdcSrc),
		uintptr(nXSrc),
		uintptr(nYSrc),
		uintptr(dwRop))

	return ret != 0
}

func GetDIBits(hdc, hbmp, uStartScan, cScanLines uint32, lpvBits *byte, lpbi *BITMAPINFO, uUsage uint32) int32 {
	ret, _, _ := syscall.Syscall9(getDIBits, 7,
		uintptr(hdc),
		uintptr(hbmp),
		uintptr(uStartScan),
		uintptr(cScanLines),
		uintptr(unsafe.Pointer(lpvBits)),
		uintptr(unsafe.Pointer(lpbi)),
		uintptr(uUsage),
		0,
		0)
	return int32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
