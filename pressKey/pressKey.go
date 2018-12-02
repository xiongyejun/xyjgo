package pressKey

import (
	"errors"
	"time"

	"github.com/xiongyejun/xyjgo/winAPI/user32/keyboard"
)

type key struct {
	WVk   uint16
	Time  time.Duration // 间隔，毫秒
	Delay uint          // 按键后延迟时间，毫秒
}
type Key struct {
	Keys  []*key
	ch    chan *key
	chLen int
	bStop bool
}

func New() *Key {
	K := new(Key)
	K.chLen = 10
	K.ch = make(chan *key, K.chLen)
	return K
}

func newkey(WVk uint16, Time time.Duration, Delay uint) (k *key) {
	return &key{WVk, Time, Delay}
}

// 增加1个按键
func (me *Key) Add(WVk uint16, Time time.Duration, Delay uint) (err error) {
	me.Keys = append(me.Keys, newkey(WVk, Time, Delay))
	return nil
}

// 按下标来删除1个按键
func (me *Key) Remove(index int) (err error) {
	if index >= len(me.Keys) {
		return errors.New("out of range")
	}

	me.Keys = append(me.Keys[0:index], me.Keys[index:]...)
	return nil
}

func (me *Key) Start() {
	for i := range me.Keys {
		me.ch <- me.Keys[i]
	}
	// 不停的循环读取channel里的数据
	for {
		for k := range me.ch {
			k.press(me.ch)
		}
		if me.bStop {
			close(me.ch)
			return
		}
		time.Sleep(time.Second)
	}

}

func (me *key) press(ch chan *key) {
	keyboard.Press(me.WVk)
	if me.Delay > 0 {
		// 在一定时间内阻止键盘输入
	}
	// 按完之后，等到了时间Time，继续插入到channel
	time.Sleep(me.Time / 1000 * time.Second)
	ch <- me
}

//func (me *Key) Pause() {

//}

//func (me *Key) Continue() {

//}

func (me *Key) Stop() {
	me.bStop = true
}
