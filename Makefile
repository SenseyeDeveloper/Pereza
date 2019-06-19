PKG=github.com/senseyedeveloper/pereza

clean:
	rm -rf .root
	rm -rf fixtures/*_easyjson.go
	rm -rf fixtures/*_perezajson.go

build:
	go build -i -o .root/bin/pereza $(PKG)/pereza/generator
	go build -i -o .root/bin/pregen $(PKG)/pereza/pregenerator

pregen: build
	.root/bin/pregen > ./core/reflect_int_size.go
	go fmt ./core/reflect_int_size.go

perezajson: build
	.root/bin/pereza ./fixtures/empty_state.go \
        ./fixtures/bool_state.go \
        ./fixtures/int_state.go \
        ./fixtures/uint_state.go \
        ./fixtures/string_state.go

easyjson:
	easyjson ./fixtures

generate: easyjson perezajson

test: generate
	go test ./benchmarks/... -v -bench=. -benchmem

all: test

fmt:
	go fmt ./benchmarks/... ./fixtures/... ./bootstrap/... ./pereza/... ./core/... ./gen/...

.PHONY: pregen test generate easyjson perezajson build clean