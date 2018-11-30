package keyboard

import (
	"fmt"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

const (
	_           uint16 = iota
	_                  //1
	_                  //2
	_                  //3
	_                  //4
	_                  //5
	_                  //6
	_                  //7
	VK_BACK            //8 退格键
	VK_TAB             //9 TAB键
	VK_SHIFT           //10 Shift键
	_                  //11
	_                  //12
	VK_RETURN          //13 回车键
	_                  //14
	_                  //15
	_                  //16
	VK_CONTROL         //17 Ctrl键
	VK_MENU            //18 Alt键
	VK_PAUSE           //19 Pause Break键
	VK_CAPITAL         //20 Caps Lock键
	_                  //21
	_                  //22
	_                  //23
	_                  //24
	_                  //25
	_                  //26
	VK_ESCAPE          //27 ESC键
	_                  //28
	_                  //29
	_                  //30
	_                  //31
	VK_SPACE           //32 空格键
	VK_PRIOR           //33 Page Up
	VK_NEXT            //34 PageDown
	VK_END             //35 End键
	VK_HOME            //36 Home键
	VK_LEFT            //37 方向键(←)
	VK_UP              //38 方向键(↑)
	VK_RIGHT           //39 方向键(→)
	VK_DOWN            //40 方向键(↓)
	_                  //41
	_                  //42
	_                  //43
	_                  //44
	VK_INSERT          //45 Insert键
	VK_DELETE          //46 Delete键
	_                  //47
	VK_0               //48 0
	VK_1               //49 1
	VK_2               //50 2
	VK_3               //51 3
	VK_4               //52 4
	VK_5               //53 5
	VK_6               //54 6
	VK_7               //55 7
	VK_8               //56 8
	VK_9               //57 9
	_                  //58
	_                  //59
	_                  //60
	_                  //61
	_                  //62
	_                  //63
	_                  //64
	VK_A               //65 A
	VK_B               //66 B
	VK_C               //67 C
	VK_D               //68 D
	VK_E               //69 E
	VK_F               //70 F
	VK_G               //71 G
	VK_H               //72 H
	VK_I               //73 I
	VK_J               //74 J
	VK_K               //75 K
	VK_L               //76 L
	VK_M               //77 M
	VK_N               //78 N
	VK_O               //79 O
	VK_P               //80 P
	VK_Q               //81 Q
	VK_R               //82 R
	VK_S               //83 S
	VK_T               //84 T
	VK_U               //85 U
	VK_V               //86 V
	VK_W               //87 W
	VK_X               //88 X
	VK_Y               //89 Y
	VK_Z               //90 Z
	VK_LWIN            //91 左徽标键
	VK_RWIN            //92 右徽标键
	VK_APPS            //93 鼠标右键快捷键
	_                  //94
	_                  //95
	VK_NUMPAD0         //96 小键盘0
	VK_NUMPAD1         //97 小键盘1
	VK_NUMPAD2         //98 小键盘2
	VK_NUMPAD3         //99 小键盘3
	VK_NUMPAD4         //100 小键盘4
	VK_NUMPAD5         //101 小键盘5
	VK_NUMPAD6         //102 小键盘6
	VK_NUMPAD7         //103 小键盘7
	VK_NUMPAD8         //104 小键盘8
	VK_NUMPAD9         //105 小键盘9
	VK_MULTIPLY        //106 小键盘*
	VK_ADD             //107 小键盘+
	_                  //108
	VK_SUBTRACT        //109 小键盘-
	VK_DECIMAL         //110 小键盘.
	VK_DIVIDE          //111 小键盘/
	VK_F1              //112 F1键
	VK_F2              //113 F2键
	VK_F3              //114 F3键
	VK_F4              //115 F4键
	VK_F5              //116 F5键
	VK_F6              //117 F6键
	VK_F7              //118 F7键
	VK_F8              //119 F8键
	VK_F9              //120 F9键
	VK_F10             //121 F10键
	VK_F11             //122 F11键
	VK_F12             //123 F12键
	_                  //124
	_                  //125
	_                  //126
	_                  //127
	_                  //128
	_                  //129
	_                  //130
	_                  //131
	_                  //132
	_                  //133
	_                  //134
	_                  //135
	_                  //136
	_                  //137
	_                  //138
	_                  //139
	_                  //140
	_                  //141
	_                  //142
	_                  //143
	VK_NUMLOCK         //144 Num Lock键
	VK_SCROLL          //145 Scroll Lock键

)

func Press(wVk uint16) (ret uint32) {
	ip := user32.KEYBD_INPUT{user32.INPUT_KEYBOARD,
		user32.KEYBDINPUT{}}
	ip.Ki.WVk = wVk
	//	ip.Ki.WScan = wVk

	ipSize := int32(unsafe.Sizeof(ip))
	fmt.Println(ipSize)
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)
	ip.Ki.DwFlags = user32.KEYEVENTF_KEYUP
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)

	return
}

func Free() {
	user32.Free()
}
