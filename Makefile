PKG=github.com/senseyedeveloper/pereza

clean:
	rm -rf .root
	rm -rf pregen
	rm -rf fixtures/*_easyjson.go
	rm -rf fixtures/*_perezajson.go

pregen-build:
	go build -i -o .root/bin/pregenref $(PKG)/pereza/pregenerator/reflect

pregen: pregen-build
	mkdir -p pregen
	.root/bin/pregenref > ./pregen/reflect_int_size.go
	go fmt ./pregen/reflect_int_size.go

build: pregen
	go build -i -o .root/bin/pereza $(PKG)/pereza/generator

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
	go fmt ./pregen/... ./benchmarks/... ./fixtures/... ./bootstrap/... ./pereza/... ./core/... ./gen/...

.PHONY: pregen-build pregen test generate easyjson perezajson build clean