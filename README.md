# Pereza
experimental json marshaler

##### Very simple structures
```golang
struct {
	State bool `json:"state"`
}
```
```text
BenchmarkBoolStateEncodingJSON   	 5000000	       306 ns/op	     177 B/op	       2 allocs/op
BenchmarkBoolStateEasyJSON       	20000000	       110 ns/op	     128 B/op	       1 allocs/op
BenchmarkBoolStatePerezaJSON     	200000000	         6.39 ns/op	       0 B/op	       0 allocs/op
```

```golang
struct {
	State int `json:"state"`
}
```
```text
BenchmarkIntStateEncodingJSON    	 5000000	       331 ns/op	     184 B/op	       2 allocs/op
BenchmarkIntStateEasyJSON        	10000000	       129 ns/op	     128 B/op	       1 allocs/op
BenchmarkIntStatePerezaJSON      	30000000	        58.1 ns/op	      32 B/op	       1 allocs/op
```
