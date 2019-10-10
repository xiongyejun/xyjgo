package pe

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

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
type EXPORT_DIRECTORY_Info struct {
	ModuleName string // 指向导出表Name字段，内容存储的是模块函数名称的ASCII字符
	FuncNames  []string
	Ordinals   []int16 // 函数序号
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
	f             *os.File
	DosHeader     *IMAGE_DOS_HEADER
	NTHeader      *IMAGE_NT_HEADERS
	Sections      []*coff.IMAGE_SECTION_HEADER
	ExportDir     *IMAGE_EXPORT_DIRECTORY
	ExportDirInfo *EXPORT_DIRECTORY_Info
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
	if _, err = me.f.Seek(int64(me.VA2FOA(me.NTHeader.OptionalHeader.FDataDirectory()[0].VirtualAddress)), 0); err != nil {
		return
	}
	if err = me.readExportDir(); err != nil {
		return
	}

	if me.ExportDirInfo, err = me.ExportDir.readInfo(me); err != nil {
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
		return errors.New("readDosHeader err:" + err.Error())
	}
	if n < dosHeaderSize {
		err = errors.New("未能读取到【IMAGE_DOS_HEADER】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.DosHeader); err != nil {
		return errors.New("readDosHeader byte2struct err:" + err.Error())
	}
	if me.DosHeader.E_magic != 0x5A4D {
		err = errors.New("IMAGE_DOS_HEADER.DosHeader != 0x5A4D(MZ)")
		return
	}
	return
}

func (me *PE) readNTHeader() (err error) {
	me.NTHeader = new(IMAGE_NT_HEADERS)
	// 读取Signature
	NTHeaderSize := binary.Size(me.NTHeader.Signature)
	var b []byte = make([]byte, NTHeaderSize)
	var n int
	n, err = me.f.Read(b)
	if err != nil {
		return errors.New("readNTHeader Signature err:" + err.Error())
	}
	if n < NTHeaderSize {
		err = errors.New("未能读取到【IMAGE_NT_HEADERS.Signature】足够的字节数。")
		return
	}

	if err = byte2struct(b, &me.NTHeader.Signature); err != nil {
		return errors.New("readNTHeader byte2struct err:" + err.Error())
	}

	if me.NTHeader.Signature != 0x00004550 {
		err = errors.New("IMAGE_NT_HEADERS.Signature != 0x00004550(PE)")
		return
	}

	// 读取FileHeader
	NTHeaderSize = binary.Size(me.NTHeader.FileHeader)
	b = make([]byte, NTHeaderSize)
	n, err = me.f.Read(b)
	if err != nil {
		return errors.New("readNTHeader FileHeader err:" + err.Error())
	}
	if n < NTHeaderSize {
		err = errors.New("未能读取到【IMAGE_NT_HEADERS.FileHeader】足够的字节数。")
		return
	}

	if err = byte2struct(b, &me.NTHeader.FileHeader); err != nil {
		return errors.New("readNTHeader me.NTHeader.FileHeader byte2struct err:" + err.Error())
	}

	// 根据FileHeader.Machine，分别按32位和64位读取IMAGE_OPTIONAL_HEADER
	if me.NTHeader.FileHeader.Machine == 0x14c {
		tmp := IMAGE_OPTIONAL_HEADER32{}
		tmpSize := binary.Size(tmp)
		b = make([]byte, tmpSize)
		n, err = me.f.Read(b)
		if err != nil {
			return errors.New("readNTHeader IMAGE_OPTIONAL_HEADER32 err:" + err.Error())
		}
		if n < tmpSize {
			err = errors.New("未能读取到【IMAGE_OPTIONAL_HEADER32】足够的字节数。")
			return
		}

		if err = byte2struct(b, &tmp); err != nil {
			return errors.New("readNTHeader IMAGE_OPTIONAL_HEADER32 byte2struct err:" + err.Error())
		}

		me.NTHeader.OptionalHeader = IMAGE_OPTIONAL_HEADER{}
		me.NTHeader.OptionalHeader.I_IMAGE_OPTIONAL_HEADER = tmp
	} else if me.NTHeader.FileHeader.Machine == 0x8664 {
		tmp := IMAGE_OPTIONAL_HEADER64{}
		tmpSize := binary.Size(tmp)
		b = make([]byte, tmpSize)
		n, err = me.f.Read(b)
		if err != nil {
			return errors.New("readNTHeader IMAGE_OPTIONAL_HEADER64 err:" + err.Error())
		}
		if n < tmpSize {
			err = errors.New("未能读取到【IMAGE_OPTIONAL_HEADER64】足够的字节数。")
			return
		}

		if err = byte2struct(b, &tmp); err != nil {
			return errors.New("readNTHeader IMAGE_OPTIONAL_HEADER64 byte2struct err:" + err.Error())
		}
		me.NTHeader.OptionalHeader = IMAGE_OPTIONAL_HEADER{}
		me.NTHeader.OptionalHeader.I_IMAGE_OPTIONAL_HEADER = tmp
	} else {
		return errors.New("readNTHeader err:" + fmt.Sprintf("未知NTHeader.FileHeader.Machine 0x%x", me.NTHeader.FileHeader.Machine))
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
			return errors.New("readSections err:" + err.Error())
		}
		if n < sectionSize {
			err = errors.New("未能读取到【IMAGE_SECTION_HEADER】足够的字节数。")
			return
		}
		tmp := new(coff.IMAGE_SECTION_HEADER)
		if err = byte2struct(b, tmp); err != nil {
			return errors.New("readSections byte2struct err:" + err.Error())
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
		return errors.New("readExportDir err:" + err.Error())
	}
	if n < ExportDirSize {
		err = errors.New("未能读取到【IMAGE_EXPORT_DIRECTORY】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.ExportDir); err != nil {
		return errors.New("readExportDir byte2struct err:" + err.Error())
	}

	if me.ExportDir.Characteristics != 0x0 {
		err = errors.New("IMAGE_EXPORT_DIRECTORY.Characteristics != 0x0")
		return
	}
	return
}

// 读取导出表的信息
func (me *IMAGE_EXPORT_DIRECTORY) readInfo(pe *PE) (ret *EXPORT_DIRECTORY_Info, err error) {
	ret = new(EXPORT_DIRECTORY_Info)
	if _, err = pe.f.Seek(int64(pe.VA2FOA(me.Name)), 0); err != nil {
		err = errors.New("readInfo Seek me.Name err:" + err.Error())
		return
	}
	// 读取dll的名称
	var bname []byte
	if bname, err = readByteEndOnZero(pe.f); err != nil {
		err = errors.New("readInfo ModuleName readByteEndOnZero err:" + err.Error())
		return
	}
	ret.ModuleName = string(bname)

	// 读取导出函数的名称
	ret.FuncNames = make([]string, me.NumberOfNames)
	// AddressOfNames记录的是导出函数名字的RVA，有NumberOfNames个
	var rva []DWORD = make([]DWORD, me.NumberOfNames)
	if _, err = pe.f.Seek(int64(pe.VA2FOA(me.AddressOfNames)), 0); err != nil {
		return
	}
	var b []byte = make([]byte, 4*int(me.NumberOfNames))
	if _, err = pe.f.Read(b); err != nil {
		return
	}
	if err = byte2struct(b, rva); err != nil {
		return
	}
	for i := range ret.FuncNames {
		if _, err = pe.f.Seek(int64(pe.VA2FOA(rva[i])), 0); err != nil {
			err = errors.New("readInfo Seek rva[i] err:" + err.Error())
			return
		}
		if bname, err = readByteEndOnZero(pe.f); err != nil {
			err = errors.New("readInfo FuncNames readByteEndOnZero err:" + err.Error())
			return
		}
		ret.FuncNames[i] = string(bname)
	}
	// 读取Ordinals
	ret.Ordinals = make([]int16, me.NumberOfFunctions)
	// 编程的时候为DWORD类型，但却是WORD类型
	b = make([]byte, 2*int(me.NumberOfFunctions))
	if _, err = pe.f.Seek(int64(pe.VA2FOA(me.AddressOfNameOrdinals)), 0); err != nil {
		err = errors.New("readInfo Seek AddressOfNameOrdinals err:" + err.Error())
		return
	}
	if _, err = pe.f.Read(b); err != nil {
		err = errors.New("readInfo f.Read AddressOfNameOrdinals err:" + err.Error())
		return
	}
	if err = byte2struct(b, ret.Ordinals); err != nil {
		err = errors.New("readInfo Ordinals byte2struct err:" + err.Error())
		return
	}
	// 导出函数序号 = 函数入口地址序号 + 基数
	for i := range ret.Ordinals {
		ret.Ordinals[i] += int16(me.Base)
	}
	return
}

func (me *IMAGE_OPTIONAL_HEADER) String() string {
	return "OPTIONAL HEADER VALUES\n" +
		fmt.Sprintf("%16X magic # (PE32)\n", me.FMagic()) +
		fmt.Sprintf("%13d.00 linker version\n", me.FMajorLinkerVersion()) +
		fmt.Sprintf("%16X size of code\n", me.FSizeOfCode()) +
		fmt.Sprintf("%16X size of initialized data\n", me.FSizeOfInitializedData()) +
		fmt.Sprintf("%16X size of uninitialized data\n", me.FSizeOfUninitializedData()) +
		fmt.Sprintf("%16X entry point\n", me.FAddressOfEntryPoint()) +
		fmt.Sprintf("%16X base of code\n", me.FBaseOfCode()) +
		fmt.Sprintf("%16X image base\n", me.FImageBase()) +
		fmt.Sprintf("%16X section alignment\n", me.FSectionAlignment()) +
		fmt.Sprintf("%16X file alignment\n", me.FFileAlignment()) +
		fmt.Sprintf("%16X operating system version\n", me.FMajorOperatingSystemVersion()) +
		fmt.Sprintf("%16X image version\n", me.FMajorImageVersion()) +
		fmt.Sprintf("%16X subsystem version\n", me.FMajorSubsystemVersion()) +
		fmt.Sprintf("%16X Win32 version\n", me.FWin32VersionValue()) +
		fmt.Sprintf("%16X size of image\n", me.FSizeOfImage()) +
		fmt.Sprintf("%16X size of headers\n", me.FSizeOfHeaders()) +
		fmt.Sprintf("%16X checksum\n", me.FCheckSum()) +
		fmt.Sprintf("%16X subsystem (Windows GUI)\n", me.FSubsystem()) +
		fmt.Sprintf("%16X DLL characteristics\n", me.FDllCharacteristics()) +
		fmt.Sprintf("%16X size of stack reserve\n", me.FSizeOfStackReserve()) +
		fmt.Sprintf("%16X size of stack commit\n", me.FSizeOfStackCommit()) +
		fmt.Sprintf("%16X size of heap header\n", me.FSizeOfHeaders()) +
		fmt.Sprintf("%16X size of heap reserve\n", me.FSizeOfHeapReserve()) +
		fmt.Sprintf("%16X size of heap commit\n", me.FSizeOfHeapCommit()) +
		fmt.Sprintf("%16X loader flags\n", me.FLoaderFlags()) +
		fmt.Sprintf("%16X number of directories\n", len(me.FDataDirectory()))
}

func (me *EXPORT_DIRECTORY_Info) String() string {
	var ret string
	ret = "EXPORT_DIRECTORY_Info VALUES\n" +
		fmt.Sprintf("%16s Module Name\n", me.ModuleName)

	for i := range me.FuncNames {
		ret += fmt.Sprintf("%16s Function %2d name, Ordinals No. %d\n", me.FuncNames[i], i, me.Ordinals[i])
	}

	return ret
}

func (me *IMAGE_EXPORT_DIRECTORY) String() string {
	return "导出表 VALUES\n" +
		fmt.Sprintf("%16X 保留，恒为0\n", me.Characteristics) +
		fmt.Sprintf("%16X time date stamp %s\n", me.TimeDateStamp, time.Unix(int64(me.TimeDateStamp), 0).Format("2006-01-02 15:04:05")) +
		fmt.Sprintf("%16X 主版本号\n", me.MajorVersion) +
		fmt.Sprintf("%16X 子版本号\n", me.MinorVersion) +
		fmt.Sprintf("%16X 指向模块名（导出表所在模块的名称）的ASCII字符的RVA\n", me.Name) +
		fmt.Sprintf("%16X 导出表用于输出导出函数序号值的基数: 导出函数序号 = 函数入口地址数组下标索引值 + 基数\n", me.Base) +
		fmt.Sprintf("%16X 导出函数入口地址表的成员个数\n", me.NumberOfFunctions) +
		fmt.Sprintf("%16X 以函数名称导出的成员个数\n", me.NumberOfNames) +
		fmt.Sprintf("%16X 函数入口地址表的相对虚拟地址(RVA)，每一个非0的项都对应一个被导出的函数名称或导出序号(序号+基数等于导出函数序号)\n", me.AddressOfFunctions) +
		fmt.Sprintf("%16X 函数名称表的相对虚拟地址(RVA),存储着指向导出函数名称的ASCII字符的RVA\n", me.AddressOfNames) +
		fmt.Sprintf("%16X 存储着函数入口地址表的数组下标索引值(序号表)，跟导出函数名称表的成员顺序对应\n", me.AddressOfNameOrdinals)
}

/* https://blog.csdn.net/qq_30145355/article/details/78848942
转换一个内存地址RVA到一个文件的偏移是多少，怎么实现？
1、把PE文件中的所有区段信息全部读取出来
2、RVA属于哪一个区段
3、RVA - RVA（用需要查找的RVA减去整个区段头的RVA，
*/

// https://bbs.pediy.com/thread-221766.htm
// 相对虚拟地址(RVA)与文件偏移地址转换(FOA)
func (me *PE) VA2FOA(virtualAddress DWORD) DWORD {
	//	for i := range me.Sections {
	//		if virtualAddress > me.Sections[i].VirtualAddress && virtualAddress < (me.Sections[i].VirtualAddress+me.Sections[i].SizeOfRawData) {
	//			/*
	//				如果此时我的RVA值为1014，那么文件偏移为多少？
	//				1、.text段的RAV(VirtualOffset)为1000，大小为4B48字节，明显我们的1014属于这个范围。对于的.text段的文件偏移(RawOffset)为400.
	//				2、（1014-1000）+ 400 = 414，这就是我们的文件偏移值
	//			*/
	//			return virtualAddress - me.Sections[i].VirtualAddress + me.Sections[i].PointerToRawData
	//		}
	//	}

	// Sections的VirtualAddress已经是升序排列的，可以使用二分法查找
	low, high := 0, len(me.Sections)-1
	end := high
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if me.Sections[mid].VirtualAddress == virtualAddress {
			break
		} else if me.Sections[mid].VirtualAddress < virtualAddress {
			if mid == end {
				break
			} else {
				// 小于的时候还要保证mid+1是大于它的
				if me.Sections[mid+1].VirtualAddress > virtualAddress {
					break
				}
			}
			// mid < virtualAddress，但是mid+1不大于，所以继续向后面找
			low = mid + 1
		} else {
			// 剩下的情况是mid > virtualAddress，所以继续向前面找
			high = mid - 1
		}
	}

	if low > high {
		return 0xffffffff
	} else {
		return virtualAddress - me.Sections[mid].VirtualAddress + me.Sections[mid].PointerToRawData
	}
}

func byte2struct(b []byte, pStruct interface{}) error {
	buf := bytes.NewBuffer(b)
	return binary.Read(buf, binary.LittleEndian, pStruct)
}

// 从f中读取byte，直到第1个0x0
func readByteEndOnZero(f *os.File) (ret []byte, err error) {
	var b []byte = make([]byte, 1)
	b[0] = 1

	for b[0] != 0x0 {
		if _, err = f.Read(b); err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			return
		}
		ret = append(ret, b[0])
	}
	return
}
