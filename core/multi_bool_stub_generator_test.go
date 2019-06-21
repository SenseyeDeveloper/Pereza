package core

import "testing"

func BenchmarkMultiBoolStubGenerator_Generate(b *testing.B) {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	data := alphabet[:16]

	for i := 0; i < b.N; i++ {
		NewMultiBoolStubGenerator(data, data).Generate()
	}
}
