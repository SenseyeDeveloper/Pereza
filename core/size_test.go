package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	stubTypes     = []string{"a", "ab", "abc", "abcd"}
	stubLongTypes = []string{"a", "ab", "abc", "abcd", "abcde"}
)

func TestStringSliceSize(t *testing.T) {
	assert.Equal(t, 10, stringSliceSize(stubTypes))
}

func BenchmarkStringSliceSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stringSliceSize(stubTypes)
	}
}
