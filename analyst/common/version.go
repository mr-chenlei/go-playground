package common

import "fmt"

// no need to set value to this. makefile will auto-inject value to this
var (
	version = "None"
	build   = "None"
)

// Version get current version
func Version() string {
	return version
}

// Build get current build
func Build() string {
	return build
}

// PrintVersion ...
func PrintVersion() {
	fmt.Println("version:", version)
}

// PrintBuild ...
func PrintBuild() {
	fmt.Println("build:", build)
}
