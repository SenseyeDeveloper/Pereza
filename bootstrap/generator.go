package bootstrap

type Generator struct {
	settings Settings
}

func NewGenerator(settings Settings) *Generator {
	return &Generator{settings: settings}
}

func (g *Generator) Run() error {
	for _, filename := range g.settings.Filenames {
		err := g.generate(filename)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generate(filename string) error {
	return nil
}
