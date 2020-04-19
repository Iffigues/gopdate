package main

import "fmt"

func main() {
	opts := NewOPT()
	e := NewManager(opts)
	fmt.Print(e.Find())
}
