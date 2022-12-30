package inm

type (
	// Random is a random generator func, takes N, return random in [0, N-1].
	Random func(int) int

	options struct {
		rnd Random
		lst *Dictionary
	}

	// Option is a functional option type.
	Option func(*options)
)

func (o *options) validate() {
	if o.rnd == nil {
		rnd := defaultRandom()
		o.rnd = rnd.Rand
	}

	if o.lst == nil {
		WithWords(&origWords)(o)
	}
}

func (o *options) buildGen() (g gen) {
	return &listGen{rnd: o.rnd, lst: o.lst}
}

// WithRandom sets random generator, if not set 'math/rand' will be used.
func WithRandom(fn Random) Option {
	return func(o *options) {
		o.rnd = fn
	}
}

// WithWords allows to set custom dictionary.
func WithWords(lst *Dictionary) Option {
	return func(o *options) {
		if o.lst == nil {
			o.lst = &Dictionary{}
		}

		o.lst.Add(lst)
	}
}
