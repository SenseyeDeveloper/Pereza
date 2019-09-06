package boolean

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/boolean"
	"testing"
)

func BenchmarkAlphabetBoolStateEncodingJSON(b *testing.B) {
	source := boolean.AlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkAlphabetBoolStateEasyJSON(b *testing.B) {
	source := boolean.EasyAlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkAlphabetBoolStatePerezaJSON(b *testing.B) {
	source := boolean.PerezaAlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
