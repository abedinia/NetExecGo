BINARY_NAME=netexecgo

all: build-linux build-mac build-windows

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}-linux main.go

build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME}-mac main.go

build-windows:
	@GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}-windows.exe main.go

run:
	@./${BINARY_NAME}

clean:
	@go clean
	@rm -f ${BINARY_NAME}-*

test:
	@go test -cover

.PHONY: all build-linux build-mac build-windows run clean test
