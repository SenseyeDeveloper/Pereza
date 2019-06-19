package core

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	intReflectTypes = []reflect.Kind{
		reflect.Int8,
		reflect.Uint8,
		reflect.Int,
		reflect.Uint,
		reflect.Int64,
		reflect.Uint64,
	}
)

func TestIntResultStub(t *testing.T) {
	for _, k := range intReflectTypes {
		assertIntResultStubOneAllocation(t, "BoolState", "SomeState", "some_state", k)
		assertIntResultStubOneAllocation(t, "PerezaBoolState", "State", "state", k)
	}
}

func BenchmarkIntResultStub(b *testing.B) {
	length := len(intReflectTypes)

	for i := 0; i < b.N; i++ {
		_ = IntResultStubByType("PerezaBoolState", "State", "state", intReflectTypes[i%length])
	}
}

func assertIntResultStubOneAllocation(t *testing.T, typeName, fieldName, jsonName string, k reflect.Kind) {
	t.Helper()

	expectSize := getIntResultStubSizeByType(typeName, fieldName, jsonName, k)
	assert.Equal(t, expectSize, len(IntResultStubByType(typeName, fieldName, jsonName, k)))
}
