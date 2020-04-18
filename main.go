package main

import (
	"runtime"
)

func main() {
	NewManager(runtime.GOOS, runtime.GOARCH)
}
