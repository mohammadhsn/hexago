package hexago

type Middleware interface {
	Exec(Cmd, func(Cmd))
}

type Chain []Middleware

func (ch Chain) Wrap(c Cmd) func(Cmd) {
	f := func(Cmd) {}

	for _, m := range ch {
		l := f
		mid := m
		f = func(c Cmd) {
			mid.Exec(c, l)
		}
	}

	return f
}
