package json

import (
	"encoding/json"
	"github.com/gopereza/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	stringStateDataProvider = map[string]string{
		"":    fixtures.ExpectStringStateEmpty,
		"abc": fixtures.ExpectStringStateSmall,
	}
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
	longString = alphabet + alphabet + alphabet + alphabet + alphabet
)

func TestStringStateEncodingJSON(t *testing.T) {
	for state, expect := range stringStateDataProvider {
		source := fixtures.StringState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestStringStateEasyJSON(t *testing.T) {
	for state, expect := range stringStateDataProvider {
		source := fixtures.EasyStringState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestStringStatePereza(t *testing.T) {
	for state, expect := range stringStateDataProvider {
		source := fixtures.PerezaStringState{
			State: state,
		}

		actual, err := source.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkStringStateEncodingJSON(b *testing.B) {
	source := fixtures.StringState{
		State: longString,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}

func BenchmarkStringStateEasyJSON(b *testing.B) {
	source := fixtures.EasyStringState{
		State: longString,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}

func BenchmarkStringStatePerezaJSON(b *testing.B) {
	source := fixtures.PerezaStringState{
		State: longString,
	}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalJSON()
	}
}
