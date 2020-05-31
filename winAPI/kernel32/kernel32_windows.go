package kernel32

import (
	"bytes"
	"errors"
	"syscall"
	"unsafe"

	"github.com/xiongyejun/xyjgo/winAPI/win"
)

var (
	lib                      uintptr
	getShortPathName         uintptr
	globalAlloc              uintptr
	globalFree               uintptr
	globalLock               uintptr
	globalUnlock             uintptr
	openProcess              uintptr
	virtualAllocEx           uintptr
	virtualFreeEx            uintptr
	virtualQueryEx           uintptr
	writeProcessMemory       uintptr
	readProcessMemory        uintptr
	closeHandle              uintptr
	createToolhelp32Snapshot uintptr
	process32First           uintptr
	process32Next            uintptr
	getLastError             uintptr
	getModuleHandle          uintptr
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
	createToolhelp32Snapshot = win.MustGetProcAddress(lib, "CreateToolhelp32Snapshot")
	process32Next = win.MustGetProcAddress(lib, "Process32Next")
	process32First = win.MustGetProcAddress(lib, "Process32First")
	getLastError = win.MustGetProcAddress(lib, "GetLastError")
	getModuleHandle = win.MustGetProcAddress(lib, "GetModuleHandleW")
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

const (
	TH32CS_INHERIT      = 0x80000000                                                                         // 声明快照句柄是可继承的。
	TH32CS_SNAPHEAPLIST = 0x00000001                                                                         // 在快照中包含在th32ProcessID中指定的进程的所有的堆。
	TH32CS_SNAPMODULE   = 0x00000008                                                                         // 在快照中包含在th32ProcessID中指定的进程的所有的模块。
	TH32CS_SNAPPROCESS  = 0x00000002                                                                         // 在快照中包含系统中所有的进程。
	TH32CS_SNAPTHREAD   = 0x00000004                                                                         // 在快照中包含系统中所有的线程。
	H32CS_SNAPALL       = (TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE) // 在快照中包含系统中所有的进程和线程。

	INVALID_HANDLE_VALUE = -1
)

func CreateToolhelp32Snapshot(dwFlags, th32ProcessID uint32) uint32 {
	ret, _, _ := syscall.Syscall(createToolhelp32Snapshot, 2,
		uintptr(dwFlags),
		uintptr(th32ProcessID),
		0)

	return uint32(ret)
}

type PROCESSENTRY32 struct {
	Size            uint32 // 结构大小；
	Usage           uint32 // 此进程的引用计数；
	ProcessID       uint32 // 进程ID;
	DefaultHeapID   uint32 // 进程默认堆ID；
	ModuleID        uint32 // 进程模块ID；
	Threads         uint32 // 此进程开启的线程计数；
	ParentProcessID uint32 // 父进程ID；
	PriClassBase    uint32 // 线程优先权；
	Flags           uint32 // 保留；
	ExeFile         [1024]byte
	//     szExeFile[MAX_PATH] uint32 // 进程全名；
}

func Process32First(hSnapshot uint32, lppe uintptr) uint32 {
	ret, _, _ := syscall.Syscall(process32First, 2,
		uintptr(hSnapshot),
		lppe,
		0)

	return uint32(ret)
}
func Process32Next(hSnapshot uint32, lppe uintptr) uint32 {
	ret, _, _ := syscall.Syscall(process32Next, 2,
		uintptr(hSnapshot),
		lppe,
		0)

	return uint32(ret)
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

func WriteProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesWritten int32) int32 {
	ret, _, _ := syscall.Syscall6(writeProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesWritten),
		0)

	return int32(ret)
}

func ReadProcessMemory(hProcess uint32, lpBaseAddress uintptr, lpBuffer uintptr, nSize, lpNumberOfBytesRead int32) int32 {
	ret, _, _ := syscall.Syscall6(readProcessMemory, 5,
		uintptr(hProcess),
		lpBaseAddress,
		lpBuffer,
		uintptr(nSize),
		uintptr(lpNumberOfBytesRead),
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

func GetPIDByProcessName(processName string) (ret uint32, err error) {
	hProcessSnap := CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0)
	if hProcessSnap == 0 {
		err = errors.New("CreateToolhelp32Snapshot出错。 ")
		return
	}
	defer CloseHandle(hProcessSnap)

	pentry := PROCESSENTRY32{}
	pentry.Size = 1024 + 9*4
	b := Process32First(hProcessSnap, uintptr(unsafe.Pointer(&pentry.Size)))
	for b == 1 {
		index := bytes.Index(pentry.ExeFile[:], []byte{0})
		str := string(pentry.ExeFile[:index])
		if str == processName {
			ret = pentry.ProcessID
			return
		}
		b = Process32Next(hProcessSnap, uintptr(unsafe.Pointer(&pentry.Size)))
	}

	return
}

func GetModuleHandle(lpModuleName *uint16) win.HINSTANCE {
	ret, _, _ := syscall.Syscall(getModuleHandle, 1,
		uintptr(unsafe.Pointer(lpModuleName)),
		0,
		0)

	return win.HINSTANCE(ret)
}

func Free() {
	syscall.FreeLibrary(syscall.Handle(lib))
}
