package cmd

import (
	"flag"

	"github.com/sakaicodes/tree/internal"
)

// DisplayTree is the main function that handles the logic for displaying the directory tree. It parses command-line flags and arguments, determines the directory to display, and calls the PrintTree function from the internal package to print the directory structure.
func DisplayTree() {
	all := flag.Bool("all", false, "Include all files and directories (including hidden ones)")
	flag.Parse()

	var dir string
	if args := flag.Args(); len(args) > 0 {
		if args[0] == "." {
			dir = internal.GetCurrDir()
		} else {
			dir = args[0]
		}
	}

	if *all {
		internal.PrintTree(dir, "", map[string]bool{})
	} else {
		internal.PrintTree(dir, "", internal.DefaultSkipDirs)
	}

}
