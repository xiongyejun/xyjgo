// Operating memory 操作内容
package opmem

import (
	"errors"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/kernel32"
	"github.com/xiongyejun/xyjgo/winAPI/user32"
)

type OpMem struct {
	Process uint32
}

func New(processID uint32) (ret *OpMem, err error) {
	ret = new(OpMem)
	ret.Process = kernel32.OpenProcess(kernel32.PROCESS_ALL_ACCESS, 0, processID)

	return
}

func NewByWindow(window string) (ret *OpMem, err error) {
	hwnd := user32.FindWindow("", window)
	if hwnd == 0 {
		err = errors.New("没有找到窗口: " + window)
		return
	}
	var processID uint32
	if user32.GetWindowThreadProcessId(hwnd, &processID) == 0 {
		err = errors.New("GetWindowThreadProcessId 出错:")
		return
	}

	return New(processID)
}

func (me *OpMem) Scan() (ret []*kernel32.MEMORY_BASIC_INFORMATION, err error) {
	/*
		--------4G------ 0xFFFFFFFF
		Operating System
		--------2G------c
		User Process
	*/
	var baseAddress uint32 = 0x400000

	for baseAddress < 0x80000000 {
		lpBuffer := kernel32.MEMORY_BASIC_INFORMATION{}

		cret := kernel32.VirtualQueryEx(me.Process, uintptr(baseAddress), uintptr(unsafe.Pointer(&lpBuffer.BaseAddress)), 28)

		if cret == 28 {
			if lpBuffer.AllocationProtect == kernel32.PAGE_READWRITE {
				ret = append(ret, &lpBuffer)
			}
			baseAddress = uint32(lpBuffer.RegionSize + lpBuffer.BaseAddress)
		} else {
			baseAddress += 1
		}
	}
	return
}

func (me *OpMem) Read(meminfo *kernel32.MEMORY_BASIC_INFORMATION) (ret []byte, err error) {
	ret = make([]byte, meminfo.RegionSize)
	if apiret := kernel32.ReadProcessMemory(me.Process, uintptr(meminfo.BaseAddress), uintptr(unsafe.Pointer(&ret[0])), meminfo.RegionSize, 0); apiret == 0 {
		err = errors.New("ReadProcessMemory返回为0")
		return
	}

	return
}
