package inm

const (
	cOf      = "of"
	cOn      = "on"
	cAnd     = "and"
	cThe     = "the"
	cUpon    = "upon"
	cWhile   = "while"
	cAmidst  = "amidst"
	cAmongst = "amongst"
)

type gen interface {
	Adjective() string
	Adverb() string
	Noun() string
	ProperNoun() string
	Verb() string
}

// SongTitleGenerator is a generator instance.
type SongTitleGenerator struct {
	gen  gen
	opts options
}

// New creates new generator instance.
func New(opts ...Option) (rv *SongTitleGenerator) {
	var stg SongTitleGenerator

	for _, o := range opts {
		o(&stg.opts)
	}

	stg.opts.validate()

	stg.gen = stg.opts.buildGen()

	return &stg
}

// String prodece new song title every upon every call.
func (stg *SongTitleGenerator) String() string {
	switch stg.opts.rnd(3) {
	case 2:
		return stg.v1()
	case 1:
		return stg.v2()
	default:
	}

	return stg.v3()
}

func (stg *SongTitleGenerator) v1() (rv string) {
	var p phrase

	p.Add(stg.gen.Adjective(), cAnd, stg.gen.Adjective(), stg.gen.Noun())

	switch stg.opts.rnd(3) {
	case 2:
		p.Add(cOf, cThe, stg.gen.Noun())
	case 1:
		p.Add(stg.gen.Adverb(), stg.gen.Verb())

		if stg.opts.rnd(2) == 1 {
			p.Add(cThe, stg.gen.Noun(), cOf, stg.gen.ProperNoun())
		} else {
			p.Add(cUpon, cThe, stg.gen.Adjective()+stg.gen.Noun())
		}
	default:
		p.Add(stg.gen.Verb())

		switch stg.opts.rnd(3) {
		case 2:
			p.Add(
				cAmidst,
				cThe,
				stg.gen.Adjective(),
				stg.gen.Adjective()+stg.gen.Noun(),
				cOf,
				cThe,
				stg.gen.Noun())
		case 1:
			p.Add(
				cThe,
				stg.gen.Adjective(),
				stg.gen.Adjective()+stg.gen.Noun(),
				cOf,
				stg.gen.ProperNoun(),
			)
		default:
		}
	}

	return p.String()
}

func (stg *SongTitleGenerator) v2() (rv string) {
	var p phrase

	p.Add(stg.gen.Adverb(), stg.gen.Verb())

	switch stg.opts.rnd(3) {
	case 2:
		p.Add(cOn, cThe)

		if stg.opts.rnd(2) == 1 {
			p.Add(stg.gen.Adjective(), stg.gen.Noun())

			if stg.opts.rnd(2) == 1 {
				p.Add(cOf, cThe, stg.gen.Adjective(), stg.gen.Noun())
			} else {
				p.Add(cUpon, cThe, stg.gen.Noun(), cOf, stg.gen.ProperNoun())
			}
		} else {
			p.Add(stg.gen.Noun())

			switch stg.opts.rnd(3) {
			case 2:
				p.Add(cOf, stg.gen.ProperNoun())
			case 1:
				p.Add(cOf, cThe, stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
			default:
				p.Add(cWhile, stg.gen.Verb(), cThe, stg.gen.Adjective(), stg.gen.Adjective(), stg.gen.Noun())
			}
		}
	case 1:
		p.Add(cUpon, cThe)

		if stg.opts.rnd(2) == 1 {
			p.Add(stg.gen.Adjective(), stg.gen.Noun())

			if stg.opts.rnd(2) == 1 {
				p.Add(cOf, stg.gen.ProperNoun())
			} else {
				p.Add(cOf, cThe)

				switch stg.opts.rnd(3) {
				case 2:
					p.Add(stg.gen.Adjective(), stg.gen.Noun())
				case 1:
					p.Add(stg.gen.Adjective(), stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				default:
					p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				}
			}
		} else {
			p.Add(stg.gen.Noun(), cOf)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.ProperNoun())
			} else {
				p.Add(
					cThe,
					stg.gen.Adjective(),
					stg.gen.Adjective(),
					stg.gen.Noun(),
					cOf,
					cThe,
					stg.gen.Noun(),
				)
			}
		}
	default:
		p.Add(cThe)

		if stg.opts.rnd(2) == 1 {
			p.Add(stg.gen.Adjective())

			switch stg.opts.rnd(3) {
			case 2:
				p.Add(stg.gen.Adjective(), stg.gen.Noun())

				if stg.opts.rnd(2) == 1 {
					p.Add(cOf, stg.gen.ProperNoun())
				} else {
					p.Add(cOf, cThe, stg.gen.Noun())
				}
			case 1:
				p.Add(
					cAnd,
					stg.gen.Adjective(),
					stg.gen.Noun(),
					cWhile,
					stg.gen.Verb(),
					cThe,
					stg.gen.Adjective(),
					stg.gen.Noun(),
				)
			default:
				p.Add(stg.gen.Noun())

				switch stg.opts.rnd(3) {
				case 2:
					p.Add(cOf, stg.gen.ProperNoun())
				case 1:
					p.Add(cWhile, stg.gen.Verb(), cThe, stg.gen.Adjective(), stg.gen.Adjective(), stg.gen.Noun())
				default:
					p.Add(cOf, cThe, stg.gen.Adjective())

					if stg.opts.rnd(2) == 1 {
						p.Add(stg.gen.Noun())
					} else {
						p.Add(stg.gen.Adjective(), stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
					}
				}
			}
		} else {
			p.Add(stg.gen.Noun())

			switch stg.opts.rnd(3) {
			case 2:
				p.Add(cUpon, cThe)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Adjective())
				}

				p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
			case 1:
				p.Add(cWhile, stg.gen.Verb(), cThe, stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Noun())
				} else {
					p.Add(cAnd, stg.gen.Adjective(), stg.gen.Noun())
				}
			default:
				p.Add(cOf)

				switch stg.opts.rnd(3) {
				case 2:
					p.Add(stg.gen.ProperNoun())
				case 1:
					p.Add(cThe, stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
				default:
					p.Add(cThe, stg.gen.Adjective())

					if stg.opts.rnd(2) == 1 {
						p.Add(stg.gen.Adjective(), stg.gen.Noun(), cOf, stg.gen.ProperNoun())
					} else {
						p.Add(stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
					}
				}
			}
		}
	}

	return p.String()
}

func (stg *SongTitleGenerator) v3() (rv string) {
	var p phrase

	p.Add(stg.gen.Adjective(), stg.gen.Noun(), stg.gen.Verb(), cThe, stg.gen.Adjective())

	switch stg.opts.rnd(3) {
	case 2:
		p.Add(stg.gen.Adjective(), stg.gen.Noun())

		switch stg.opts.rnd(4) {
		case 3:
			p.Add(cAmidst, cThe, stg.gen.Noun(), cOf, stg.gen.ProperNoun())
		case 2:
			p.Add(cAmongst, cThe, stg.gen.Noun(), cOf)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Adjective(), stg.gen.Noun())
			} else {
				p.Add(stg.gen.ProperNoun())
			}
		case 1:
			p.Add(cUpon, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(cAnd, stg.gen.Adjective())
				}

				p.Add(stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
			} else {
				p.Add(stg.gen.Noun(), cOf)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.ProperNoun())
				} else {
					p.Add(cThe, stg.gen.Adjective(), stg.gen.Noun())
				}
			}
		default:
			p.Add(cOn, cThe)

			if stg.opts.rnd(2) == 1 {
				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Adjective())
				}

				p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
			} else {
				p.Add(stg.gen.Adjective(), cAnd, stg.gen.Adjective(), stg.gen.Noun(), cOf)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.ProperNoun())
				} else {
					p.Add(cThe, stg.gen.Adjective(), stg.gen.Noun())
				}
			}
		}
	case 1:
		p.Add(cAnd, stg.gen.Adjective(), stg.gen.Noun())

		switch stg.opts.rnd(4) {
		case 3:
			p.Add(
				cUpon,
				cThe,
				stg.gen.Adjective(),
				cAnd,
				stg.gen.Adjective(),
				stg.gen.Noun(),
				cOf,
				cThe,
				stg.gen.Noun(),
			)
		case 2:
			p.Add(cAmongst, cThe, stg.gen.Adjective(), stg.gen.Noun(), cOf)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Noun())
			} else {
				p.Add(stg.gen.ProperNoun())
			}
		case 1:
			p.Add(cAmidst, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Noun(), cOf, stg.gen.Adjective(), stg.gen.Noun())
			} else {
				p.Add(stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				} else {
					p.Add(cAnd, stg.gen.Adjective(), stg.gen.Noun(), cOf, stg.gen.Adjective(), stg.gen.Noun())
				}
			}
		default:
			p.Add(cOn, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Noun(), cOf)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.ProperNoun())
				} else {
					p.Add(cThe, stg.gen.Noun())
				}
			} else {
				p.Add(stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Adjective(), stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				} else {
					p.Add(stg.gen.Noun(), cOf, cThe, stg.gen.Adjective(), stg.gen.Noun())
				}
			}
		}
	default:
		p.Add(stg.gen.Noun())

		switch stg.opts.rnd(4) {
		case 3:
			p.Add(cAmongst, cThe)

			if stg.opts.rnd(2) == 1 {
				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Adjective())
				}

				p.Add(stg.gen.Noun(), cOf, stg.gen.Noun())
			} else {
				p.Add(
					stg.gen.Adjective(),
					cAnd,
					stg.gen.Adjective(),
					stg.gen.Noun(),
					cOf,
					stg.gen.Adjective(),
					stg.gen.Noun(),
				)
			}
		case 2:
			p.Add(cAmidst, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(cAnd, stg.gen.Adjective(), stg.gen.Noun(), cOf, stg.gen.Noun())
				} else {
					p.Add(stg.gen.Noun(), cOf, stg.gen.Adjective(), stg.gen.Noun())
				}
			} else {
				p.Add(stg.gen.Noun(), cOf)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Noun())
				} else {
					p.Add(stg.gen.ProperNoun())
				}
			}
		case 1:
			p.Add(cOn, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Adjective())

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				} else {
					if stg.opts.rnd(2) == 1 {
						p.Add(stg.gen.Adjective())
					} else {
						p.Add(cAnd, stg.gen.Adjective())
					}

					p.Add(stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
				}
			} else {
				p.Add(stg.gen.Noun(), cOf)

				switch stg.opts.rnd(3) {
				case 2:
					p.Add(stg.gen.ProperNoun())
				case 1:
					p.Add(cThe, stg.gen.Noun())
				default:
					p.Add(cThe, stg.gen.Adjective(), stg.gen.Noun())
				}
			}
		default:
			p.Add(cUpon, cThe)

			if stg.opts.rnd(2) == 1 {
				p.Add(stg.gen.Noun(), cOf)

				if stg.opts.rnd(2) == 1 {
					p.Add(stg.gen.ProperNoun())
				} else {
					p.Add(cThe, stg.gen.Adjective(), stg.gen.Noun())
				}
			} else {
				p.Add(stg.gen.Adjective())

				switch stg.opts.rnd(3) {
				case 2:
					p.Add(
						cAnd,
						stg.gen.Adjective(),
						stg.gen.Noun(),
						cOf,
						cThe,
						stg.gen.Adjective(),
						stg.gen.Noun(),
					)
				case 1:
					p.Add(stg.gen.Noun(), cOf, cThe, stg.gen.Noun())
				default:
					if stg.opts.rnd(2) == 1 {
						p.Add(stg.gen.Adjective())
					}

					p.Add(stg.gen.Noun(), cOf, stg.gen.ProperNoun())
				}
			}
		}
	}

	return p.String()
}
