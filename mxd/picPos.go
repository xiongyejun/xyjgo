package main

import (
	"errors"
	//	"fmt"
	"image"
	//	"time"

	"github.com/xiongyejun/xyjgo/winAPI/user32/keyboard"

	"github.com/xiongyejun/xyjgo/pic"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

// 根据图片 pos.png 来定位
// 首先找到pos.png的开始坐标，然后不停的截图
// 截图范围：3个pos.png的高度
// 			左右分别在1/4个屏幕宽
// 如果出现了图片pos.png，就切换方向
// 加1个计时器，如果超过了30秒，也强制换方向

type picPos struct {
	p                                                       *pic.Pic
	picname                                                 string // 固定是pos.png
	leftX, leftY, rightX, rightY, screenHeight, screenWidth int32
}

// 检查是否走到了边缘地带（1/4屏幕距离）
func (me *skey) check() (b bool, err error) {
	if me.moveValue == keyboard.VK_LEFT {
		return me.checkImg(me.leftX, me.leftY, keyboard.VK_RIGHT)
	} else {
		return me.checkImg(me.rightX, me.rightY, keyboard.VK_LEFT)
	}
}

func (me *skey) checkImg(x, y int32, vk uint16) (b bool, err error) {
	var img *image.RGBA
	//	fmt.Println("开始截图", time.Now())
	if img, err = pic.Screen(me.hwnd, x, y, me.screenWidth, me.screenHeight); err != nil {
		return false, err
	} else {
		//		fmt.Println("完成截图", time.Now())
		var retSimilar float64
		//		fmt.Println()
		//		fmt.Println("开始分析", time.Now())
		if _, _, retSimilar, err = me.p.FindSimilar(img, 0.95); err != nil {
			return false, err
		} else {
			//			fmt.Println("完成分析", time.Now())
			if retSimilar >= 0.95 {
				me.moveValue = vk
				return true, nil
			} else {
				//				fmt.Println("retSimilar=", retSimilar)
			}
		}
	}

	return
}

func (me *skey) getPos() (err error) {
	var rect *user32.RECT = new(user32.RECT)
	me.hwnd = user32.FindWindow("", "MapleStory")
	user32.GetClientRect(me.hwnd, rect)

	// 一般都在一半以下的高度（勋章）
	var offsetY int32 = (rect.Bottom - rect.Top) / 2

	var img *image.RGBA
	if img, err = pic.Screen(me.hwnd, 0, offsetY, rect.Right-rect.Left, rect.Bottom-rect.Top-offsetY); err != nil {
		return
	}
	if me.p, err = pic.New(me.path + me.picname); err != nil {
		return
	} else {
		me.p.GetRGBA()
		me.p.Sqrt()

		var retSimilar float64
		var retY int
		if _, retY, retSimilar, err = me.p.FindSimilar(img, 0.95); err != nil {
			return
		} else {
			if retSimilar < 0.95 {
				return errors.New("retSimilar < 0.95")
			}

			me.screenHeight = int32(me.p.Bounds().Dy() + me.p.Bounds().Dx()/2) // 多1/2个高度
			me.screenWidth = (rect.Right - rect.Left) / 4
			me.leftX = 0
			me.leftY = int32(retY-me.p.Bounds().Dx()/4) + offsetY
			me.rightY = me.leftY
			me.rightX = (rect.Right - rect.Left) * 3 / 4

		}
	}

	return
}
