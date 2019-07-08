REM ±àÒë32Î»µÄ

set GOARCH=386
set GOROOT=C:\go386

set "str=%path%"
set path=C:\mingw32\MinGW\bin;%str%


C:\go386\bin\go.exe build -o txt×ª»»32.exe %GOPATH%\src\github.com\xiongyejun\xyjgo\txtConv



set path = %str%
set GOROOT=C:\Go
set GOARCH=amd64


pause