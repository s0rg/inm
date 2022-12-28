package inm

type (
	Random func(int) int

	options struct {
		rnd Random
		lst WordList
	}

	Option func(*options)
)

func (o *options) validate() {
	if o.rnd == nil {
		rnd := defaultRandom()
		o.rnd = rnd.Rand
	}

	if o.lst.Len() == 0 {
		WithWords(&origWords)(o)
	}
}

func (o *options) buildGen() (g gen) {
	return &listGen{rnd: o.rnd, lst: o.lst}
}

func WithRandom(fn Random) Option {
	return func(o *options) {
		o.rnd = fn
	}
}

func WithWords(lst *WordList) Option {
	return func(o *options) {
		o.lst.Add(lst)
	}
}
