package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	multiBoolJSONPatternData = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
)

func TestMultiBoolJSONPattern(t *testing.T) {
	assert.Equal(t, `{"active":%t}`, MultiBoolJSONPattern([]string{"active"}))
	assert.Equal(t, `{"a":%t,"b":%t,"c":%t,"d":%t,"e":%t,"f":%t,"g":%t,"h":%t}`, MultiBoolJSONPattern(multiBoolJSONPatternData))
}

func BenchmarkMultiBoolJSONPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MultiBoolJSONPattern(multiBoolJSONPatternData)
	}
}
