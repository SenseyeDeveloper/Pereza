package benchmarks

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"testing"
)

func BenchmarkAlphabetBoolStateEncodingJSON(b *testing.B) {
	source := fixtures.AlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkAlphabetBoolStateEasyJSON(b *testing.B) {
	source := fixtures.EasyAlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkAlphabetBoolStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaAlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
