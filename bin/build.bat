set GOOS=%1
set GOARCH=%2
set VERSION=%3
mkdir .\bin\%VERSION%\%GOOS%_%GOARCH%
go env -w GOOS=%GOOS%
go env -w GOARCH=%GOARCH%
go build -o .\bin\%VERSION%\%GOOS%_%GOARCH% .\...
7z a .\bin\%VERSION%\%VERSION%_%GOOS%_%GOARCH%_statusek.zip .\bin\%VERSION%\%GOOS%_%GOARCH% 