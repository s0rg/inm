package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/s0rg/inm"
)

const (
	appName = "Impaled Northern Moon fortune"
	appSite = "https://github.com/s0rg/inm"
)

// build-time values.
var (
	GitTag    string
	GitHash   string
	BuildDate string
)

// command-line flags.
var (
	fVersion = flag.Bool("version", false, "version info")
)

func version() string {
	return fmt.Sprintf("%s %s-%s build at: %s with %s site: %s",
		appName,
		GitTag,
		GitHash,
		BuildDate,
		runtime.Version(),
		appSite,
	)
}

func main() {
	flag.Parse()

	if *fVersion {
		fmt.Println(version())

		return
	}

	gen := inm.New()

	fmt.Println(gen)
}
