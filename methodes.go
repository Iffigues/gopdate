package main

import (
	"log"
	"net/http"
	"golang.org/x/net/html"
	"strings"
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
	s.Archived = ok
	s.Unstable = t.unstable
	step := 0
	f = func(g *html.Node) {
		if g.Type == html.ElementNode && g.Data == "td" {
			step = step + 1
			if step == 1 {
				s.Filename = getFilename(g)
			}
			if step == 2 {
				s.Kind = getFilename(g)
			}
			if step == 3 {
				s.Os = strings.ToLower(getFilename(g))
			}
			if step == 4 {
				s.Arch = getFilename(g)
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
}


func (t *Manager) Down(n *html.Node) {
	var f func(*html.Node)
	f = func(g *html.Node) {
		if g.Type == html.ElementNode && g.Data == "span" {
			for _, a := range g.Attr {
				if a.Key == "class" && a.Val == "filename"{
				}
			}
		}
		for c := g.FirstChild; c != nil; c =c.NextSibling {
			f(c)
		}
	}
	f(n)
}

func (t *Manager) addVersion(n *html.Node, ff string, ok bool) {
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
		if n.Type == html.ElementNode && n.Data == "h3" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == "unstable" {
					t.unstable = true
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			tt := false
			var ttt string
			for _,a := range n.Attr {
				if a.Key == "class" && a.Val == "download downloadBox" {
					tt = true

				}
				if a.Key == "href" {
					ttt = a.Val
				}
			}
			if tt {
				t.feat = append(t.feat, ttt[4:])
			}
		}
		if n.Type == html.ElementNode && n.Data == "div" {
			ff := ""
			ok := false
		attr:
			for _, a := range n.Attr {
				oks := false
				if a.Key == "class"&& a.Val == "toggleVisible" && !oks{
					ok = true
				}
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
					t.unstable = false
					break attr
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	for ee,gb := range t.goVersion {
		for _, n := range t.feat {
			if gb.Filename == n {
				println("zaazl")
				t.goVersion[ee].Feature = true
			}
		}
	}
}
