package inm

import (
	"math/rand"
	"time"
)

type random struct {
	src rand.Source
}

func defaultRandom() (r *random) {
	return &random{src: rand.NewSource(time.Now().UnixNano())}
}

func (r *random) Rand(n int) (rv int) {
	return int(r.src.Int63() % int64(n))
}
