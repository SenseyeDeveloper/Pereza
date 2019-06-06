package parser

import "fmt"

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

func (p *Parser) Parse(filename string, isDir bool) error {
	packagePath, err := getPackagePathByEnv(p.prefixes, p.pwd, filename, isDir)

	if err != nil {
		return err
	}

	fmt.Println(packagePath)

	return nil
}
