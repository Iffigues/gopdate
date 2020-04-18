package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func getFilename(n *html.Node) (a string) {
	b := n.LastChild
	if b != nil {
		a = b.Data
		b = b.LastChild
	}
	if b != nil {
		a = b.Data
	}
	return a
}

func (t *Manager) getiingMe(n *html.Node, ff string, ok bool) {
	var f func(*html.Node)
	var s version
	s.archived = ok
	step := 0
	f = func(g *html.Node) {
		if g.Type == html.ElementNode && g.Data == "td" {
			step = step + 1
			if step == 1 {
				s.filename = getFilename(g)
			}
			if step == 2 {
				s.kind = getFilename(g)
			}
			if step == 3 {
				s.os = getFilename(g)
			}
			if step == 4 {
				s.arch = getFilename(g)
			}
		}
		for c := g.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	if step > 0 {
		t.goVersion = append(t.goVersion, s)
	}
	fmt.Println(t.goVersion)
}

func (t *Manager) addVersion(n *html.Node, ff string, ok bool) {
	//val := ff[2:]
	//var s version
	var f func(*html.Node)
	f = func(g *html.Node) {
		if g.Type == html.ElementNode && g.Data == "tr" {
			t.getiingMe(g, ff, ok)
		}
		for c := g.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
}

func (t *Manager) getVersion() {

	resp, err := http.Get(t.url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			ff := ""
			ok := false
		attr:
			for _, a := range n.Attr {
				oks := false
				if a.Key == "class" && a.Val == "toggle" {
					ok = true
				}
				if a.Key == "id" && a.Val == "archive" {
					oks = true
				}
				if a.Key == "id" {
					ff = a.Val
				}
				if ff != "" && ok {
					t.addVersion(n, ff, oks)
					break attr
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
