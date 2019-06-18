PKG=github.com/senseyedeveloper/pereza

clean:
	rm -rf .root
	rm -rf fixtures/*_easyjson.go
	rm -rf fixtures/*_perezajson.go

build:
	go build -i -o .root/bin/pereza $(PKG)/pereza

perezajson: build
	.root/bin/pereza ./fixtures/empty_state.go \
        ./fixtures/bool_state.go \
        ./fixtures/int_state.go \
        ./fixtures/uint_state.go \
        ./fixtures/string_state.go

easyjson:
	easyjson ./fixtures/empty_state.go \
	    ./fixtures/bool_state.go \
	    ./fixtures/int_state.go \
	    ./fixtures/uint_state.go \
	    ./fixtures/string_state.go

generate: easyjson perezajson

test: generate
	go test ./benchmarks/... -v -bench=. -benchmem

all: test

fmt:
	go fmt ./benchmarks/... ./fixtures/... ./bootstrap/... ./pereza/... ./core/... ./gen/...

.PHONY: test generate easyjson perezajson build clean