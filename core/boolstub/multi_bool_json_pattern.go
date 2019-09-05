package boolstub

import "github.com/gopereza/pereza/core/common"

const (
	multiBoolJSONResultHeader    = "return []byte(`{"
	multiBoolJSONResultFooter    = "}`), nil"
	multiBoolJSONResultFixedSize = len(multiBoolJSONResultHeader) + len(multiBoolJSONResultFooter)
)

type MultiBoolJSONResultGenerator struct {
	jsonNames   []string
	buffer      []byte
	last        int
	avgCapacity int
}

func NewMultiBoolJSONResultGenerator(jsonNames []string) *MultiBoolJSONResultGenerator {
	length := len(jsonNames)

	const (
		wrapTrue  = 7 // len(`"":true`)
		wrapFalse = 8 // len(`"":true`)
	)

	commaCount := length - 1
	jsonNameLength := common.StringSliceSize(jsonNames)

	minCapacity := multiBoolJSONResultFixedSize + jsonNameLength + length*wrapTrue + commaCount
	maxCapacity := multiBoolJSONResultFixedSize + jsonNameLength + length*wrapFalse + commaCount

	return &MultiBoolJSONResultGenerator{
		jsonNames:   jsonNames,
		buffer:      make([]byte, 0, maxCapacity),
		last:        commaCount,
		avgCapacity: (minCapacity + maxCapacity + 1) / 2,
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

func (g *MultiBoolJSONResultGenerator) AvgCapacity() int {
	return g.avgCapacity
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
