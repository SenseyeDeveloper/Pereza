package benchmarks

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"testing"
)

func BenchmarkOctoBoolStateEncodingJSON(b *testing.B) {
	source := fixtures.OctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkOctoBoolStateEasyJSON(b *testing.B) {
	source := fixtures.EasyOctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkOctoBoolStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaOctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
