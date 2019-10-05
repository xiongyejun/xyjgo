package pe

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"

	"github.com/xiongyejun/xyjgo/fileFormat/coff"
)

type SHORT = int16
type WORD = uint16
type DWORD = uint32
type ULONGLONG = uint64

func init() {

}

// typedef struct _IMAGE_DOS_HEADER {
//    E_magic WORD
//    E_cblp WORD
//    E_cp WORD
//    E_crlc WORD
//    E_cparhdr WORD
//    E_minalloc WORD
//    E_maxalloc WORD
//    E_ss WORD
//    E_sp WORD
//    E_csum WORD
//    E_ip WORD
//    E_cs WORD
//    E_lfarlc WORD
//    E_ovno WORD
//    E_res[4] WORD
//    E_oemid WORD
//    E_oeminfo WORD
//    E_res2[10] WORD
//    LONG E_lfanew WORD
//  } IMAGE_DOS_HEADER,*PIMAGE_DOS_HEADER WORD

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

// typedef struct _IMAGE_DATA_DIRECTORY {
//   DWORD VirtualAddress;
//   DWORD Size;
// } IMAGE_DATA_DIRECTORY,*PIMAGE_DATA_DIRECTORY;

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress DWORD
	Size           DWORD
}

const IMAGE_NUMBEROF_DIRECTORY_ENTRIES = 16

// #ifdef _WIN64
//     typedef IMAGE_OPTIONAL_HEADER64 IMAGE_OPTIONAL_HEADER DWORD
//     typedef PIMAGE_OPTIONAL_HEADER64 PIMAGE_OPTIONAL_HEADER DWORD
// #define IMAGE_SIZEOF_NT_OPTIONAL_HEADER IMAGE_SIZEOF_NT_OPTIONAL64_HEADER
// #define IMAGE_NT_OPTIONAL_HDR_MAGIC IMAGE_NT_OPTIONAL_HDR64_MAGIC
// #else  /* _WIN64 */
//     typedef IMAGE_OPTIONAL_HEADER32 IMAGE_OPTIONAL_HEADER DWORD
//     typedef PIMAGE_OPTIONAL_HEADER32 PIMAGE_OPTIONAL_HEADER DWORD
// #define IMAGE_SIZEOF_NT_OPTIONAL_HEADER IMAGE_SIZEOF_NT_OPTIONAL32_HEADER
// #define IMAGE_NT_OPTIONAL_HDR_MAGIC IMAGE_NT_OPTIONAL_HDR32_MAGIC
// #endif /* _WIN64 */

//     typedef struct _IMAGE_NT_HEADERS64 {
//       Signature DWORD
//       IMAGE_FILE_HEADER FileHeader DWORD
//       IMAGE_OPTIONAL_HEADER64 OptionalHeader DWORD
//     } IMAGE_NT_HEADERS64,*PIMAGE_NT_HEADERS64 DWORD

//     typedef struct _IMAGE_NT_HEADERS {
//       Signature DWORD
//       IMAGE_FILE_HEADER FileHeader DWORD
//       IMAGE_OPTIONAL_HEADER32 OptionalHeader DWORD
//     } IMAGE_NT_HEADERS32,*PIMAGE_NT_HEADERS32 DWORD

//     typedef struct _IMAGE_ROM_HEADERS {
//       IMAGE_FILE_HEADER FileHeader DWORD
//       IMAGE_ROM_OPTIONAL_HEADER OptionalHeader DWORD
//     } IMAGE_ROM_HEADERS,*PIMAGE_ROM_HEADERS DWORD
type IMAGE_NT_HEADERS struct {
	Signature      DWORD
	FileHeader     coff.IMAGE_FILE_HEADER
	OptionalHeader IMAGE_OPTIONAL_HEADER
}

type PE struct {
	f         *os.File
	DosHeader *IMAGE_DOS_HEADER
	NTHeader  *IMAGE_NT_HEADERS
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
func byte2struct(b []byte, pStruct interface{}) error {
	buf := bytes.NewBuffer(b)
	return binary.Read(buf, binary.LittleEndian, pStruct)
}
