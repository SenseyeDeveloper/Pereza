package benchmarks

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	doubleBoolStateDataProvider = map[bool]map[bool][]byte{
		true: {
			true:  []byte(`{"active":true,"approve":true}`),
			false: []byte(`{"active":true,"approve":false}`),
		},
		false: {
			true:  []byte(`{"active":false,"approve":true}`),
			false: []byte(`{"active":false,"approve":false}`),
		},
	}
)

func TestDoubleBoolStateEncodingJSON(t *testing.T) {
	for active, approveExpectMap := range doubleBoolStateDataProvider {
		for approve, expect := range approveExpectMap {
			source := fixtures.DoubleBoolState{
				Active:  active,
				Approve: approve,
			}

			actual, err := json.Marshal(source)
			assert.NoError(t, err)
			assert.Equal(t, expect, actual)
		}
	}
}

func TestDoubleBoolStateEasyJSON(t *testing.T) {
	for active, approveExpectMap := range doubleBoolStateDataProvider {
		for approve, expect := range approveExpectMap {
			source := fixtures.EasyDoubleBoolState{
				Active:  active,
				Approve: approve,
			}

			actual, err := source.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, expect, actual)
		}
	}
}

func TestDoubleBoolStatePereza(t *testing.T) {
	t.Skip("todo")

	for active, approveExpectMap := range doubleBoolStateDataProvider {
		for approve, expect := range approveExpectMap {
			source := fixtures.PerezaDoubleBoolState{
				Active:  active,
				Approve: approve,
			}

			actual, err := source.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, expect, actual)
		}
	}
}

func BenchmarkDoubleBoolStateEncodingJSON(b *testing.B) {
	source := fixtures.DoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkDoubleBoolStateEasyJSON(b *testing.B) {
	source := fixtures.EasyDoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkDoubleBoolStatePerezaJSON(b *testing.B) {
	b.Skip("todo")

	source := fixtures.PerezaDoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
