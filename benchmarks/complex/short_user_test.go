package complex

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/complex"
	"testing"
)

func BenchmarkEmptyStateEncodingJSON(b *testing.B) {
	source := complex.EasyShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkEmptyStateEasyJSON(b *testing.B) {
	source := complex.EasyShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkEmptyStatePerezaJSON(b *testing.B) {
	source := complex.PerezaShortUser{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

