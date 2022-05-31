package test


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




import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

        "github.com/testground/sdk-go/runtime"
	"github.com/ipfs/go-ipfs/plugin/loader"
)

var flagExp = flag.Bool("experimental", false, "enable experimental features")

// Testa
func AddDir(runenv *runtime.RunEnv) error {
	var (
		iterations = runenv.IntParam("iterations")
	)
	runenv.RecordMessage("started test instance, iterations=%d", iterations)

	flag.Parse()

	runenv.RecordMessage("all done")
}

