package boolean

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/boolean"
	"testing"
)

func BenchmarkHexaBoolStateEncodingJSON(b *testing.B) {
	source := boolean.HexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkHexaBoolStateEasyJSON(b *testing.B) {
	source := boolean.EasyHexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkHexaBoolStatePerezaJSON(b *testing.B) {
	source := boolean.PerezaHexaBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
