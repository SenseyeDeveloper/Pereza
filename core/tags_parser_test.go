package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parseFieldTagsdataProvider = map[string]FieldTags{
	"": {},
	"-": {
		Omit: true,
	},
	"-,": {
		Name: "-",
	},
	"id": {
		Name: "id",
	},
	"value": {
		Name: "value",
	},
	"id,string": {
		Name:     "id",
		AsString: true,
	},
	"value,string": {
		Name:     "value",
		AsString: true,
	},
	"id,omitempty": {
		Name:      "id",
		OmitEmpty: true,
	},
	"value,omitempty": {
		Name:      "value",
		OmitEmpty: true,
	},
	"id,omitempty,string": {
		Name:      "id",
		OmitEmpty: true,
		AsString:  true,
	},
	"value,omitempty,string": {
		Name:      "value",
		OmitEmpty: true,
		AsString:  true,
	},
}

func TestParseFieldTags(t *testing.T) {
	for tags, expect := range parseFieldTagsdataProvider {
		assert.Equal(t, expect, ParseFieldTags(tags))
	}
}

func TestFieldTagsOnlyName(t *testing.T) {
	assert.True(t, ParseFieldTags("name").OnlyName())
	assert.True(t, ParseFieldTags("name,").OnlyName())
	assert.True(t, ParseFieldTags("-,").OnlyName())

	assert.False(t, ParseFieldTags("name,omitempty").OnlyName())
	assert.False(t, ParseFieldTags("name,string").OnlyName())
	assert.False(t, ParseFieldTags("-").OnlyName())
}

func BenchmarkParseFieldTags(b *testing.B) {
	n := b.N / len(parseFieldTagsdataProvider)

	for i := 0; i < n; i++ {
		for tags := range parseFieldTagsdataProvider {
			_ = ParseFieldTags(tags)
		}
	}
}

func BenchmarkEasyParseFieldTags(b *testing.B) {
	n := b.N / len(parseFieldTagsdataProvider)

	for i := 0; i < n; i++ {
		for tags := range parseFieldTagsdataProvider {
			_ = EasyParseFieldTags(tags)
		}
	}
}

func BenchmarkParseFieldTagsDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseFieldTags("name")
	}
}

func BenchmarkEasyParseFieldTagsDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = EasyParseFieldTags("name")
	}
}

func BenchmarkFieldTagsOnlyNameTrue(b *testing.B) {
	name := EasyParseFieldTags("name")

	for i := 0; i < b.N; i++ {
		_ = name.OnlyName()
	}
}

func BenchmarkFieldTagsOnlyNameFalse(b *testing.B) {
	name := EasyParseFieldTags("")

	for i := 0; i < b.N; i++ {
		_ = name.OnlyName()
	}
}
