package main

import "fmt"

func main() {
	opts := newOPT()
	e := NewManager(opts)
	fmt.Print(e.find())
}
