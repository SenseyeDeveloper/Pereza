package boolstub

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	hexaFieldNames = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
)

func TestFastConditionMap(t *testing.T) {
	actual := FastConditionMap([]string{"A", "B", "C"})

	assert.Equal(
		t,
		map[string][]byte{
			"A": []byte("if v.A {\n"),
			"B": []byte("if v.B {\n"),
			"C": []byte("if v.C {\n"),
		},
		actual,
	)
}

func BenchmarkOneBoolStubGenerator_Generate(b *testing.B) {
	fieldNames := []string{"A"}
	jsonNames := []string{"a"}

	for i := 0; i < b.N; i++ {
		generator := NewCombinatorBoolStubGenerator(fieldNames, jsonNames)

		_ = generator.Generate()
	}
}

func BenchmarkMultiBoolStubGenerator_Generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator := NewCombinatorBoolStubGenerator(hexaFieldNames, hexaFieldNames)

		_ = generator.Generate()
	}
}

func BenchmarkFastConditionMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastConditionMap(hexaFieldNames)
	}
}
