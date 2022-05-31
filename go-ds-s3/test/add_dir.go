package test

import (
	"flag"

	"github.com/testground/sdk-go/runtime"
)

// var flagExp = flag.Bool("experimental", false, "enable experimental features")

// Testa
func AddDir(runenv *runtime.RunEnv) error {
	var (
		iterations = runenv.IntParam("iterations")
	)
	runenv.RecordMessage("started test instance, iterations=%d", iterations)

	flag.Parse()

	runenv.RecordMessage("all done")
	return nil
}
