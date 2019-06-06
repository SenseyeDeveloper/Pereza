PKG=github.com/senseyedeveloper/pereza

clean:
	rm -rf .root
	rm -rf fixtures/*_easyjson.go

build:
	go build -i -o .root/bin/pereza $(PKG)/pereza

generate: build
	.root/bin/pereza

easyjson:
	easyjson ./fixtures/empty_state.go
	easyjson ./fixtures/bool_state.go
	easyjson ./fixtures/int_state.go
	easyjson ./fixtures/string_state.go

test: easyjson generate
	go test ./benchmarks/... -v -bench=. -benchmem

all: test

.PHONY: test easyjson generate build clean