REM 编译32位的dll

set GOARCH=386
set GOROOT=C:\go386

set "str=%path%"
set path=C:\mingw32\MinGW\bin;%str%

C:\go386\bin\go.exe build -v -x -buildmode=c-archive -o c\go.a
gcc.exe c\stdcall.c c\go.def c\go.a -shared -lwinmm -lWs2_32 -o go.dll -Wl,--enable-stdcall-fixup,--out-implib,go.lib

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64