# Pereza
experimental json marshaler

##### Very simple structure
```golang
struct {
	State bool `json:"state"`
}
```
```text
BenchmarkEncodingJSON   	 5000000	       309 ns/op	     177 B/op	       2 allocs/op
BenchmarkEasyJSON       	20000000	       114 ns/op	     128 B/op	       1 allocs/op
BenchmarkPerezaJSON     	200000000	         6.30 ns/op	       0 B/op	       0 allocs/op
```
