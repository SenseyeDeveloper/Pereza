package complexstub

import "github.com/gopereza/pereza/core/common"

const (
	wrap = 3 // len(`"":`)
)

func WrapMultiSize(jsonNames []string) int {
	const brackets = 2

	length := len(jsonNames)
	commaCount := length - 1
	jsonNameLength := common.StringSliceSize(jsonNames)

	return brackets + jsonNameLength + length*wrap + commaCount
}
