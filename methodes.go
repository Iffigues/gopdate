package main

import (
	"golang.org/x/net/html"
	"log"
	"fmt"
	"net/http"
)

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
			f := ""
			ok := false
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "toggle" {
					ok = true
				}
				if a.Key == "id" {
					f = a.Val
				}
				if  f != "" && ok {
					fmt.Println(f)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
