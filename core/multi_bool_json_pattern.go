package core

const (
	multiBoolJSONResultHeader    = "return []byte(`{"
	multiBoolJSONResultFooter    = "}`), nil"
	multiBoolJSONResultFixedSize = len(multiBoolJSONResultHeader) + len(multiBoolJSONResultFooter)
)

type MultiBoolJSONResultGenerator struct {
	jsonNames []string
	buffer    []byte
	last      int
	capacity  int
}

func NewMultiBoolJSONResultGenerator(jsonNames []string) *MultiBoolJSONResultGenerator {
	length := len(jsonNames)

	const wrap = 8 // len(`"":false`)

	commaCount := length - 1

	capacity := multiBoolJSONResultFixedSize + stringSliceSize(jsonNames) + length*wrap + commaCount

	return &MultiBoolJSONResultGenerator{
		jsonNames: jsonNames,
		buffer:    make([]byte, 0, capacity),
		last:      commaCount,
		capacity:  capacity,
	}
}

func (g *MultiBoolJSONResultGenerator) Generate(values []bool) []byte {
	result := g.buffer[:0]

	result = append(result, multiBoolJSONResultHeader...)

	for i := 0; i < g.last; i++ {
		result = AppendBool(result, g.jsonNames[i], values[i])

		result = append(result, ',')
	}

	result = AppendBool(result, g.jsonNames[g.last], values[g.last])

	result = append(result, multiBoolJSONResultFooter...)

	return result
}

func (g *MultiBoolJSONResultGenerator) Capacity() int {
	return g.capacity
}

func AppendBool(source []byte, jsonName string, value bool) []byte {
	result := append(source, '"')
	result = append(result, jsonName...)
	result = append(result, '"', ':')

	if value {
		result = append(result, "true"...)
	} else {
		result = append(result, "false"...)
	}

	return result
}
