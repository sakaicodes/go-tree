package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetCurrDir returns the current working directory as a string. If there is an error while getting the current directory, it prints the error and panics.
func GetCurrDir() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		panic(err)
	}
	return path
}

// PrintTree recursively prints the directory structure starting from startDir.
//
// Parameters:
// - startDir: The directory from which to start printing the tree.
// - prefix: A string used to format the output, indicating the level of the tree.
// - skipDirs: A map of directory names to skip when printing the tree.
func PrintTree(startDir string, prefix string, skipDirs map[string]bool) {
	entries, err := os.ReadDir(startDir)
	if err != nil {
		panic(err)
	}
	for i, entry := range entries {
		if skipDirs[entry.Name()] {
			continue
		}
		isLast := i == len(entries)-1
		connector := "├── "
		if isLast {
			connector = "└── "
		}
		fmt.Println(prefix + connector + entry.Name())

		if entry.IsDir() {
			extension := "│   "
			if isLast {
				extension = "    "
			}
			PrintTree(filepath.Join(startDir, entry.Name()), prefix+extension, skipDirs)
		}

	}
}
