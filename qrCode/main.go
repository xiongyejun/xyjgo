package main

import (
	"fmt"
	"os"

	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

func main() {
	if len(os.Args) == 1 {
		uiInit()
	} else if len(os.Args) == 3 {
		if str, err := decode(os.Args[2]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`
 qrcode d <picpath> -- decode qrcode	
	`)
}

type uiControl struct {
	form    *walk.MainWindow
	tb      *walk.TextEdit
	iv      *walk.ImageView
	lbTiShi *walk.Label

	images []walk.Image
	pImage int
}

var ct *uiControl = new(uiControl)

func uiInit() {
	mw := declarative.MainWindow{
		AssignTo: &ct.form,
		Title:    "qrCode",
		Size:     declarative.Size{800, 500},

		Layout: declarative.VBox{Margins: declarative.Margins{1, 1, 1, 1}},
		Children: []declarative.Widget{
			declarative.TextEdit{
				AssignTo: &ct.tb,
				Font:     declarative.Font{PointSize: 10},
				HScroll:  true,
				VScroll:  true,

				OnTextChanged: func() {
					if ct.tb.Text() != "" {
						ct.images = getImage(ct.tb.Text())
						ct.pImage = 0
						showImage()
					}
				},
			},

			//			declarative.Separator{},

			declarative.Label{
				AssignTo: &ct.lbTiShi,
				Text:     "提示",
			},

			declarative.ImageView{
				AssignTo: &ct.iv,

				OnMouseUp: func(x, y int, button walk.MouseButton) {
					showImage()
				},
			},
		},
	}
	if err := mw.Create(); err != nil {
		walk.MsgBox(nil, "", err.Error(), walk.MsgBoxIconInformation)
		fmt.Println(err)
		return
	}

	setMiddlePos(ct.form.Handle(), nil)
	ct.form.Run()
}

func showImage() {
	if len(ct.images) == 0 {
		return
	}
	ct.lbTiShi.SetText(fmt.Sprintf("共%d张 当前第%d张", len(ct.images), ct.pImage+1))
	ct.iv.SetImage(ct.images[ct.pImage])
	ct.pImage++
	ct.pImage %= len(ct.images)
}

// 设置窗口居中
func setMiddlePos(hwnd win.HWND, owner walk.Form) {
	var srcWidth, srcHeight int32

	if owner == nil {
		srcWidth = win.GetSystemMetrics(win.SM_CXSCREEN)
		srcHeight = win.GetSystemMetrics(win.SM_CYSCREEN)
	} else {
		srcWidth = int32(owner.Width()) + 2*int32(owner.X())
		srcHeight = int32(owner.Height()) + 2*int32(owner.Y())
	}

	rect := new(win.RECT)
	win.GetWindowRect(hwnd, rect)
	win.SetWindowPos(hwnd, win.HWND_TOPMOST,
		(srcWidth-(rect.Right-rect.Left))/2,
		(srcHeight-(rect.Bottom-rect.Top))/2,
		rect.Right-rect.Left,
		rect.Bottom-rect.Top,
		win.SWP_SHOWWINDOW)
}
