package main

import (
	"os"

	"github.com/vldKasatonov/btc-indexer-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
