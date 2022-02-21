BINDIR := bin
BIN := ${BINDIR}/anagrambler

.PHONY: all build clean
.PHONY: fmt lint test bench

all: build fmt lint test

${BINDIR}:
	mkdir -p ${BINDIR}

${BIN}: build

build: ${BINDIR}
	go build -v -o ${BIN} cmd/anagrambler/main.go

clean: ${BINDIR}
	rm -f ${BIN}

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

bench:
	go test -run=none -bench=. ./...
