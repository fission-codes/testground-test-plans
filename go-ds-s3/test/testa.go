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
//   --test-param cpuprof_path=/tmp/cpu.prof \
//   --test-param memprof_path=/tmp/mem.prof
//

// Testa
func Testa(runenv *runtime.RunEnv) error {
	runenv.RecordMessage("started test instance")
	return nil
}
