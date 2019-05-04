REM 编译32位的dll

set GOARCH=386
set GOROOT=C:\Users\Administrator\Downloads\go386

set "str=%path%"
set path=C:\mingw32\MinGW\bin;%str%

C:\Users\Administrator\Downloads\go386\bin\go.exe build -v -x -buildmode=c-archive -o go.a

set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64