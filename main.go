package main

func main() {
	opts := newOPT()
	e := NewManager(opts)
	e.Show(e.find())
}
