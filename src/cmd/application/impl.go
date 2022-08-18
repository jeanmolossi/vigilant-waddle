package application

type app struct {
	query   Query
	command Command
}

func New(opts ...Option) Application {
	a := &app{}

	for _, opt := range opts {
		if err := opt(a); err != nil {
			panic(err.Error())
		}
	}

	return a
}

func (a *app) Run() {}

func (a *app) Query() Query {
	return a.query
}

func (a *app) Command() Command {
	return a.command
}
