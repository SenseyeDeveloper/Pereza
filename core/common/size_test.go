package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	stubTypes     = []string{"a", "ab", "abc", "abcd"}
)

func TestStringSliceSize(t *testing.T) {
	assert.Equal(t, 10, StringSliceSize(stubTypes))
}

func BenchmarkStringSliceSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringSliceSize(stubTypes)
	}
}
