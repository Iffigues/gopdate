package main

import (
	"os"
	"fmt"
	"log"
	"runtime"
	"path/filepath"
)

func main() {
	NewManager(runtime.GOOS, runtime.GOARCH)
dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(dir)
	ex, err := os.Executable()
	fmt.Println(ex)
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}