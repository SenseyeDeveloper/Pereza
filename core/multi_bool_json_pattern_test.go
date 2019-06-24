package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	multiBoolJSONPatternData = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
)

func TestMultiBoolJSONResultGenerator_Generate(t *testing.T) {
	generator := NewMultiBoolJSONResultGenerator(multiBoolJSONPatternData)

	assert.Equal(
		t,
		multiBoolJSONResultHeader+`"a":true,"b":true,"c":true,"d":true,"e":true,"f":true,"g":true,"h":true`+multiBoolJSONResultFooter,
		string(generator.Generate([]bool{true, true, true, true, true, true, true, true})),
	)

	assert.Equal(
		t,
		multiBoolJSONResultHeader+`"a":false,"b":false,"c":false,"d":false,"e":false,"f":false,"g":false,"h":false`+multiBoolJSONResultFooter,
		string(generator.Generate([]bool{false, false, false, false, false, false, false, false})),
	)
}

func BenchmarkMultiBoolJSONResultGenerator_Generate(b *testing.B) {
	generator := NewMultiBoolJSONResultGenerator(multiBoolJSONPatternData)
	states := []bool{false, false, false, false, false, false, false, false}

	for i := 0; i < b.N; i++ {
		_ = generator.Generate(states)
	}
}
