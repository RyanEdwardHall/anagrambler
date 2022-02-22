BINDIR := bin
BIN := ${BINDIR}/anagrambler

.PHONY: build clean fmt lint test bench

build: | ${BINDIR}
	go build -v -o ${BIN} cmd/anagrambler/main.go

clean: | ${BINDIR}
	rm -f ${BIN}

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

bench:
	go test -run=none -bench=. ./...

${BINDIR}:
	mkdir -p ${BINDIR}
