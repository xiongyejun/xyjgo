package pe

// 包装IMAGE_OPTIONAL_HEADER的interface
type IMAGE_OPTIONAL_HEADER struct {
	I_IMAGE_OPTIONAL_HEADER
}

// 用接口兼容32位和64位的结构
type I_IMAGE_OPTIONAL_HEADER interface {
	FMagic() WORD
	FMajorLinkerVersion() byte
	FMinorLinkerVersion() byte
	FSizeOfCode() DWORD
	FSizeOfInitializedData() DWORD
	FSizeOfUninitializedData() DWORD
	FAddressOfEntryPoint() DWORD
	FBaseOfCode() DWORD
	FBaseOfData() DWORD
	FImageBase() ULONGLONG
	FSectionAlignment() DWORD
	FFileAlignment() DWORD
	FMajorOperatingSystemVersion() WORD
	FMinorOperatingSystemVersion() WORD
	FMajorImageVersion() WORD
	FMinorImageVersion() WORD
	FMajorSubsystemVersion() WORD
	FMinorSubsystemVersion() WORD
	FWin32VersionValue() DWORD
	FSizeOfImage() DWORD
	FSizeOfHeaders() DWORD
	FCheckSum() DWORD
	FSubsystem() WORD
	FDllCharacteristics() WORD
	FSizeOfStackReserve() ULONGLONG
	FSizeOfStackCommit() ULONGLONG
	FSizeOfHeapReserve() ULONGLONG
	FSizeOfHeapCommit() ULONGLONG
	FLoaderFlags() DWORD
	FNumberOfRvaAndSizes() DWORD
	FDataDirectory() [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY
}

// 32
type IMAGE_OPTIONAL_HEADER32 struct {
	Magic                       WORD
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	BaseOfData                  DWORD
	ImageBase                   DWORD
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          DWORD
	SizeOfStackCommit           DWORD
	SizeOfHeapReserve           DWORD
	SizeOfHeapCommit            DWORD
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY
}

func (me IMAGE_OPTIONAL_HEADER32) FMagic() WORD {
	return me.Magic
}
func (me IMAGE_OPTIONAL_HEADER32) FMajorLinkerVersion() byte {
	return me.MajorLinkerVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMinorLinkerVersion() byte {
	return me.MinorLinkerVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfCode() DWORD {
	return me.SizeOfCode
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfInitializedData() DWORD {
	return me.SizeOfInitializedData
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfUninitializedData() DWORD {
	return me.SizeOfUninitializedData
}
func (me IMAGE_OPTIONAL_HEADER32) FAddressOfEntryPoint() DWORD {
	return me.AddressOfEntryPoint
}
func (me IMAGE_OPTIONAL_HEADER32) FBaseOfCode() DWORD {
	return me.BaseOfCode
}
func (me IMAGE_OPTIONAL_HEADER32) FBaseOfData() DWORD {
	return me.BaseOfData
}
func (me IMAGE_OPTIONAL_HEADER32) FImageBase() ULONGLONG {
	return ULONGLONG(me.ImageBase)
}
func (me IMAGE_OPTIONAL_HEADER32) FSectionAlignment() DWORD {
	return me.SectionAlignment
}
func (me IMAGE_OPTIONAL_HEADER32) FFileAlignment() DWORD {
	return me.FileAlignment
}
func (me IMAGE_OPTIONAL_HEADER32) FMajorOperatingSystemVersion() WORD {
	return me.MajorOperatingSystemVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMinorOperatingSystemVersion() WORD {
	return me.MinorOperatingSystemVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMajorImageVersion() WORD {
	return me.MajorImageVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMinorImageVersion() WORD {
	return me.MinorImageVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMajorSubsystemVersion() WORD {
	return me.MajorSubsystemVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FMinorSubsystemVersion() WORD {
	return me.MinorSubsystemVersion
}
func (me IMAGE_OPTIONAL_HEADER32) FWin32VersionValue() DWORD {
	return me.Win32VersionValue
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfImage() DWORD {
	return me.SizeOfImage
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfHeaders() DWORD {
	return me.SizeOfHeaders
}
func (me IMAGE_OPTIONAL_HEADER32) FCheckSum() DWORD {
	return me.CheckSum
}
func (me IMAGE_OPTIONAL_HEADER32) FSubsystem() WORD {
	return me.Subsystem
}
func (me IMAGE_OPTIONAL_HEADER32) FDllCharacteristics() WORD {
	return me.DllCharacteristics
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfStackReserve() ULONGLONG {
	return ULONGLONG(me.SizeOfStackReserve)
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfStackCommit() ULONGLONG {
	return ULONGLONG(me.SizeOfStackCommit)
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfHeapReserve() ULONGLONG {
	return ULONGLONG(me.SizeOfHeapReserve)
}
func (me IMAGE_OPTIONAL_HEADER32) FSizeOfHeapCommit() ULONGLONG {
	return ULONGLONG(me.SizeOfHeapCommit)
}
func (me IMAGE_OPTIONAL_HEADER32) FLoaderFlags() DWORD {
	return me.LoaderFlags
}
func (me IMAGE_OPTIONAL_HEADER32) FNumberOfRvaAndSizes() DWORD {
	return me.NumberOfRvaAndSizes
}
func (me IMAGE_OPTIONAL_HEADER32) FDataDirectory() [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY {
	return me.DataDirectory
}

// 64
type IMAGE_OPTIONAL_HEADER64 struct {
	Magic                       WORD
	MajorLinkerVersion          byte
	MinorLinkerVersion          byte
	SizeOfCode                  DWORD
	SizeOfInitializedData       DWORD
	SizeOfUninitializedData     DWORD
	AddressOfEntryPoint         DWORD
	BaseOfCode                  DWORD
	ImageBase                   ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion WORD
	MinorOperatingSystemVersion WORD
	MajorImageVersion           WORD
	MinorImageVersion           WORD
	MajorSubsystemVersion       WORD
	MinorSubsystemVersion       WORD
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   WORD
	DllCharacteristics          WORD
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY
}

func (me IMAGE_OPTIONAL_HEADER64) FMagic() WORD {
	return me.Magic
}
func (me IMAGE_OPTIONAL_HEADER64) FMajorLinkerVersion() byte {
	return me.MajorLinkerVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMinorLinkerVersion() byte {
	return me.MinorLinkerVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfCode() DWORD {
	return me.SizeOfCode
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfInitializedData() DWORD {
	return me.SizeOfInitializedData
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfUninitializedData() DWORD {
	return me.SizeOfUninitializedData
}
func (me IMAGE_OPTIONAL_HEADER64) FAddressOfEntryPoint() DWORD {
	return me.AddressOfEntryPoint
}
func (me IMAGE_OPTIONAL_HEADER64) FBaseOfCode() DWORD {
	return me.BaseOfCode
}
func (me IMAGE_OPTIONAL_HEADER64) FBaseOfData() DWORD {
	//	return me.BaseOfData
	return 0
}
func (me IMAGE_OPTIONAL_HEADER64) FImageBase() ULONGLONG {
	return me.ImageBase
}
func (me IMAGE_OPTIONAL_HEADER64) FSectionAlignment() DWORD {
	return me.SectionAlignment
}
func (me IMAGE_OPTIONAL_HEADER64) FFileAlignment() DWORD {
	return me.FileAlignment
}
func (me IMAGE_OPTIONAL_HEADER64) FMajorOperatingSystemVersion() WORD {
	return me.MajorOperatingSystemVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMinorOperatingSystemVersion() WORD {
	return me.MinorOperatingSystemVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMajorImageVersion() WORD {
	return me.MajorImageVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMinorImageVersion() WORD {
	return me.MinorImageVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMajorSubsystemVersion() WORD {
	return me.MajorSubsystemVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FMinorSubsystemVersion() WORD {
	return me.MinorSubsystemVersion
}
func (me IMAGE_OPTIONAL_HEADER64) FWin32VersionValue() DWORD {
	return me.Win32VersionValue
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfImage() DWORD {
	return me.SizeOfImage
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfHeaders() DWORD {
	return me.SizeOfHeaders
}
func (me IMAGE_OPTIONAL_HEADER64) FCheckSum() DWORD {
	return me.CheckSum
}
func (me IMAGE_OPTIONAL_HEADER64) FSubsystem() WORD {
	return me.Subsystem
}
func (me IMAGE_OPTIONAL_HEADER64) FDllCharacteristics() WORD {
	return me.DllCharacteristics
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfStackReserve() ULONGLONG {
	return me.SizeOfStackReserve
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfStackCommit() ULONGLONG {
	return me.SizeOfStackCommit
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfHeapReserve() ULONGLONG {
	return me.SizeOfHeapReserve
}
func (me IMAGE_OPTIONAL_HEADER64) FSizeOfHeapCommit() ULONGLONG {
	return me.SizeOfHeapCommit
}
func (me IMAGE_OPTIONAL_HEADER64) FLoaderFlags() DWORD {
	return me.LoaderFlags
}
func (me IMAGE_OPTIONAL_HEADER64) FNumberOfRvaAndSizes() DWORD {
	return me.NumberOfRvaAndSizes
}
func (me IMAGE_OPTIONAL_HEADER64) FDataDirectory() [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY {
	return me.DataDirectory
}
