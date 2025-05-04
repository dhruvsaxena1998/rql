package main

import (
	"fmt"
	"os"

	"github.com/dhruvsaxena1998/rel/cmd/cli"
)

func main() {
	if err := cli.RootCommand.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
