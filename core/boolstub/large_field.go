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

	stub = AppendFirstField(stub, fastConditionMap[fieldNames[0]], jsonNames[0])

	last := len(fieldNames) - 1
	for i := 1; i < last; i++ {
		stub = AppendField(stub, fastConditionMap[fieldNames[i]], jsonNames[i])
	}

	stub = AppendLastField(stub, fastConditionMap[fieldNames[last]], jsonNames[last])

	stub = append(stub, '\n')
	stub = append(stub, "return result, nil"...)
	stub = append(stub, '\n')

	stub = common.AppendFooter(stub)

	return stub
}

func AppendFirstField(result []byte, condition []byte, jsonName string) []byte {
	result = append(result, condition...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `{\""+jsonName+"\":true`...)\n"...)
	result = append(result, "} else {\n"...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `{\""+jsonName+"\":false`...)\n"...)
	result = append(result, "}\n"...)

	return result
}

func AppendField(result []byte, condition []byte, jsonName string) []byte {
	result = append(result, condition...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `,\""+jsonName+"\":true`...)\n"...)
	result = append(result, "} else {\n"...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `,\""+jsonName+"\":false`...)\n"...)
	result = append(result, "}\n"...)

	return result
}

func AppendLastField(result []byte, condition []byte, jsonName string) []byte {
	result = append(result, condition...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `,\""+jsonName+"\":true}`...)\n"...)
	result = append(result, "} else {\n"...)
	result = append(result, '\t')
	result = append(result, "result = append(result, `,\""+jsonName+"\":false}`...)\n"...)
	result = append(result, "}\n"...)

	return result
}
