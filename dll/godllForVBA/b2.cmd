REM 编译32位的dll

set GOARCH=386
set GOROOT=C:\go386

set "str=%path%"
set path=C:\mingw32\MinGW\bin;%str%

gcc.exe c\stdcall.c c\go.def c\go.a -shared -lwinmm -lWs2_32 -o go32.dll -Wl,--enable-stdcall-fixup,--out-implib,go.lib

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64

copy go32.dll C:\Users\Administrator\Documents\VBAProject\godllForVBA32.dll

pause