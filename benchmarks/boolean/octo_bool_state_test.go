package boolean

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/boolean"
	"testing"
)

func BenchmarkOctoBoolStateEncodingJSON(b *testing.B) {
	source := boolean.OctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkOctoBoolStateEasyJSON(b *testing.B) {
	source := boolean.EasyOctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkOctoBoolStatePerezaJSON(b *testing.B) {
	source := boolean.PerezaOctoBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
