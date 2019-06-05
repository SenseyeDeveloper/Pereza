all: test

test: easyjson
	go test ./benchmarks/... -v -bench=. -benchmem

easyjson:
	easyjson ./fixtures/state.go
