all: test

test: easyjson
	go test ./benchmarks/... -v -bench=. -benchmem

easyjson:
	easyjson ./fixtures/empty_state.go
	easyjson ./fixtures/bool_state.go
	easyjson ./fixtures/int_state.go
	easyjson ./fixtures/string_state.go
