package test

import "github.com/testground/sdk-go/runtime"

// Testa
func Testb(runenv *runtime.RunEnv) error {
	// 📝  Consume test parameters from the runtime environment.
	var (
		iterations = runenv.IntParam("iterations")
	)
	runenv.RecordMessage("started test instance, iterations=%d", iterations)
	return nil
}
