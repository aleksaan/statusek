# get test coverage of the package

go test ./tests -cover -coverpkg ./database -coverprofile cover.out

# see report in html

go tool cover -html cover.out

# readme

go mod init ...
go clean -modcache
go get gorm.io/gorm  
go mod tidy

# Как делать бинарники

.\bin\build_all_platforms.bat
