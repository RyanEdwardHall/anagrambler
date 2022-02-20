BIN = bin
CMD := ${BIN}/anagrambler

.PHONY: all
all: fmt test

${BIN}:
	mkdir -p ${BIN}

${BIN}/%: ${BIN}
	go build -v -o $@ cmd/$(@F)/main.go

.PHONY: build
build: ${CMD}

.PHONY: rebuild
rebuild: clean build

.PHONY: clean
clean: ${BIN}
	rm -f ${CMD}

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...
