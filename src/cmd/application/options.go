package application

func WithQuery(q Query) Option {
	return func(a *app) error {
		a.query = q
		return nil
	}
}

func WithCommand(c Command) Option {
	return func(a *app) error {
		a.command = c
		return nil
	}
}
