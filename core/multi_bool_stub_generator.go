package core

const (
	conditionStart     = "if v."
	conditionEnd       = " {\n"
	conditionFixedSize = len(conditionStart) + len(conditionEnd)
)

type MultiBoolStubGenerator struct {
	fieldNames       []string
	fastConditionMap map[string][]byte
	pattern          *MultiBoolJSONResultGenerator
	replacer         *BoolStateReplacer
	buffer           []byte
	returnDepth      int
	capacity         int
}

func NewMultiBoolStubGenerator(fieldNames, jsonNames []string) *MultiBoolStubGenerator {
	pattern := NewMultiBoolJSONResultGenerator(jsonNames)

	length := len(fieldNames)

	capacity := pattern.Capacity() * length * length * length * length * 2

	return &MultiBoolStubGenerator{
		fieldNames:       fieldNames,
		fastConditionMap: FastConditionMap(fieldNames),
		pattern:          pattern,
		replacer:         NewBoolStateReplacer(length),
		buffer:           make([]byte, 0, capacity), // dynamic allocate
		returnDepth:      length - 1,
		capacity:         capacity,
	}
}

func (g *MultiBoolStubGenerator) Generate() []byte {
	g.generate(0, FillBooleans(len(g.fieldNames), true))

	return g.buffer
}

func (g *MultiBoolStubGenerator) generate(depth int, states []bool) {
	fieldName := g.fieldNames[depth]

	trueState := g.replacer.Replace(states, depth, true)
	falseState := g.replacer.Replace(states, depth, false)

	if depth == g.returnDepth {
		g.append(g.fastConditionMap[fieldName])
		g.append(g.pattern.Generate(trueState))
		g.conditionClose()

		g.append(g.pattern.Generate(falseState))
	} else {
		g.append(g.fastConditionMap[fieldName])
		g.generate(depth+1, trueState)
		g.conditionClose()

		g.generate(depth+1, falseState)
	}

	g.replacer.PoolPut(trueState)
	g.replacer.PoolPut(falseState)
}

func (g *MultiBoolStubGenerator) appendString(code string) {
	g.buffer = append(g.buffer, code...)
}

func (g *MultiBoolStubGenerator) append(data []byte) {
	g.buffer = append(g.buffer, data...)
}

func (g *MultiBoolStubGenerator) conditionClose() {
	g.buffer = append(g.buffer, '}', '\n')
}

func FastConditionMap(fieldNames []string) map[string][]byte {
	length := len(fieldNames)

	fastConditionMap := make(map[string][]byte, length)

	capacity := stringSliceSize(fieldNames) + length*conditionFixedSize
	once := make([]byte, 0, capacity)

	for _, fieldName := range fieldNames {
		current := conditionFixedSize + len(fieldName)

		once = append(once, conditionStart...)
		once = append(once, fieldName...)
		once = append(once, conditionEnd...)

		fastConditionMap[fieldName] = once[:current]

		once = once[current:]
	}

	return fastConditionMap
}
