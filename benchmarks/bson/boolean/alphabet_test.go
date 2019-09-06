package boolean

import (
	"github.com/gopereza/pereza/fixtures/bson/boolean"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var (
	alphabetDataProvider = map[bool]string{
		false: "\x6d\x00\x00\x00\x08a\x00\x00\x08b\x00\x00\x08c\x00\x00\x08d\x00\x00\x08e\x00\x00\x08f\x00\x00\x08g\x00\x00\x08h\x00\x00\x08i\x00\x00\x08j\x00\x00\x08k\x00\x00\x08l\x00\x00\x08m\x00\x00\x08n\x00\x00\x08o\x00\x00\x08p\x00\x00\x08q\x00\x00\x08r\x00\x00\x08s\x00\x00\x08t\x00\x00\x08u\x00\x00\x08v\x00\x00\x08w\x00\x00\x08x\x00\x00\x08y\x00\x00\x08z\x00\x00\x00",
		true:  "\x6d\x00\x00\x00\x08a\x00\x01\x08b\x00\x01\x08c\x00\x01\x08d\x00\x01\x08e\x00\x01\x08f\x00\x01\x08g\x00\x01\x08h\x00\x01\x08i\x00\x01\x08j\x00\x01\x08k\x00\x01\x08l\x00\x01\x08m\x00\x01\x08n\x00\x01\x08o\x00\x01\x08p\x00\x01\x08q\x00\x01\x08r\x00\x01\x08s\x00\x01\x08t\x00\x01\x08u\x00\x01\x08v\x00\x01\x08w\x00\x01\x08x\x00\x01\x08y\x00\x01\x08z\x00\x01\x00",
	}
)

func TestAlphabetMongoMarshalBSON(t *testing.T) {
	for state, expect := range alphabetDataProvider {
		source := boolean.AlphabetBoolState{
			A: state,
			B: state,
			C: state,
			D: state,
			E: state,
			F: state,
			G: state,
			H: state,
			I: state,
			J: state,
			K: state,
			L: state,
			M: state,
			N: state,
			O: state,
			P: state,
			Q: state,
			R: state,
			S: state,
			T: state,
			U: state,
			V: state,
			W: state,
			X: state,
			Y: state,
			Z: state,
		}

		actual, err := bson.Marshal(source)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func TestAlphabetPerezaMarshalBSON(t *testing.T) {
	for state, expect := range alphabetDataProvider {
		source := boolean.PerezaAlphabetBoolState{
			A: state,
			B: state,
			C: state,
			D: state,
			E: state,
			F: state,
			G: state,
			H: state,
			I: state,
			J: state,
			K: state,
			L: state,
			M: state,
			N: state,
			O: state,
			P: state,
			Q: state,
			R: state,
			S: state,
			T: state,
			U: state,
			V: state,
			W: state,
			X: state,
			Y: state,
			Z: state,
		}

		actual, err := source.MarshalBSON()
		assert.NoError(t, err)
		assert.Equal(t, []byte(expect), actual)
	}
}

func BenchmarkAlphabetMongoMarshalBSON(b *testing.B) {
	source := boolean.AlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = bson.Marshal(source)
	}
}

func BenchmarkAlphabetPerezaMarshalBSON(b *testing.B) {
	source := boolean.PerezaAlphabetBoolState{}

	for i := 0; i < b.N; i++ {
		_, _ = source.MarshalBSON()
	}
}
