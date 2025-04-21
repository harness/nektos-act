package main

import (
	_ "embed"

	"github.com/harness/nektos-act/CI-17031-test/cmd"
	"github.com/harness/nektos-act/CI-17031-test/pkg/common"
)

//go:embed VERSION
var version string

func main() {
	ctx, cancel := common.CreateGracefulJobCancellationContext()
	defer cancel()

	// run the command
	cmd.Execute(ctx, version)
}
