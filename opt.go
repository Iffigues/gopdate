package main

import "os"

func commander(b string) (a string) {
	f := []string{"help"}
	for _, g := range f {
		if b == g {
			return b
		}
	}
	return "help"
}

func newOPT() (opts *opt) {
	opts = &opt{}
	opts.command = "help"
	for i, val := range os.Args {
		if i > 0 {
			if val != "" {
				if i == 1 {
					opts.command = commander(val)
				}
			}
		}
	}
	return
}
