BINARY_NAME=proxgo

all: build

build:
	@go build -o ${BINARY_NAME} main.go

run:
	@./${BINARY_NAME}

clean:
	@go clean
	@rm -f ${BINARY_NAME}

.PHONY: all build run clean