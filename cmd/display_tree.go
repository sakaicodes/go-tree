package cmd

import (
	"flag"

	"github.com/sakaicodes/tree/internal"
)

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
