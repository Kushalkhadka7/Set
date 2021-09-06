build:
	@go build 
	
install:
	@go install 

test-verbose:
	@go test -v . 

test:
	@go test 