PKG=github.com/senseyedeveloper/pereza

clean:
	rm -rf .root
	rm -rf fixtures/*_easyjson.go

build:
	go build -i -o .root/bin/pereza $(PKG)/pereza

generate: build
	.root/bin/pereza ./fixtures/empty_state.go
	.root/bin/pereza ./fixtures/bool_state.go
	.root/bin/pereza ./fixtures/int_state.go
	.root/bin/pereza ./fixtures/string_state.go

easyjson:
	easyjson ./fixtures/empty_state.go
	easyjson ./fixtures/bool_state.go
	easyjson ./fixtures/int_state.go
	easyjson ./fixtures/string_state.go

test: easyjson generate
	go test ./benchmarks/... -v -bench=. -benchmem

all: test

.PHONY: test easyjson generate build clean