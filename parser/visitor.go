package parser

import (
	"go/ast"
	"strings"
)

const structComment = "pereza:json"

type visitor struct {
	packageName string
	structNames []string
	name        string
	explicit    bool
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v

	case *ast.File:
		v.packageName = n.Name.String()

		return v

	case *ast.GenDecl:
		v.explicit = needType(n.Doc.Text())

		if !v.explicit {
			return nil
		}

		return v

	case *ast.TypeSpec:
		v.name = n.Name.String()

		// Allow to specify non-structs explicitly independent of '-all' flag.
		if v.explicit {
			v.structNames = append(v.structNames, v.name)
			return nil
		}

		return v

	case *ast.StructType:
		v.structNames = append(v.structNames, v.name)

		return nil
	}

	return nil
}

func needType(comments string) bool {
	for _, v := range strings.Split(comments, "\n") {
		if strings.HasPrefix(v, structComment) {
			return true
		}
	}

	return false
}
