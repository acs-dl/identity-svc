package main

import (
	"os"

	"gitlab.com/distributed_lab/acs/identity-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
