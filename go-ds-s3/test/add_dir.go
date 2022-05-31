package test

import (
	crand "crypto/rand"
	"flag"
	"math/rand"
	"os"
	"time"

	randomfiles "github.com/jbenet/go-random-files"
	"github.com/testground/sdk-go/runtime"
)

func AddDir(runenv *runtime.RunEnv) error {
	var (
		iterations = runenv.IntParam("iterations")
	)

	opts := randomfiles.Options{
		FileSize:     4096,
		FanoutDepth:  2,
		FanoutDirs:   5,
		FanoutFiles:  10,
		RandomSeed:   0,
		RandomFanout: false,
		RandomSize:   true,
		Alphabet:     randomfiles.RunesEasy,
		Out:          os.Stdout,
		Source:       crand.Reader,
	}

	rand.Seed(time.Now().UnixNano())
	runenv.RecordMessage("started test instance, iterations=%d", iterations)

	flag.Parse()

	root := "/tmp/testground"
	if err := os.MkdirAll(root, 0755); err != nil {
		return err
	}

	dirPath, err := runenv.CreateRandomDirectory(root, 1)
	if err != nil {
		return err
	}

	err = randomfiles.WriteRandomFiles(dirPath, 1, &opts)
	if err != nil {
		return err
	}

	runenv.RecordMessage("all done")
	return nil
}
