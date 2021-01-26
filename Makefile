get:
	go get github.com/zero-yy/export-config

exportconfig:
	export-config --config=default.toml

build: exportconfig
	go build -o gendata ./test/go/gen/conf/cmd/gendata/main.go

gendata: build
	./gendata --config=default.toml

init-mod:
	go mod init "github.com/zero-yy/export-config-test"

all: gendata

allall: init-mod get all


.PHONY: get exportconfig build gendata init-mod all allall