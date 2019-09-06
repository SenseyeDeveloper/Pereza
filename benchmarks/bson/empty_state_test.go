package json

import (
	"github.com/gopereza/pereza/fixtures"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

// http://bsonspec.org/spec.html
// \x05 = Binary data
const (
	expectEmptyStateBSON = "\x05\x00\x00\x00\x00"
)

func TestEmptyStateEncodingJSON(t *testing.T) {
	source := fixtures.EmptyState{}

	actual, err := bson.Marshal(source)
	assert.NoError(t, err)
	assert.Equal(t, []byte(expectEmptyStateBSON), actual)
}

func BenchmarkEmptyStateEncodingJSON(b *testing.B) {
	source := fixtures.EmptyState{}

	for i := 0; i < b.N; i++ {
		_, _ = bson.Marshal(source)
	}
}
