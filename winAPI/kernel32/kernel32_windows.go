package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                uintptr
	getShortPathName   uintptr
	globalAlloc        uintptr
	globalFree         uintptr
	globalLock         uintptr
	globalUnlock       uintptr
	openProcess        uintptr
	virtualAllocEx     uintptr
	virtualFreeEx      uintptr
	virtualQueryEx     uintptr
	writeProcessMemory uintptr
	readProcessMemory  uintptr
	closeHandle        uintptr
	getLastError       uintptr
)

// GlobalAlloc flags
const (
	GHND          = 0x0042
	GMEM_FIXED    = 0x0000
	GMEM_MOVEABLE = 0x0002
	GMEM_ZEROINIT = 0x0040
	GPTR          = 0x004
)

func init() {
	// Library
	lib = win.MustLoadLibrary("kernel32.dll")

	// Functions
	getShortPathName = win.MustGetProcAddress(lib, "GetShortPathNameA")
	globalAlloc = win.MustGetProcAddress(lib, "GlobalAlloc")
	globalFree = win.MustGetProcAddress(lib, "GlobalFree")
	globalLock = win.MustGetProcAddress(lib, "GlobalLock")
	globalUnlock = win.MustGetProcAddress(lib, "GlobalUnlock")
	openProcess = win.MustGetProcAddress(lib, "OpenProcess")
	virtualAllocEx = win.MustGetProcAddress(lib, "VirtualAllocEx")
	virtualFreeEx = win.MustGetProcAddress(lib, "VirtualFreeEx")
	virtualQueryEx = win.MustGetProcAddress(lib, "VirtualQueryEx")
	writeProcessMemory = win.MustGetProcAddress(lib, "WriteProcessMemory")
	readProcessMemory = win.MustGetProcAddress(lib, "ReadProcessMemory")
	closeHandle = win.MustGetProcAddress(lib, "CloseHandle")
	getLastError = win.MustGetProcAddress(lib, "GetLastError")
}

func GlobalAlloc(uFlags uint32, dwBytes uintptr) uint32 {
	ret, _, _ := syscall.Syscall(globalAlloc, 2,
		uintptr(uFlags),
		dwBytes,
		0)

	return uint32(ret)
}

func GlobalFree(hMem uint32) uint32 {
	ret, _, _ := syscall.Syscall(globalFree, 1,
		uintptr(hMem),
		0,
		0)

	return uint32(ret)
}

func GlobalLock(hMem uint32) unsafe.Pointer {
	ret, _, _ := syscall.Syscall(globalLock, 1,
		uintptr(hMem),
		0,
		0)

	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem uint32) bool {
	ret, _, _ := syscall.Syscall(globalUnlock, 1,
		uintptr(hMem),
		0,
		0)

	return ret != 0
}

func GetShortPathName(longPath string, shortPath uintptr, shortBufferSize uint32) int32 {
	ret, _, _ := syscall.Syscall6(getShortPathName, 3,
		win.StrPtr(longPath, win.CODE_UTF16),
		shortPath,
		uintptr(shortBufferSize),
		0,
		0,
		0)

	return int32(ret)
}

const (
	SYNCHRONIZE               = 0x100000
	STANDARD_RIGHTS_REQUIRED  = 0xF0000
	PROCESS_ALL_ACCESS        = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFF)
	PROCESS_VM_OPERATION      = 0x8
	PROCESS_VM_READ           = 0x10
	PROCESS_VM_WRITE          = 0x20
	PROCESS_QUERY_INFORMATION = 0x0400
)

func OpenProcess(dwDesiredAccess, bInheritHandle, dwProcessId uint32) uint32 {
	ret, _, _ := syscall.Syscall(openProcess, 3,
		uintptr(dwDesiredAccess),
		uintptr(bInheritHandle),
		uintptr(dwProcessId))

	return uint32(ret)
}
func CloseHandle(hObject uint32) bool {
	ret, _, _ := syscall.Syscall(closeHandle, 1,
		uintptr(hObject),
		0,
		0)

	return ret != 0
}
func VirtualAllocEx(hProcess uint32, lpAddress uintptr, dwSize uintptr, flAllocationType, flProtect uint32) uintptr {
	ret, _, _ := syscall.Syscall6(virtualAllocEx, 5,
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(flAllocationType),
		uintptr(flProtect),
		0)

	return uintptr(ret)
}
func VirtualFreeEx(hProcess uint32, lpAddress uintptr, dwSize uintptr, dwFreeType uint32) int32 {
	ret, _, _ := syscall.Syscall6(virtualFreeEx, 4,
		uintptr(hProcess),
		lpAddress,
		dwSize,
		uintptr(dwFreeType),
		0,
		0)

	return int32(ret)
}

// AllocationProtect
const (
	PAGE_NOACCESS          = 0x01 // 禁止一切访问
	PAGE_READONLY          = 0x02
	PAGE_READWRITE         = 0x04
	PAGE_WRITECOPY         = 0x08
	PAGE_TARGETS_INVALID   = 0x40000000
	PAGE_TARGETS_NO_UPDATE = 0x40000000
	PAGE_EXECUTE           = 0x10 // 只允许执行代码，对该区域试图进行读写操作将引发访问违规。
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_EXECUTE_WRITECOPY = 0x80 // 对于该地址空间的区域，不管执行什么操作，都不会引发访问违规。如果试图在该页面上的内存中进行写入操作，就会将它自己的私有页面（受页文件的支持）拷贝赋予该进程。

	PAGE_GUARD        = 0x100 // 在页面上写入一个字节时使应用程序收到一个通知（通过一个异常条件）。该标志有一些非常巧妙的用法。Windows 2000在创建线程堆栈时使用该标志。
	PAGE_NOCACHE      = 0x200 // 停用已提交页面的高速缓存。一般情况下最好不要使用该标志，因为它主要是供需要处理内存缓冲区的硬件设备驱动程序的开发人员使用的。
	PAGE_WRITECOMBINE = 0x400
)

type MEMORY_BASIC_INFORMATION struct {
	/*
		使用VirtualAlloc分配虚拟内存时,您将始终获得一个BaseAddress等于AllocationBase的块.但是,如果您随后更改了此块中一个或多个页面的页面保护,那么您可以观察到此块被细分为不同的BaseAddress.
		BaseAddress一定在AllocationBase之内
	*/
	BaseAddress       int32 // 区域基地址。
	AllocationBase    int32 // 分配基地址。
	AllocationProtect int32 // 区域被初次保留时赋予的保护属性。
	RegionSize        int32 // 区域大小（以字节为计量单位）。
	State             int32 // 状态（MEM_FREE、MEM_RESERVE或 MEM_COMMIT）。
	Protect           int32 // 保护属性。
	Type              int32 // 类型。
}

func VirtualQueryEx(hProcess uint32, lpAddress, lpBuffer uintptr, dwLength int32) int32 {
	ret, _, _ := syscall.Syscall6(virtualQueryEx, 4,
		uintptr(hProcess),
		lpAddress,
		lpBuffer,
		uintptr(dwLength),
		0,
		0)

	return int32(ret)
}

func WriteProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesWrittenu int32) int32 {
	ret, _, _ := syscall.Syscall6(writeProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesWrittenu),
		0)

	return int32(ret)
}

func ReadProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesWrittenu int32) int32 {
	ret, _, _ := syscall.Syscall6(readProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesWrittenu),
		0)

	return int32(ret)
}

func GetLastError() uint32 {
	ret, _, _ := syscall.Syscall(getLastError, 0,
		0,
		0,
		0)

	return uint32(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
