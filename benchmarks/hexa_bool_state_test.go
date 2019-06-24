package benchmarks

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"testing"
)

func BenchmarkHexaBoolStateEncodingJSON(b *testing.B) {
	source := fixtures.HexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkHexaBoolStateEasyJSON(b *testing.B) {
	source := fixtures.EasyHexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkHexaBoolStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaHexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
