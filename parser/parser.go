package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Parser struct {
	prefixes []string
	pwd      string
}

func NewParser(gopath string, pwd string) *Parser {
	return &Parser{
		prefixes: prefixes(gopath),
		pwd:      pwd,
	}
}

func (p *Parser) Parse(filename string, isDir bool) (Result, error) {
	packagePath, err := getPackagePathByEnv(p.prefixes, p.pwd, filename, isDir)

	if err != nil {
		return Result{}, err
	}

	fset := token.NewFileSet()
	resultVisitor := &visitor{}

	if isDir {
		packages, err := parser.ParseDir(fset, filename, nil, parser.ParseComments)
		if err != nil {
			return Result{}, err
		}

		for _, astPackage := range packages {
			ast.Walk(resultVisitor, astPackage)
		}
	} else {
		astFile, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			return Result{}, err
		}

		ast.Walk(resultVisitor, astFile)
	}

	return Result{
		PackagePath: packagePath,
		PackageName: resultVisitor.packageName,
		StructNames: resultVisitor.structNames,
	}, err
}
