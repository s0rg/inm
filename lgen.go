package inm

type listGen struct {
	rnd Random
	lst *Dictionary
}

func (lg *listGen) Adjective() string {
	return lg.lst.Adjectives[lg.rnd(len(lg.lst.Adjectives))]
}

func (lg *listGen) Adverb() string {
	return lg.lst.Adverbs[lg.rnd(len(lg.lst.Adverbs))]
}

func (lg *listGen) Noun() string {
	return lg.lst.Nouns[lg.rnd(len(lg.lst.Nouns))]
}

func (lg *listGen) ProperNoun() string {
	return lg.lst.ProperNouns[lg.rnd(len(lg.lst.ProperNouns))]
}

func (lg *listGen) Verb() string {
	return lg.lst.Verbs[lg.rnd(len(lg.lst.Verbs))]
}
