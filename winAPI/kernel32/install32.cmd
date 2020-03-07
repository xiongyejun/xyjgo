REM 编译32位

set GOARCH=386
set GOROOT=C:\go386

set "str=%path%"
set path=C:\mingw32\MinGW\bin;%str%

C:\go386\bin\go.exe install

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64

pause