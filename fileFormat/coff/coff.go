// COFF(Common Object File Format)
//		PE（Protable Executable）32位win可执行文件
//		PE的前身就是COFF
package coff

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"time"
)

// typedef struct _IMAGE_FILE_HEADER {
//   WORD Machine;
//   WORD NumberOfSections;
//   DWORD TimeDateStamp;
//   DWORD PointerToSymbolTable;
//   DWORD NumberOfSymbols;
//   WORD SizeOfOptionalHeader;
//   WORD Characteristics;
// } IMAGE_FILE_HEADER,*PIMAGE_FILE_HEADER;

type SHORT = int16
type WORD = uint16
type DWORD = uint32

// 映像（Image)
// PE文件在装载时被直接映射到进程的虚拟空间中运行，它是进程的虚拟空间映像。
// 所以PE可执行文件很多时候被叫做映像文件（Image File）
type IMAGE_FILE_HEADER struct {
	Machine              WORD
	NumberOfSections     WORD
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader WORD
	Characteristics      WORD
}

// typedef struct _IMAGE_SECTION_HEADER {
//   BYTE Name[IMAGE_SIZEOF_SHORT_NAME];
//   union {
// DWORD PhysicalAddress;
// DWORD VirtualSize;
//   } Misc;
//   DWORD VirtualAddress;
//   DWORD SizeOfRawData;
//   DWORD PointerToRawData;
//   DWORD PointerToRelocations;
//   DWORD PointerToLinenumbers;
//   WORD NumberOfRelocations;
//   WORD NumberOfLinenumbers;
//   DWORD Characteristics;
// } IMAGE_SECTION_HEADER,*PIMAGE_SECTION_HEADER;
const IMAGE_SIZEOF_SHORT_NAME = 8

type IMAGE_SECTION_HEADER struct {
	Name                 [IMAGE_SIZEOF_SHORT_NAME]byte
	Misc                 DWORD
	VirtualAddress       DWORD
	SizeOfRawData        DWORD
	PointerToRawData     DWORD
	PointerToRelocations DWORD
	PointerToLinenumbers DWORD
	NumberOfRelocations  WORD
	NumberOfLinenumbers  WORD
	Characteristics      DWORD
}

// typedef struct _IMAGE_SYMBOL {
//     union {
// BYTE ShortName[8];
// struct {
//  DWORD Short;
//  DWORD Long;
// } Name;
// DWORD LongName[2];
//     } N;
//     DWORD Value;
//     SHORT SectionNumber;
//     WORD Type;
//     BYTE StorageClass;
//     BYTE NumberOfAuxSymbols;
//   } IMAGE_SYMBOL;
//   typedef IMAGE_SYMBOL UNALIGNED *PIMAGE_SYMBOL;
type IMAGE_SYMBOL struct {
	ShortName          [8]byte
	Value              DWORD
	SectionNumber      SHORT
	Type               WORD
	StorageClass       byte
	NumberOfAuxSymbols byte
}

// objdump -h SimpleSection.o 查看木变文件的结构和内容
// .text	Code Section	代码段
// .data	Data Section	数据段	已初始化的全局变量和局部静态变量
// .bss		Block Started by Symbol		--		未初始化的全局变量和局部静态变量

type COFF struct {
	f        *os.File
	Header   *IMAGE_FILE_HEADER
	Sections []*IMAGE_SECTION_HEADER
	Symbols  []*IMAGE_SYMBOL
}

func init() {

}

// 解析文件
func (me *COFF) Parse(f *os.File) (err error) {
	me.f = f
	if err = me.readHeader(); err != nil {
		return
	}

	if err = me.readSections(); err != nil {
		return
	}

	// 设置文件读取的开始地址
	if _, err = me.f.Seek(int64(me.Header.PointerToSymbolTable), 0); err != nil {
		return
	}
	if err = me.readSymbol(); err != nil {
		return
	}
	return
}

// 读取文件头
func (me *COFF) readHeader() (err error) {
	me.Header = new(IMAGE_FILE_HEADER)
	headerSize := binary.Size(me.Header)
	var b []byte = make([]byte, headerSize)
	var n int
	n, err = me.f.Read(b)
	if err != nil {
		return
	}
	if n < headerSize {
		err = errors.New("未能读取到【IMAGE_FILE_HEADER】足够的字节数。")
		return
	}

	if err = byte2struct(b, me.Header); err != nil {
		return
	}
	return
}

// 读取sections
func (me *COFF) readSections() (err error) {
	me.Sections = make([]*IMAGE_SECTION_HEADER, me.Header.NumberOfSections)
	sectionSize := binary.Size(new(IMAGE_SECTION_HEADER))

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
		tmp := new(IMAGE_SECTION_HEADER)
		if err = byte2struct(b, tmp); err != nil {
			return
		}
		me.Sections[i] = tmp
	}

	return
}

// 读取Symbol
func (me *COFF) readSymbol() (err error) {
	me.Symbols = make([]*IMAGE_SYMBOL, me.Header.NumberOfSymbols)
	symbolSize := binary.Size(new(IMAGE_SYMBOL))

	// 逐个读取symbol
	for i := range me.Symbols {
		var b []byte = make([]byte, symbolSize)
		var n int
		n, err = me.f.Read(b)
		if err != nil {
			return
		}
		if n < symbolSize {
			err = errors.New("未能读取到【IMAGE_SYMBOL】足够的字节数。")
			return
		}
		tmp := new(IMAGE_SYMBOL)
		if err = byte2struct(b, tmp); err != nil {
			return
		}
		me.Symbols[i] = tmp
	}

	return
}

func (me *IMAGE_FILE_HEADER) GetCoffHeaderPrintStr() string {
	return "FILE HEADER VALUES\n" +
		fmt.Sprintf("%16X machine (x86)\n", me.Machine) +
		fmt.Sprintf("%16X number of sections\n", me.NumberOfSections) +
		fmt.Sprintf("%16X time date stamp %s\n", me.TimeDateStamp, time.Unix(int64(me.TimeDateStamp), 0).Format("2006-01-02 15:04:05")) +
		fmt.Sprintf("%16X file pointer to symbol table\n", me.PointerToSymbolTable) +
		fmt.Sprintf("%16X number of symbols\n", me.NumberOfSymbols) +
		fmt.Sprintf("%16X size of optional header\n", me.SizeOfOptionalHeader) +
		fmt.Sprintf("%16X characteristics\n", me.Characteristics)
}
func byte2struct(b []byte, pStruct interface{}) error {
	buf := bytes.NewBuffer(b)
	return binary.Read(buf, binary.LittleEndian, pStruct)
}
