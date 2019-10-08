package pe

// 32
type IMAGE_OPTIONAL_HEADER struct {
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

// 64
// type IMAGE_OPTIONAL_HEADER struct {
// 	Magic                       WORD
// 	MajorLinkerVersion          byte
// 	MinorLinkerVersion          byte
// 	SizeOfCode                  DWORD
// 	SizeOfInitializedData       DWORD
// 	SizeOfUninitializedData     DWORD
// 	AddressOfEntryPoint         DWORD
// 	BaseOfCode                  DWORD
// 	ImageBase                   ULONGLONG
// 	SectionAlignment            DWORD
// 	FileAlignment               DWORD
// 	MajorOperatingSystemVersion WORD
// 	MinorOperatingSystemVersion WORD
// 	MajorImageVersion           WORD
// 	MinorImageVersion           WORD
// 	MajorSubsystemVersion       WORD
// 	MinorSubsystemVersion       WORD
// 	Win32VersionValue           DWORD
// 	SizeOfImage                 DWORD
// 	SizeOfHeaders               DWORD
// 	CheckSum                    DWORD
// 	Subsystem                   WORD
// 	DllCharacteristics          WORD
// 	SizeOfStackReserve          ULONGLONG
// 	SizeOfStackCommit           ULONGLONG
// 	SizeOfHeapReserve           ULONGLONG
// 	SizeOfHeapCommit            ULONGLONG
// 	LoaderFlags                 DWORD
// 	NumberOfRvaAndSizes         DWORD
// 	DataDirectory               [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]IMAGE_DATA_DIRECTORY
// }
