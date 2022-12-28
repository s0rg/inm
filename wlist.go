package inm

type WordList struct {
	Adjectives  []string
	Adverbs     []string
	Nouns       []string
	ProperNouns []string
	Verbs       []string
}

func (w *WordList) Add(v *WordList) {
	w.Adjectives = append(w.Adjectives, v.Adjectives...)
	w.Adverbs = append(w.Adverbs, v.Adverbs...)
	w.Nouns = append(w.Nouns, v.Nouns...)
	w.ProperNouns = append(w.ProperNouns, v.ProperNouns...)
	w.Verbs = append(w.Verbs, v.Verbs...)
}

func (w *WordList) Len() int {
	return len(w.Adjectives) + len(w.Adverbs) + len(w.Nouns) + len(w.ProperNouns) + len(w.Verbs)
}
