# get test coverage of the package
go test ./tests -cover -coverpkg ./database -coverprofile cover.out

# see report in html
go tool cover -html cover.out