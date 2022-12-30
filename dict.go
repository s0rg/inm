package inm

// Dictionary holds grouped words for generator.
type Dictionary struct {
	Adjectives  []string
	Adverbs     []string
	Nouns       []string
	ProperNouns []string
	Verbs       []string
}

// Add merges two dictionaries.
func (d *Dictionary) Add(v *Dictionary) {
	d.Adjectives = append(d.Adjectives, v.Adjectives...)
	d.Adverbs = append(d.Adverbs, v.Adverbs...)
	d.Nouns = append(d.Nouns, v.Nouns...)
	d.ProperNouns = append(d.ProperNouns, v.ProperNouns...)
	d.Verbs = append(d.Verbs, v.Verbs...)
}
