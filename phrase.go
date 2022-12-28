package inm

import "strings"

const space = ' '

type phrase struct {
	b strings.Builder
}

func (p *phrase) Add(v ...string) {
	for _, s := range v {
		p.b.WriteString(s)
		p.b.WriteRune(space)
	}
}

func (p *phrase) String() (rv string) {
	return p.b.String()
}
