package inm

import (
	"bufio"
	"embed"
	"path"
	"strings"
)

// assets is original generator words lists.
//
//go:embed assets/*.txt
var assets embed.FS

const assetsRoot = "assets"

var origWords = WordList{
	Adjectives:  load("adjectives.txt"),
	Adverbs:     load("adverbs.txt"),
	Nouns:       load("nouns.txt"),
	ProperNouns: load("proper-nouns.txt"),
	Verbs:       load("verbs.txt"),
}

func load(name string) (rv []string) {
	fd, _ := assets.Open(path.Join(assetsRoot, name))
	defer fd.Close()

	var line string

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		rv = append(rv, line)
	}

	return rv
}
