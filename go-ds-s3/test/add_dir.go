package test

import (
	"flag"
	"os"

	randomfiles "github.com/jbenet/go-random-files"
	"github.com/testground/sdk-go/runtime"
)

// var flagExp = flag.Bool("experimental", false, "enable experimental features")

// Testa
func AddDir(runenv *runtime.RunEnv) error {
	var (
		iterations = runenv.IntParam("iterations")
	)

	opts := randomfiles.Options{
		FileSize:    1024,
		FanoutDepth: 5,
		FanoutFiles: 10,
		FanoutDirs:  5,
		RandomSize:  true,
	}

	runenv.RecordMessage("started test instance, iterations=%d", iterations)

	flag.Parse()

	root := "/tmp/testground"

	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}

	err := randomfiles.WriteRandomFiles(root, 1, &opts)
	if err != nil {
		return err
	}

	runenv.RecordMessage("random string = ", s)
	runenv.RecordMessage("all done")
	return nil
}
