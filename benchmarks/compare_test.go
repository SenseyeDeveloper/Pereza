package benchmarks

import (
	"encoding/json"
	"github.com/senseyedeveloper/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodingJSON(t *testing.T) {
	dataProvider := map[bool]string{
		false: fixtures.ExpectBoolStateFalse,
		true:  fixtures.ExpectBoolStateTrue,
	}

	for state, expect := range dataProvider {
		source := fixtures.BoolState{
			State: state,
		}

		actual, err := json.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkEncodingJSON(b *testing.B) {
	source := fixtures.BoolState{
		State: true,
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(source)
	}
}
