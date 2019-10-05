package pe

// typedef struct _IMAGE_OPTIONAL_HEADER {

//      WORD Magic DWORD
//      BYTE MajorLinkerVersion DWORD
//      BYTE MinorLinkerVersion DWORD
//      SizeOfCode DWORD
//      SizeOfInitializedData DWORD
//      SizeOfUninitializedData DWORD
//      AddressOfEntryPoint DWORD
//      BaseOfCode DWORD
//      BaseOfData DWORD
//      ImageBase DWORD
//      SectionAlignment DWORD
//      FileAlignment DWORD
//      WORD MajorOperatingSystemVersion DWORD
//      WORD MinorOperatingSystemVersion DWORD
//      WORD MajorImageVersion DWORD
//      WORD MinorImageVersion DWORD
//      WORD MajorSubsystemVersion DWORD
//      WORD MinorSubsystemVersion DWORD
//      Win32VersionValue DWORD
//      SizeOfImage DWORD
//      SizeOfHeaders DWORD
//      CheckSum DWORD
//      WORD Subsystem DWORD
//      WORD DllCharacteristics DWORD
//      SizeOfStackReserve DWORD
//      SizeOfStackCommit DWORD
//      SizeOfHeapReserve DWORD
//      SizeOfHeapCommit DWORD
//      LoaderFlags DWORD
//      NumberOfRvaAndSizes DWORD
//      IMAGE_DATA_DIRECTORY DataDirectory[IMAGE_NUMBEROF_DIRECTORY_ENTRIES] DWORD
//    } IMAGE_OPTIONAL_HEADER32,*PIMAGE_OPTIONAL_HEADER32 DWORD

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
