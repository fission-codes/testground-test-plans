package test

import (
	"flag"
	"math/rand"

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

	s := randomString(32)

	runenv.RecordMessage("random string = ", s)
	runenv.RecordMessage("all done")
	return nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
