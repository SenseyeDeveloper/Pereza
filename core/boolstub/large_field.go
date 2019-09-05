package boolstub

import (
	"github.com/gopereza/pereza/core/common"
	"strconv"
)

func LargeFieldStub(typeName string, fieldNames, jsonNames []string) []byte {
	fastConditionMap := FastConditionMap(fieldNames)

	var stub []byte
	stub = common.AppendHeader(stub, typeName)

	stub = append(stub, "result := make([]byte, 0, "...)
	stub = strconv.AppendUint(stub, uint64(MultiSize(jsonNames)), 10)
	stub = append(stub, ')', '\n', '\n')

	{
		fieldName := fieldNames[0]
		jsonName := jsonNames[0]
		stub = append(stub, fastConditionMap[fieldName]...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `{\""+jsonName+"\":true`...)\n"...)
		stub = append(stub, "} else {\n"...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `{\""+jsonName+"\":false`...)\n"...)
		stub = append(stub, "}\n"...)
	}

	last := len(fieldNames) - 1
	for i := 1; i < last; i++ {
		fieldName := fieldNames[i]
		jsonName := jsonNames[i]
		stub = append(stub, fastConditionMap[fieldName]...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `,\""+jsonName+"\":true`...)\n"...)
		stub = append(stub, "} else {\n"...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `,\""+jsonName+"\":false`...)\n"...)
		stub = append(stub, "}\n"...)
	}

	{
		fieldName := fieldNames[last]
		jsonName := jsonNames[last]
		stub = append(stub, fastConditionMap[fieldName]...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `,\""+jsonName+"\":true}`...)\n"...)
		stub = append(stub, "} else {\n"...)
		stub = append(stub, '\t')
		stub = append(stub, "result = append(result, `,\""+jsonName+"\":false}`...)\n"...)
		stub = append(stub, "}\n"...)
	}

	stub = append(stub, '\n')
	stub = append(stub, "return result, nil"...)
	stub = append(stub, '\n')

	stub = common.AppendFooter(stub)

	return stub
}
