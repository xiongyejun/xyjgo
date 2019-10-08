package pe

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/xiongyejun/xyjgo/fileFormat/coff"
)

type SHORT = int16
type WORD = uint16
type DWORD = uint32
type ULONGLONG = uint64

func init() {

}

// winnt.h
type IMAGE_DOS_HEADER struct {
	E_magic    WORD
	E_cblp     WORD
	E_cp       WORD
	E_crlc     WORD
	E_cparhdr  WORD
	E_minalloc WORD
	E_maxalloc WORD
	E_ss       WORD
	E_sp       WORD
	E_csum     WORD
	E_ip       WORD
	E_cs       WORD
	E_lfarlc   WORD
	E_ovno     WORD
	E_res      [4]WORD
	E_oemid    WORD
	E_oeminfo  WORD
	E_res2     [10]WORD
	E_lfanew   int32 // 表明IMAGE_NT_HEADERS在文件中的偏移
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress DWORD
	Size           DWORD
}

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       DWORD
	TimeDateStamp         DWORD
	MajorVersion          WORD
	MinorVersion          WORD
	Name                  DWORD
	Base                  DWORD
	NumberOfFunctions     DWORD
	NumberOfNames         DWORD
	AddressOfFunctions    DWORD
	AddressOfNames        DWORD
	AddressOfNameOrdinals DWORD
}

const IMAGE_NUMBEROF_DIRECTORY_ENTRIES = 16

type IMAGE_NT_HEADERS struct {
	Signature      DWORD
	FileHeader     coff.IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER
}

/*
	PE文件格式
-------------------------------
	IMAGE_DOS_HEADER
	Image DOS Stub
	IMAGE_NT_HEADERS
	IMAGE_FILE_HEADERS
	IMAGE_OPTIONAL_HEADER
	IMAGE_SECTION_HEADER[]
	.text
	.data
	.drectve
	.debug$S
	...
	other sections
	SYMBOL Table
*/
type PE struct {
	f         *os.File
	DosHeader *IMAGE_DOS_HEADER
	NTHeader  *IMAGE_NT_HEADERS
	Sections  []*coff.IMAGE_SECTION_HEADER
	ExportDir *IMAGE_EXPORT_DIRECTORY
}

// 解析文件
func (me *PE) Parse(f *os.File) (err error) {
	me.f = f
	if err = me.readDosHeader(); err != nil {
		return
	}
	if _, err = me.f.Seek(int64(me.DosHeader.E_lfanew), 0); err != nil {
		return
	}
	if err = me.readNTHeader(); err != nil {
		return
	}
	if err = me.readSections(); err != nil {
		return
	}
	// #define IMAGE_DIRECTORY_ENTRY_EXPORT 0
	// #define IMAGE_DIRECTORY_ENTRY_IMPORT 1
	// #define IMAGE_DIRECTORY_ENTRY_RESOURCE 2
	// #define IMAGE_DIRECTORY_ENTRY_EXCEPTION 3
	// #define IMAGE_DIRECTORY_ENTRY_SECURITY 4
	// #define IMAGE_DIRECTORY_ENTRY_BASERELOC 5
	// #define IMAGE_DIRECTORY_ENTRY_DEBUG 6
	// #define IMAGE_DIRECTORY_ENTRY_ARCHITECTURE 7
	// #define IMAGE_DIRECTORY_ENTRY_GLOBALPTR 8
	// #define IMAGE_DIRECTORY_ENTRY_TLS 9
	// #define IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG 10
	// #define IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT 11
	// #define IMAGE_DIRECTORY_ENTRY_IAT 12
	// #define IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT 13
	// #define IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR 14

	// DataDirectory共16个元素，第1个元素就是导出表的结构的地址和长度
	if _, err = me.f.Seek(int64(me.NTHeader.OptionalHeader.DataDirectory[0].VirtualAddress), 0); err != nil {
		return
	}
	if err = me.readExportDir(); err != nil {
		return
	}

	return
}
func (me *PE) readDosHeader() (err error) {
	me.DosHeader = new(IMAGE_DOS_HEADER)
	dosHeaderSize := binary.Size(me.DosHeader)
	var b []byte = make([]byte, dosHeaderSize)
	var n int
	n, err = me.f.Read(b)
	if err != nil {
		return
	}
	if n < dosHeaderSize {
		err = errors.New("未能读取到【IMAGE_DOS_HEADER】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.DosHeader); err != nil {
		return
	}
	if me.DosHeader.E_magic != 0x5A4D {
		err = errors.New("IMAGE_DOS_HEADER.DosHeader != 0x5A4D(MZ)")
		return
	}
	return
}

func (me *PE) readNTHeader() (err error) {
	me.NTHeader = new(IMAGE_NT_HEADERS)
	NTHeaderSize := binary.Size(me.NTHeader)
	var b []byte = make([]byte, NTHeaderSize)
	var n int
	n, err = me.f.Read(b)
	if err != nil {
		return
	}
	if n < NTHeaderSize {
		err = errors.New("未能读取到【IMAGE_NT_HEADERS】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.NTHeader); err != nil {
		return
	}

	if me.NTHeader.Signature != 0x00004550 {
		err = errors.New("IMAGE_NT_HEADERS.Signature != 0x00004550(PE)")
		return
	}
	return
}

// 读取sections
func (me *PE) readSections() (err error) {
	me.Sections = make([]*coff.IMAGE_SECTION_HEADER, me.NTHeader.FileHeader.NumberOfSections)
	sectionSize := binary.Size(new(coff.IMAGE_SECTION_HEADER))

	// 逐个读取section
	for i := range me.Sections {
		var b []byte = make([]byte, sectionSize)
		var n int
		n, err = me.f.Read(b)
		if err != nil {
			return
		}
		if n < sectionSize {
			err = errors.New("未能读取到【IMAGE_SECTION_HEADER】足够的字节数。")
			return
		}
		tmp := new(coff.IMAGE_SECTION_HEADER)
		if err = byte2struct(b, tmp); err != nil {
			return
		}
		me.Sections[i] = tmp
	}

	return
}

func (me *PE) readExportDir() (err error) {
	me.ExportDir = new(IMAGE_EXPORT_DIRECTORY)
	ExportDirSize := binary.Size(me.ExportDir)
	var b []byte = make([]byte, ExportDirSize)
	var n int
	n, err = me.f.Read(b)
	if err != nil {
		return
	}
	if n < ExportDirSize {
		err = errors.New("未能读取到【IMAGE_EXPORT_DIRECTORY】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.ExportDir); err != nil {
		return
	}

	return
}

func (me *IMAGE_OPTIONAL_HEADER) GetPrintStr() string {
	return "OPTIONAL HEADER VALUES\n" +
		fmt.Sprintf("%16X magic # (PE32)\n", me.Magic) +
		fmt.Sprintf("%13d.00 linker version\n", me.MajorLinkerVersion) +
		fmt.Sprintf("%16X size of code\n", me.SizeOfCode) +
		fmt.Sprintf("%16X size of initialized data\n", me.SizeOfInitializedData) +
		fmt.Sprintf("%16X size of uninitialized data\n", me.SizeOfUninitializedData) +
		fmt.Sprintf("%16X entry point\n", me.AddressOfEntryPoint) +
		fmt.Sprintf("%16X base of code\n", me.BaseOfCode) +
		fmt.Sprintf("%16X image base\n", me.ImageBase) +
		fmt.Sprintf("%16X section alignment\n", me.SectionAlignment) +
		fmt.Sprintf("%16X file alignment\n", me.FileAlignment) +
		fmt.Sprintf("%16X operating system version\n", me.MajorOperatingSystemVersion) +
		fmt.Sprintf("%16X image version\n", me.MajorImageVersion) +
		fmt.Sprintf("%16X subsystem version\n", me.MajorSubsystemVersion) +
		fmt.Sprintf("%16X Win32 version\n", me.Win32VersionValue) +
		fmt.Sprintf("%16X size of image\n", me.SizeOfImage) +
		fmt.Sprintf("%16X size of headers\n", me.SizeOfHeaders) +
		fmt.Sprintf("%16X checksum\n", me.CheckSum) +
		fmt.Sprintf("%16X subsystem (Windows GUI)\n", me.Subsystem) +
		fmt.Sprintf("%16X DLL characteristics\n", me.DllCharacteristics) +
		fmt.Sprintf("%16X size of stack reserve\n", me.SizeOfStackReserve) +
		fmt.Sprintf("%16X size of stack commit\n", me.SizeOfStackCommit) +
		fmt.Sprintf("%16X size of heap header\n", me.SizeOfHeaders) +
		fmt.Sprintf("%16X size of heap reserve\n", me.SizeOfHeapReserve) +
		fmt.Sprintf("%16X size of heap commit\n", me.SizeOfHeapCommit) +
		fmt.Sprintf("%16X loader flags\n", me.LoaderFlags) +
		fmt.Sprintf("%16X number of directories\n", len(me.DataDirectory))
}

func byte2struct(b []byte, pStruct interface{}) error {
	buf := bytes.NewBuffer(b)
	return binary.Read(buf, binary.LittleEndian, pStruct)
}
