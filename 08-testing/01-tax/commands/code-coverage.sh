# GENERATE CODE COVERAGE
go test -coverprofile=coverage.out

# SEE CODE COVERAGE IN HTML
go tool cover -html=coverage.out