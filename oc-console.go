package main // import "github.com/cblecker/oc-console"

import (
	"os"

	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/cblecker/oc-console/pkg/console"
)

func main() {
	flags := pflag.NewFlagSet("oc-console", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := console.NewCmdConsoleConfig(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
