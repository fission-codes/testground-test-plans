package main

import (
	test "github.com/fission-suite/testground-test-plans/go-ds-s3/test"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

func main() {
	run.Invoke(runner)
}

func runner(runenv *runtime.RunEnv) error {
	switch c := runenv.TestCase; c {
	case "add_dir":
		return test.AddDir(runenv)
	case "testa":
		return test.Testa(runenv)
	case "testb":
		return test.Testb(runenv)
	case "testc":
		return test.Testc(runenv)
	default:
		panic("unrecognized test case")
	}
}
