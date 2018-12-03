package keyboard

import (
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type keyTable struct {
	vk   byte
	scan byte
}

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

func Press(wVk int) (ret uint32) {
	ip := user32.KEYBD_INPUT{user32.INPUT_KEYBOARD,
		user32.KEYBDINPUT{}}
	ip.Ki.WVk = uint16(arrKeyTable[wVk].vk)
	ip.Ki.WScan = uint16(arrKeyTable[wVk].scan)

	ipSize := int32(unsafe.Sizeof(ip))
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)
	ip.Ki.DwFlags = user32.KEYEVENTF_KEYUP
	ret = user32.SendInput(1, unsafe.Pointer(&ip), ipSize)

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
