REM 编译32位的dll

set GOARCH=386
set GOROOT=C:\go386

set "str=%path%"
set path=C:\mingw32\mingw_tools\MinGW\bin;%str%

C:\go386\bin\go.exe build -v -x -buildmode=c-archive -o go.a
gcc go.def go.a -shared -lwinmm -lWs2_32 -o go.dll -Wl,--out-implib,go.lib

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64