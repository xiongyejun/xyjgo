package pressKey

import (
	"errors"
	"time"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
	"github.com/xiongyejun/xyjgo/winAPI/user32/keyboard"
)

type key struct {
	WVk   int
	Time  time.Duration // 间隔，毫秒
	Delay time.Duration // 按键后延迟时间，毫秒
}

type Key struct {
	WindowText string
	Keys       []*key
	ch         chan *key
	chLen      int
	bStop      bool
}

func New() *Key {
	K := new(Key)
	K.chLen = 10
	return K
}

func newkey(WVk int, Time time.Duration, Delay time.Duration) (k *key) {
	return &key{WVk, Time, Delay}
}

// 增加1个按键
func (me *Key) Add(WVk int, Time time.Duration, Delay time.Duration) (err error) {
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
	me.ch = make(chan *key, me.chLen)

	for i := range me.Keys {
		me.ch <- me.Keys[i]
	}
	// 不停的循环读取channel里的数据
	for {
		for k := range me.ch {
			k.press(me.ch, me.WindowText == user32.GetWindowText(user32.GetForegroundWindow()))
			if me.bStop {
				close(me.ch)
				return
			}
		}

		time.Sleep(time.Second)
	}

}

// 是否是需要的活动窗口
func (me *key) press(ch chan *key, bWindowText bool) {
	if bWindowText {
		keyboard.Press(me.WVk)
		if me.Delay > 0 {
			// 在一定时间内阻止键盘输入
			time.Sleep(me.Delay)
		}
	}

	// 按完之后，等到了时间Time，继续插入到channel
	go func() {
		time.Sleep(me.Time / 1000 * time.Second)
		ch <- me
	}()

}

//func (me *Key) Pause() {

//}

//func (me *Key) Continue() {

//}

func (me *Key) Stop() {
	me.bStop = true
}
