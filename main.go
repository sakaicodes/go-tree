package main

func main() {
	defaultSkipDirs := map[string]bool{
		".git":         true,
		"node_modules": true,
		".DS_Store":    true,
	}

	//TODO: Add command line flag to skip default directories
	skip := true // Set to true to skip default directories

	skipDirs := make(map[string]bool)
	if skip {
		skipDirs = defaultSkipDirs
	}

	start_dir := GetCurrDir()

	PrintTree(start_dir, "", skipDirs)
}
