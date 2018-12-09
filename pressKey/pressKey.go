package pressKey

import (
	"errors"
	"time"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
	"github.com/xiongyejun/xyjgo/winAPI/user32/keyboard"
)

type key struct {
	WVk   uint16
	Time  time.Duration // 间隔，毫秒
	Delay time.Duration // 按键后延迟时间，毫秒
}

type Key struct {
	WindowText string
	Keys       []*key
	ch         chan *key
	chLen      int
}

var bStop bool

func New() *Key {
	K := new(Key)
	K.chLen = 100
	return K
}

func newkey(WVk uint16, Time time.Duration, Delay time.Duration) (k *key) {
	return &key{WVk, Time, Delay}
}

// 换1个键
func (me *Key) Change(index int, WVk uint16) (err error) {
	if index >= len(me.Keys) {
		return errors.New("out of range")
	}

	me.Keys[index].WVk = WVk
	return nil
}

// 增加1个按键
func (me *Key) Add(WVk uint16, Time time.Duration, Delay time.Duration) (err error) {
	me.Keys = append(me.Keys, newkey(WVk, Time, Delay))
	return nil
}

// 按下标来删除1个按键
func (me *Key) Remove(index int) (err error) {
	if index >= len(me.Keys) {
		return errors.New("out of range")
	}

	me.Keys = append(me.Keys[0:index], me.Keys[index+1:]...)
	return nil
}

func (me *Key) Start() {
	defer keyboard.Free()
	hwnd := user32.FindWindow("", me.WindowText)

	me.ch = make(chan *key, me.chLen)
	bStop = false

	for i := range me.Keys {
		me.ch <- me.Keys[i]
	}
	// 不停的循环读取channel里的数据
	for {
		for k := range me.ch {
			k.press(me.ch, hwnd == user32.GetForegroundWindow())
			if bStop {
				close(me.ch)
				return
			}
		}

		time.Sleep(time.Second)
	}

}

// bWindow 是否是需要的活动窗口
func (me *key) press(ch chan *key, bWindow bool) {
	if bWindow {
		keyboard.Press(me.WVk, me.Delay/1000*time.Second)
	} else {
		//		print("窗口不对啊")
	}

	// 按完之后，等到了时间Time，继续插入到channel
	go func() {
		time.Sleep(me.Time / 1000 * time.Second)
		if !bStop {
			ch <- me
		}
	}()
}

// 一定要保证最后成功执行了BlockInput(0)
func blockInput() {
	print("in blockInput")
	var ret uint32 = 0
	for ret == 0 {
		ret = user32.BlockInput(0)
	}
	print("out blockInput")
}

//func (me *Key) Pause() {

//}

//func (me *Key) Continue() {

//}

func (me *Key) Stop() {
	bStop = true
}
