package hexago

type Middleware interface {
	Exec(Cmd, func(Cmd))
}

type Chain struct {
	mw []Middleware
}

func NewChain(mw ...Middleware) *Chain {
	return &Chain{mw}
}

func (ch *Chain) Wrap(c Cmd) func(Cmd) {
	f := func(Cmd) {}

	for _, m := range ch.mw {
		l := f
		mid := m
		f = func(c Cmd) {
			mid.Exec(c, l)
		}
	}

	return f
}

func (ch *Chain) Add(mw Middleware) {
	ch.mw = append(ch.mw, mw)
}
