package test

import "github.com/testground/sdk-go/runtime"

//
// To run use:
//
// ./testground run s go-ds-s3/testa \
//   --builder=exec:go \
//   --runner="local:exec" \
//   --dep="github.com/fission-suite/go-ipfs=master" \
//   -instances=8 \
//   --test-param iterations=100
//

// Testa
func Testa(runenv *runtime.RunEnv) error {
	// 📝  Consume test parameters from the runtime environment.
	var (
		iterations = runenv.IntParam("iterations")
	)
	runenv.RecordMessage("started test instance, iterations=%d", iterations)
	return nil
}
