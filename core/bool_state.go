package core

type BoolStateReplacer struct {
	stateSize int
	pool      [][]bool
	poolSize  int
}

func NewBoolStateReplacer(size int) *BoolStateReplacer {
	return &BoolStateReplacer{
		stateSize: size,
		pool:      nil,
		poolSize:  0,
	}
}

func (r *BoolStateReplacer) Replace(source []bool, n int, value bool) []bool {
	if source[n] == value {
		return source
	}

	result := r.poolGet()

	copy(result, source)

	result[n] = value

	return result
}

func (r *BoolStateReplacer) PoolPut(source []bool) {
	r.pool = append(r.pool, source)
	r.poolSize++
}

func (r *BoolStateReplacer) poolGet() []bool {
	if r.poolSize > 0 {
		r.poolSize--

		result := r.pool[r.poolSize]

		r.pool = r.pool[:r.poolSize]

		return result
	}

	return make([]bool, r.stateSize)
}

func FillBooleans(length int, value bool) []bool {
	result := make([]bool, length)

	for i := 0; i < length; i++ {
		result[i] = value
	}

	return result
}
