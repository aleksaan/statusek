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

# Как тестировать
go test ./...

# Как добавить тэг
git tag v2.0
git push origin v2.0