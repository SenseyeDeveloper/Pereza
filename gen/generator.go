package gen

import (
	"github.com/gopereza/pereza/core"
	"io"
	"reflect"
)

type Generator struct {
	packagePath string
	packageName string
	hashString  string

	types []reflect.Type
}

func NewGenerator(packagePath, packageName, filename string) *Generator {
	ret := &Generator{
		packagePath: packagePath,
		packageName: packageName,
		hashString:  unique(filename),

		types: make([]reflect.Type, 0, 1),
	}

	return ret
}

func (g *Generator) Add(obj interface{}) {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	g.types = append(g.types, t)
}

// Run runs the generator and outputs generated code to out.
func (g *Generator) Run(out io.Writer) error {
	out.Write(g.header())

	for _, t := range g.types {
		out.Write(g.genStructEncoder(t))
	}

	return nil
}

// header prints package declaration and imports.
func (g *Generator) header() []byte {
	const header = `// Code generated by pereza for marshaling/unmarshaling. DO NOT EDIT.

package `

	result := make([]byte, 0, len(header)+len(g.packageName)+2)

	result = append(result, header...)
	result = append(result, g.packageName...)
	result = append(result, '\n', '\n')

	return result
}

func (g *Generator) genStructEncoder(t reflect.Type) []byte {
	length := t.NumField()

	if length == 0 {
		return core.EmptyResultStub(t.Name())
	}

	switch length {
	case 1:
		field := t.Field(0)

		jsonName, standard := core.StandardStructureField(field)

		if standard {
			kind := field.Type.Kind()

			switch kind {
			case reflect.Bool:
				return core.BoolResultStub(t.Name(), field.Name, jsonName)
			case reflect.String:
				return core.StringResultStub(t.Name(), field.Name, jsonName)
			case reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64:

				return core.IntResultStubByType(t.Name(), field.Name, jsonName, kind)
			}
		}
	default:
		fieldsNames, jsonNames, standard := core.MultiBoolStandardStructure(t)

		if standard {
			if core.MatchAllBooleanFields(t) {
				if len(fieldsNames) > core.MultiBoolMaxProperties {
					// TODO
				}

				return core.CombinatorBoolResultStub(t.Name(), fieldsNames, jsonNames)
			}
		}
	}

	return core.EmptyResultStub(t.Name())
}
