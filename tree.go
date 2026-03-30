package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCurrDir() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		panic(err)
	}
	return path
}

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
