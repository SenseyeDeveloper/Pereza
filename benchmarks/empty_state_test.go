package benchmarks

import (
	"encoding/json"
	"github.com/senseyedeveloper/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStateEncodingJSON(t *testing.T) {
	source := fixtures.EmptyState{}

	actual, err := json.Marshal(source)
	assert.NoError(t, err)
	assert.Equal(t, []byte(fixtures.ExpectEmptyState), actual)
}

func TestEmptyStateEasyJSON(t *testing.T) {
	source := fixtures.EasyEmptyState{}

	actual, err := source.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(fixtures.ExpectEmptyState), actual)
}

func TestEmptyStatePereza(t *testing.T) {
	source := fixtures.PerezaEmptyState{}

	actual, err := source.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(fixtures.ExpectEmptyState), actual)
}

func BenchmarkEmptyStateEncodingJSON(b *testing.B) {
	source := fixtures.EmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkEmptyStateEasyJSON(b *testing.B) {
	source := fixtures.EasyEmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkEmptyStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaEmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
