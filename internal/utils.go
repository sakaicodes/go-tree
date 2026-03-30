package internal

var DefaultSkipDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	".DS_Store":    true,
}
