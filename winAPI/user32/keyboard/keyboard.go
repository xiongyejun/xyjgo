package keyboard

import (
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type keyTable struct {
	vk   byte
	scan byte
}

const (
	_           uint16 = iota //1
	_                         //2
	_                         //3
	_                         //4
	_                         //5
	_                         //6
	_                         //7
	VK_BACK                   //8 退格键
	VK_TAB                    //9 TAB键
	VK_SHIFT                  //10 Shift键
	_                         //11
	_                         //12
	VK_RETURN                 //13 回车键
	_                         //14
	_                         //15
	_                         //16
	VK_CONTROL                //17 Ctrl键
	VK_MENU                   //18 Alt键
	VK_PAUSE                  //19 Pause Break键
	VK_CAPITAL                //20 Caps Lock键
	_                         //21
	_                         //22
	_                         //23
	_                         //24
	_                         //25
	_                         //26
	VK_ESCAPE                 //27 ESC键
	_                         //28
	_                         //29
	_                         //30
	_                         //31
	VK_SPACE                  //32 空格键
	VK_PRIOR                  //33 Page Up
	VK_NEXT                   //34 PageDown
	VK_END                    //35 End键
	VK_HOME                   //36 Home键
	VK_LEFT                   //37 方向键(←)
	VK_UP                     //38 方向键(↑)
	VK_RIGHT                  //39 方向键(→)
	VK_DOWN                   //40 方向键(↓)
	_                         //41
	_                         //42
	_                         //43
	_                         //44
	VK_INSERT                 //45 Insert键
	VK_DELETE                 //46 Delete键
	_                         //47
	VK_0                      //48 0
	VK_1                      //49 1
	VK_2                      //50 2
	VK_3                      //51 3
	VK_4                      //52 4
	VK_5                      //53 5
	VK_6                      //54 6
	VK_7                      //55 7
	VK_8                      //56 8
	VK_9                      //57 9
	_                         //58
	_                         //59
	_                         //60
	_                         //61
	_                         //62
	_                         //63
	_                         //64
	VK_A                      //65 A
	VK_B                      //66 B
	VK_C                      //67 C
	VK_D                      //68 D
	VK_E                      //69 E
	VK_F                      //70 F
	VK_G                      //71 G
	VK_H                      //72 H
	VK_I                      //73 I
	VK_J                      //74 J
	VK_K                      //75 K
	VK_L                      //76 L
	VK_M                      //77 M
	VK_N                      //78 N
	VK_O                      //79 O
	VK_P                      //80 P
	VK_Q                      //81 Q
	VK_R                      //82 R
	VK_S                      //83 S
	VK_T                      //84 T
	VK_U                      //85 U
	VK_V                      //86 V
	VK_W                      //87 W
	VK_X                      //88 X
	VK_Y                      //89 Y
	VK_Z                      //90 Z
	VK_LWIN                   //91 左徽标键
	VK_RWIN                   //92 右徽标键
	VK_APPS                   //93 鼠标右键快捷键
	_                         //94
	_                         //95
	VK_NUMPAD0                //96 小键盘0
	VK_NUMPAD1                //97 小键盘1
	VK_NUMPAD2                //98 小键盘2
	VK_NUMPAD3                //99 小键盘3
	VK_NUMPAD4                //100 小键盘4
	VK_NUMPAD5                //101 小键盘5
	VK_NUMPAD6                //102 小键盘6
	VK_NUMPAD7                //103 小键盘7
	VK_NUMPAD8                //104 小键盘8
	VK_NUMPAD9                //105 小键盘9
	VK_MULTIPLY               //106 小键盘*
	VK_ADD                    //107 小键盘+
	_                         //108
	VK_SUBTRACT               //109 小键盘-
	VK_DECIMAL                //110 小键盘.
	VK_DIVIDE                 //111 小键盘/
	VK_F1                     //112 F1键
	VK_F2                     //113 F2键
	VK_F3                     //114 F3键
	VK_F4                     //115 F4键
	VK_F5                     //116 F5键
	VK_F6                     //117 F6键
	VK_F7                     //118 F7键
	VK_F8                     //119 F8键
	VK_F9                     //120 F9键
	VK_F10                    //121 F10键
	VK_F11                    //122 F11键
	VK_F12                    //123 F12键
	_                         //124
	_                         //125
	_                         //126
	_                         //127
	_                         //128
	_                         //129
	_                         //130
	_                         //131
	_                         //132
	_                         //133
	_                         //134
	_                         //135
	_                         //136
	_                         //137
	_                         //138
	_                         //139
	_                         //140
	_                         //141
	_                         //142
	_                         //143
	VK_NUMLOCK                //144 Num Lock键
	VK_SCROLL                 //145 Scroll Lock键

)

var arrKeyTable []keyTable

// https://www.cnblogs.com/wqw/archive/2009/08/30/1556618.html
func init() {
	arrKeyTable = make([]keyTable, 256)

	arrKeyTable[97].vk, arrKeyTable[97].scan = 0x61, 0x1E   //a
	arrKeyTable[98].vk, arrKeyTable[98].scan = 0x62, 0x30   //b
	arrKeyTable[99].vk, arrKeyTable[99].scan = 0x63, 0x2E   //c
	arrKeyTable[100].vk, arrKeyTable[100].scan = 0x64, 0x20 //d
	arrKeyTable[101].vk, arrKeyTable[101].scan = 0x65, 0x12 //e
	arrKeyTable[102].vk, arrKeyTable[102].scan = 0x66, 0x21 //f
	arrKeyTable[103].vk, arrKeyTable[103].scan = 0x67, 0x22 //g
	arrKeyTable[104].vk, arrKeyTable[104].scan = 0x68, 0x23 //h
	arrKeyTable[105].vk, arrKeyTable[105].scan = 0x69, 0x17 //i
	arrKeyTable[106].vk, arrKeyTable[106].scan = 0x6A, 0x24 //j
	arrKeyTable[107].vk, arrKeyTable[107].scan = 0x6B, 0x25 //k
	arrKeyTable[108].vk, arrKeyTable[108].scan = 0x6C, 0x26 //l
	arrKeyTable[109].vk, arrKeyTable[109].scan = 0x6D, 0x32 //m
	arrKeyTable[110].vk, arrKeyTable[110].scan = 0x6E, 0x31 //n
	arrKeyTable[111].vk, arrKeyTable[111].scan = 0x6F, 0x18 //o
	arrKeyTable[112].vk, arrKeyTable[112].scan = 0x70, 0x19 //p
	arrKeyTable[113].vk, arrKeyTable[113].scan = 0x71, 0x10 //q
	arrKeyTable[114].vk, arrKeyTable[114].scan = 0x72, 0x13 //r
	arrKeyTable[115].vk, arrKeyTable[115].scan = 0x73, 0x1F //s
	arrKeyTable[116].vk, arrKeyTable[116].scan = 0x74, 0x14 //t
	arrKeyTable[117].vk, arrKeyTable[117].scan = 0x75, 0x16 //u
	arrKeyTable[118].vk, arrKeyTable[118].scan = 0x76, 0x2F //v
	arrKeyTable[119].vk, arrKeyTable[119].scan = 0x77, 0x11 //w
	arrKeyTable[120].vk, arrKeyTable[120].scan = 0x78, 0x2D //x
	arrKeyTable[121].vk, arrKeyTable[121].scan = 0x79, 0x15 //y
	arrKeyTable[122].vk, arrKeyTable[122].scan = 0x7A, 0x2C //z
	arrKeyTable[97].vk, arrKeyTable[97].scan = 0x61, 0x1E   //A
	arrKeyTable[98].vk, arrKeyTable[98].scan = 0x62, 0x30   //B
	arrKeyTable[99].vk, arrKeyTable[99].scan = 0x63, 0x2E   //C
	arrKeyTable[100].vk, arrKeyTable[100].scan = 0x64, 0x20 //D
	arrKeyTable[101].vk, arrKeyTable[101].scan = 0x65, 0x12 //E
	arrKeyTable[102].vk, arrKeyTable[102].scan = 0x66, 0x21 //F
	arrKeyTable[103].vk, arrKeyTable[103].scan = 0x67, 0x22 //G
	arrKeyTable[104].vk, arrKeyTable[104].scan = 0x68, 0x23 //H
	arrKeyTable[105].vk, arrKeyTable[105].scan = 0x69, 0x17 //I
	arrKeyTable[106].vk, arrKeyTable[106].scan = 0x6A, 0x24 //J
	arrKeyTable[107].vk, arrKeyTable[107].scan = 0x6B, 0x25 //K
	arrKeyTable[108].vk, arrKeyTable[108].scan = 0x6C, 0x26 //L
	arrKeyTable[109].vk, arrKeyTable[109].scan = 0x6D, 0x32 //M
	arrKeyTable[110].vk, arrKeyTable[110].scan = 0x6E, 0x31 //N
	arrKeyTable[111].vk, arrKeyTable[111].scan = 0x6F, 0x18 //O
	arrKeyTable[112].vk, arrKeyTable[112].scan = 0x70, 0x19 //P
	arrKeyTable[113].vk, arrKeyTable[113].scan = 0x71, 0x10 //Q
	arrKeyTable[114].vk, arrKeyTable[114].scan = 0x72, 0x13 //R
	arrKeyTable[115].vk, arrKeyTable[115].scan = 0x73, 0x1F //S
	arrKeyTable[116].vk, arrKeyTable[116].scan = 0x74, 0x14 //T
	arrKeyTable[117].vk, arrKeyTable[117].scan = 0x75, 0x16 //U
	arrKeyTable[118].vk, arrKeyTable[118].scan = 0x76, 0x2F //V
	arrKeyTable[119].vk, arrKeyTable[119].scan = 0x77, 0x11 //W
	arrKeyTable[120].vk, arrKeyTable[120].scan = 0x78, 0x2D //X
	arrKeyTable[121].vk, arrKeyTable[121].scan = 0x79, 0x15 //Y
	arrKeyTable[122].vk, arrKeyTable[122].scan = 0x7A, 0x2C //Z
	arrKeyTable[48].vk, arrKeyTable[48].scan = 0x30, 0x0b   //0
	arrKeyTable[49].vk, arrKeyTable[49].scan = 0x31, 0x02   //1
	arrKeyTable[50].vk, arrKeyTable[50].scan = 0x32, 0x03   //2
	arrKeyTable[51].vk, arrKeyTable[51].scan = 0x33, 0x04   //3
	arrKeyTable[52].vk, arrKeyTable[52].scan = 0x34, 0x05   //4
	arrKeyTable[53].vk, arrKeyTable[53].scan = 0x35, 0x06   //5
	arrKeyTable[54].vk, arrKeyTable[54].scan = 0x36, 0x07   //6
	arrKeyTable[55].vk, arrKeyTable[55].scan = 0x37, 0x08   //7
	arrKeyTable[56].vk, arrKeyTable[56].scan = 0x38, 0x09   //8
	arrKeyTable[57].vk, arrKeyTable[57].scan = 0x39, 0x0a   //9

}

//func PressKeys(wVk []uint16) (ret uint32) {
//	count := len(wVk)
//	if 0 == count {
//		return 0
//	}

//	ips := make([]user32.KEYBD_INPUT, count)
//	for i := range ips {
//		ips[i] = user32.KEYBD_INPUT{user32.INPUT_KEYBOARD,
//			user32.KEYBDINPUT{}}

//		ips[i].Ki.WVk = wVk[i]
//	}

//	return
//}

func Press(wVk uint16) (ret uint32) {
	ip := user32.KEYBD_INPUT{user32.INPUT_KEYBOARD,
		user32.KEYBDINPUT{}}
	ip.Ki.WVk = uint16(wVk)
	ip.Ki.WScan = uint16(user32.MapVirtualKey(uint32(wVk), 0))

	ipSize := int32(unsafe.Sizeof(ip))
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)
	ip.Ki.DwFlags = user32.KEYEVENTF_KEYUP
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)

	return
}

// https://www.cnblogs.com/tary2017/articles/8031782.html   wMsg参数常量值

//lParam的0到15位为该键在键盘上的重复次数，经常设为1，即按键1次；
//16至23位为键盘的扫描码，通过MapVirtualKey配合其参数可以得到；
//24位为扩展键，即某些右ALT和CTRL；29、30、31位按照说明设置即可
//（第30位对于keydown在和shift等结合的时候通常要设置为1）。  0100 0000 0000 0000 0000 0000 0000 0000
func SendMessage(hWnd uint32, wVk uint16) (ret uint32) {
	ret = user32.SendMessage(hWnd, user32.WM_KEYDOWN, wVk, 0x00000001|user32.MapVirtualKey(uint32(wVk), 0)|0x40000000)
	ret = user32.SendMessage(hWnd, user32.WM_KEYUP, wVk, 0x00000001|user32.MapVirtualKey(uint32(wVk), 0))

	return
}

//func Press(wVk int) (ret uint32) {
//	user32.Keybd_event(arrKeyTable[wVk].vk, arrKeyTable[wVk].scan, 0, 0)
//	user32.Keybd_event(arrKeyTable[wVk].vk, arrKeyTable[wVk].scan, 2, 0)
//	return
//}

func Free() {
	user32.Free()
}
