package main

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func (t *Manager) addVersion(n *html.Node, ff string) {
	//val := ff[2:]
	//var s version
	var f func(*html.Node)
	f = func(g *html.Node) {
		log.Println(g.Data)
		if g.Type == html.ElementNode && g.Data == "td" {
			for _, a := range g.Attr {
				log.Println(a)
			}
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
				if a.Key == "class" && a.Val == "toggle" {
					ok = true
				}
				if a.Key == "id" && a.Val == "archive" {
					break attr
				}
				if a.Key == "id" {
					ff = a.Val
				}
				if ff != "" && ok {
					t.addVersion(n, ff)
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
