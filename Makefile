BINARY_NAME=netexecgo

all: build

build:
	@go build -o ${BINARY_NAME} main.go

run:
	@./${BINARY_NAME}

clean:
	@go clean
	@rm -f ${BINARY_NAME}

test:
	@go test -cover

.PHONY: all build run clean test