build:
	@go build 

run:
	@go run main.go 

install:
	@go install 

test-verbose:
	@go test -v . 

test:
	@go test 