package json

import (
	fixtures "github.com/gopereza/pereza/fixtures/bson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

// http://bsonspec.org/spec.html
// \x05 = Binary data
const (
	expectEmptyStateBSON = "\x05\x00\x00\x00\x00"
)

func TestEmptyStateMongoMarshalBSON(t *testing.T) {
	source := fixtures.EmptyState{}

	actual, err := bson.Marshal(source)
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectEmptyStateBSON), actual)
}

func TestEmptyStatePerezaMarshalBSON(t *testing.T) {
	source := fixtures.PerezaEmptyState{}

	actual, err := source.MarshalBSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectEmptyStateBSON), actual)
}

func BenchmarkEmptyStateMongoMarshalBSON(b *testing.B) {
	source := fixtures.EmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = bson.Marshal(source)
	}
}

func BenchmarkEmptyStatePerezaMarshalBSON(b *testing.B) {
	source := fixtures.PerezaEmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalBSON()
	}
}
