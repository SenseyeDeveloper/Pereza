PKG=github.com/senseyedeveloper/pereza

all: test

dep:
	dep ensure

clean:
	rm -rf .root
	rm -rf pregen
	rm -rf fixtures/pregen
	rm -rf benchmarks/pregen
	rm -rf fixtures/*_easyjson.go
	rm -rf fixtures/*_perezajson.go

pregen-build:
	go build -i -o .root/bin/pregenref $(PKG)/pereza/pregenerator/reflect
	go build -i -o .root/bin/pregentest $(PKG)/pereza/pregenerator/test

pregen: pregen-build
	mkdir -p pregen
	.root/bin/pregenref > ./pregen/reflect_int_size.go
	go fmt ./pregen/...

	mkdir -p ./fixtures/pregen
	mkdir -p ./benchmarks/pregen
	.root/bin/pregentest
	go fmt ./fixtures/pregen/...

build: pregen
	go build -i -o .root/bin/pereza $(PKG)/pereza/generator

perezajson: build
	.root/bin/pereza ./fixtures/empty_state.go \
        ./fixtures/bool_state.go \
        ./fixtures/string_state.go \
        ./fixtures/pregen/int_state.go \
        ./fixtures/pregen/int8_state.go \
        ./fixtures/pregen/int16_state.go \
        ./fixtures/pregen/int32_state.go \
        ./fixtures/pregen/int64_state.go \
        ./fixtures/pregen/uint_state.go \
        ./fixtures/pregen/uint8_state.go \
        ./fixtures/pregen/uint16_state.go \
        ./fixtures/pregen/uint32_state.go \
        ./fixtures/pregen/uint64_state.go

easyjson:
	easyjson ./fixtures ./fixtures/pregen

generate: perezajson easyjson

test: generate dep
	go test ./benchmarks/... -v -bench=. -benchmem

fmt:
	go fmt ./pregen/... ./benchmarks/... ./fixtures/... ./bootstrap/... ./pereza/... ./core/... ./gen/...

.PHONY: dep pregen-build pregen test generate easyjson perezajson build clean