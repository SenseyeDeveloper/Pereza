# Pereza [![Build Status](https://travis-ci.org/gopereza/pereza.svg?branch=master)](https://travis-ci.org/gopereza/pereza)[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/gopereza/pereza/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/gopereza/pereza/?branch=master)
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

```golang
struct {
	State string `json:"state"`
}
```
```text
BenchmarkStringStateEncodingJSON   	 2000000	       761 ns/op	     480 B/op	       3 allocs/op
BenchmarkStringStateEasyJSON       	 2000000	       719 ns/op	     720 B/op	       4 allocs/op
BenchmarkStringStatePerezaJSON     	20000000	        69.6 ns/op	     144 B/op	       1 allocs/op
```
