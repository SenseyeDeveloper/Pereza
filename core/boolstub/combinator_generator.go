package boolstub

import "github.com/gopereza/pereza/core/common"

const (
	conditionStart     = "if v."
	conditionEnd       = " {\n"
	conditionFixedSize = len(conditionStart) + len(conditionEnd)
)

// Dynamic allocate
func CombinatorBoolResultStub(typeName string, fieldNames, jsonNames []string) []byte {
	generator := NewCombinatorGenerator(fieldNames, jsonNames)

	body := generator.Generate()

	result := make([]byte, 0, common.WrapSignatureSize+len(body))

	result = common.AppendHeader(result, typeName)
	result = append(result, body...)
	result = common.AppendFooter(result)

	return result
}

type CombinatorGenerator struct {
	fieldNames       []string
	fastConditionMap map[string][]byte
	pattern          *DumpGenerator
	replacer         *BoolStateReplacer
	buffer           []byte
	returnDepth      int
}

func NewCombinatorGenerator(fieldNames, jsonNames []string) *CombinatorGenerator {
	pattern := NewDumpGenerator(jsonNames)

	length := len(fieldNames)

	return &CombinatorGenerator{
		fieldNames:       fieldNames,
		fastConditionMap: FastConditionMap(fieldNames),
		pattern:          pattern,
		replacer:         NewBoolStateReplacer(length),
		buffer:           make([]byte, 0, 1024), // dynamic allocate
		returnDepth:      length - 1,
	}
}

func (g *CombinatorGenerator) Generate() []byte {
	g.generate(0, FillBooleans(len(g.fieldNames), true))

	return g.buffer
}

func (g *CombinatorGenerator) generate(depth int, states []bool) {
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

func (g *CombinatorGenerator) append(data []byte) {
	g.buffer = append(g.buffer, data...)
}

func (g *CombinatorGenerator) conditionClose() {
	g.buffer = append(g.buffer, '}', '\n')
}

func FastConditionMap(fieldNames []string) map[string][]byte {
	length := len(fieldNames)

	fastConditionMap := make(map[string][]byte, length)

	capacity := common.StringSliceSize(fieldNames) + length*conditionFixedSize
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
