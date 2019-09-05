package boolstub

import "github.com/gopereza/pereza/core/common"

const (
	wrapTrue  = 7 // len(`"":true`)
	wrapFalse = 8 // len(`"":false`)
)

func MultiSize(jsonNames []string) int {
	const brackets = 2

	length := len(jsonNames)
	commaCount := length - 1
	jsonNameLength := common.StringSliceSize(jsonNames)

	return brackets + jsonNameLength + length*wrapFalse + commaCount
}