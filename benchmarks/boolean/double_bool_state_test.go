package boolean

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures/boolean"
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
			source := boolean.DoubleBoolState{
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
			source := boolean.EasyDoubleBoolState{
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
	for active, approveExpectMap := range doubleBoolStateDataProvider {
		for approve, expect := range approveExpectMap {
			source := boolean.PerezaDoubleBoolState{
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
	source := boolean.DoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkDoubleBoolStateEasyJSON(b *testing.B) {
	source := boolean.EasyDoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkDoubleBoolStatePerezaJSON(b *testing.B) {
	source := boolean.PerezaDoubleBoolState{
		Active:  true,
		Approve: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
