package main

import "fmt"

func main() {
	opts := newOPT()
	e := NewManager(opts)
	//fmt.Println("\n", e.goArch, e.goOs)
	fmt.Print(e.find())
	//fmt.Println("\n", e.goArch, e.goOs)
}
