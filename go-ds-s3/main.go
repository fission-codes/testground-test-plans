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
	default:
		panic("unrecognized test case")
	}
}
