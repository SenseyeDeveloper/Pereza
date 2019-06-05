# Pereza
experimental json marshaler

##### Very simple structure
```golang
type BoolState struct {
	State bool `json:"state"`
}
```
```text
BenchmarkEncodingJSON   	 5000000	       311 ns/op	     177 B/op	       2 allocs/op
```
