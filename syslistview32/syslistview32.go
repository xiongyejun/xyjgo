package syslistview32

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/user32"

	"github.com/axgle/mahonia"
	"github.com/xiongyejun/xyjgo/winAPI/kernel32"
)

type ListView32 struct {
	hwndFolder   uint32 // CabinetWClass
	hwndListView uint32 // SysListView32
	//	hwndSysHeader win.HWND // SysHeader32
}

type iItemStruct struct {
	mask       uint32
	iItem      int32
	iSubItem   int32
	state      uint32
	stateMask  uint32
	pszText    uintptr
	cchTextMax int32
	iImage     int32
	IParam     int32
	// _WIN32_ie >= 0X0300
	iIndent int32
	// _WIN32_WINNT >= 0X0501
	iGroupId  int32
	cColumns  uint32
	puColumns uint32
	// _WIN32_WINNT >= 0X0600
	piColFmt int32
	iGroup   int32
}

func NewListView32(hwndFolder uint32) (l *ListView32, err error) {
	l = new(ListView32)

	if hwndFolder == 0 {
		return nil, errors.New("hwndFolder不能为0")
	}
	l.hwndFolder = hwndFolder

	if l.hwndListView, err = FindWindowEx(l.hwndFolder, []string{"SHELLDLL_DefView",
		"DUIViewWndClassName",
		"DirectUIHWND",
		"CtrlNotifySink",
		"SysListView32"}); err != nil {
		return
	}
	return
}

func FindWindowEx(hwndParent uint32, wins []string) (ret uint32, err error) {
	for i := range wins {
		if hwndParent = user32.FindWindowEx(hwndParent, 0, wins[i], ""); 0 == hwndParent {
			return 0, errors.New("没有找到SHELLDLL_DefView窗口")
		}
	}
	return hwndParent, nil
}

const (
	LVM_FIRST        = 0x1000
	LVM_GETITEMW     = (LVM_FIRST + 75)
	LVM_GETITEMA     = (LVM_FIRST + 5)
	LVM_GETNEXTITEM  = (LVM_FIRST + 12)
	LVNI_SELECTED    = 0x2
	LVIS_SELECTED    = 0x2
	LVIF_TEXT        = 0x1
	LVM_GETITEMCOUNT = (LVM_FIRST + 4)
	LVM_GETITEMTEXTA = (LVM_FIRST + 45)
	LVM_GETITEMTEXTW = (LVM_FIRST + 115)
	LVM_GETITEMSTATE = (LVM_FIRST + 44)

	MEM_RESERVE    = 0x2000
	MEM_COMMIT     = 0x1000
	PAGE_READWRITE = 0x4
	MEM_RELEASE    = 0x8000
)

func (me *ListView32) GetFolderPath() (ret string, err error) {
	var hwndPath uint32
	if hwndPath, err = FindWindowEx(me.hwndFolder, []string{
		"WorkerW",
		"ReBarWindow32",
		"ComboBoxEx32"}); err != nil {
		return
	}
	fmt.Printf("%x\r\n", hwndPath)
	iLen := user32.SendMessage(hwndPath, user32.WM_GETTEXTLENGTH, 0, 0) + 1
	if iLen == 1 {
		return "", errors.New("标题获取出错")
	}
	buf := make([]uint16, iLen)
	user32.SendMessage(hwndPath, user32.WM_GETTEXT, uintptr(iLen), uintptr(unsafe.Pointer(&buf[0])))
	ret = syscall.UTF16ToString(buf)
	if ret == "我的文档" {
		ret = os.Getenv("USERPROFILE") + `\My Documents`
	}
	return
}

func (me *ListView32) GetItemsCount() int {
	return int(user32.SendMessage(me.hwndListView, LVM_GETITEMCOUNT, 0, 0))
}
func (me *ListView32) GetSlectedItemIndex() int {
	return int(user32.SendMessage(me.hwndListView, LVM_GETNEXTITEM, 0, LVNI_SELECTED))
}
func (me *ListView32) GetItemString(index int) (ret string, err error) {
	if index < 0 || index > me.GetItemsCount() {
		return "", errors.New("GetItemString下标越界。")
	}

	var pid uint32
	user32.GetWindowThreadProcessId(me.hwndFolder, &pid)
	vProcessId := kernel32.OpenProcess(kernel32.PROCESS_ALL_ACCESS, 0, pid)

	if 0 == vProcessId {
		return "", errors.New("OpenProcess出错。")
	}
	defer kernel32.CloseHandle(vProcessId)
	l := new(iItemStruct)
	l.cchTextMax = 512
	l.iSubItem = 0
	nSize := int32(unsafe.Sizeof(*l)) + l.cchTextMax - 32

	vPointer := kernel32.VirtualAllocEx(vProcessId, 0, uintptr(nSize), MEM_COMMIT, PAGE_READWRITE)
	defer kernel32.VirtualFreeEx(vProcessId, vPointer, 0, MEM_RELEASE)

	pItem := kernel32.VirtualAllocEx(vProcessId, 0, uintptr(l.cchTextMax), MEM_COMMIT, PAGE_READWRITE)
	defer kernel32.VirtualFreeEx(vProcessId, pItem, 0, MEM_RELEASE)

	bItem := make([]byte, l.cchTextMax)
	l.pszText = pItem
	kernel32.WriteProcessMemory(vProcessId, vPointer, uintptr(unsafe.Pointer(l)), nSize, 0)
	user32.SendMessage(me.hwndListView, LVM_GETITEMTEXTA, uintptr(index), vPointer)
	kernel32.ReadProcessMemory(vProcessId, pItem, uintptr(unsafe.Pointer(&bItem[0])), l.cchTextMax, 0)

	return string(getByte(bItem)), nil
}

func getByte(b []byte) []byte {
	index := bytes.Index(b, []byte{0})
	if index > 0 {
		return tslByte(b[:index])
	} else {
		return tslByte(b)
	}
}

func tslByte(b []byte) []byte {
	decoder := mahonia.NewDecoder("gbk")
	return []byte(decoder.ConvertString(string(b)))
}
