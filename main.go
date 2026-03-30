package main

import (
	"os"

	"github.com/sakaicodes/tree/cmd"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: go-tree [directory]")
		os.Exit(1)
	}
	cmd.DisplayTree()
}
