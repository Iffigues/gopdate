package main

import "os"

func NewOPT() (opts *opt) {
	opts = &opt{}
	for i, val := range os.Args {
		if i > 0 {
			if val == "" {

			}
		}
	}
	return
}