PKG=github.com/gopereza/pereza

all: test

dep:
	dep ensure

.root/src/$(PKG):
	mkdir -p $@
	for i in $$PWD/* ; do ln -s $$i $@/`basename $$i` ; done

root: .root/src/$(PKG)

clean:
	rm -rf .root
	rm -rf pregen
	rm -rf fixtures/json/pregen
	rm -rf benchmarks/json/pregen
	rm -rf fixtures/*_easyjson.go
	rm -rf fixtures/*_perezajson.go
	rm -rf fixtures/boolean/*_easyjson.go
	rm -rf fixtures/boolean/*_perezajson.go

pregen-build:
	go build -i -o .root/bin/pregenref $(PKG)/pereza/pregenerator/reflect
	go build -i -o .root/bin/pregentest $(PKG)/pereza/pregenerator/test

pregen: pregen-build
	mkdir -p pregen
	.root/bin/pregenref > ./pregen/reflect_int_size.go
	go fmt ./pregen/...

	mkdir -p ./fixtures/json/pregen
	mkdir -p ./benchmarks/json/pregen
	.root/bin/pregentest
	go fmt ./fixtures/json/pregen/...
	go fmt ./benchmarks/json/pregen...

build: pregen
	go build -i -o .root/bin/pereza $(PKG)/pereza/generator

perezajson: build
	.root/bin/pereza ./fixtures/empty_state.go \
        ./fixtures/boolean/bool_state.go \
        ./fixtures/boolean/double_bool_state.go \
        ./fixtures/boolean/octo_bool_state.go \
        ./fixtures/boolean/hexa_bool_state.go \
        ./fixtures/boolean/alphabet_bool_state.go \
        ./fixtures/string_state.go \
        ./fixtures/json/pregen/int_state.go \
        ./fixtures/json/pregen/int8_state.go \
        ./fixtures/json/pregen/int16_state.go \
        ./fixtures/json/pregen/int32_state.go \
        ./fixtures/json/pregen/int64_state.go \
        ./fixtures/json/pregen/uint_state.go \
        ./fixtures/json/pregen/uint8_state.go \
        ./fixtures/json/pregen/uint16_state.go \
        ./fixtures/json/pregen/uint32_state.go \
        ./fixtures/json/pregen/uint64_state.go \
        ./fixtures/complex/short_user.go

easyjson:
	easyjson ./fixtures \
        ./fixtures/json/pregen \
        ./fixtures/boolean \
        ./fixtures/complex

generate: root perezajson easyjson

test: generate dep
	go test ./benchmarks/... -v -bench=. -benchmem

dev-hexa:
	.root/bin/pereza ./fixtures/boolean/hexa_bool_state.go
	go test ./benchmarks/boolean/... -v -run=HexaBool -bench=HexaBool -benchmem

dev-alphabet:
	.root/bin/pereza ./fixtures/boolean/alphabet_bool_state.go
	go test ./benchmarks/boolean/... -v -run=AlphabetBool -bench=AlphabetBool -benchmem

dev-user:
	.root/bin/pereza ./fixtures/complex/short_user.go
	go test ./benchmarks/complex/... -v -run=ShortUser -bench=ShortUser -benchmem

fmt:
	go fmt ./pregen/... ./benchmarks/... ./fixtures/... ./bootstrap/... ./pereza/... ./core/... ./gen/...

.PHONY: dep pregen-build pregen test generate easyjson perezajson build clean fmt development