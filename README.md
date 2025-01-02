# go-testing-playground

## Useful built in CLI feature for testing 

### Run tests

`go test .`

### Run verbose tests

`go test -v .`

### Run one specific test 

`go test -run <TEST FUNC NAME>`

### Check the testing coverage: 

`go test -cover .`

### Build coverage.out and check it with a html ui

```go
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```