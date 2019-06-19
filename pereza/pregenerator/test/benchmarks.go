package main

const (
	benchmarkOutputFilePattern = "%s_state_test.go"
)

type minmax struct {
	min, max, mathMin, mathMax string
}

func benchmarks(path string) {
	intBenchmarks(path)
	uintBenchmarks(path)
}
