package main

import (
	"go/build"
)

type HtmlCollector struct {
	class   string
	archive string
}

type Manager struct {
	goOs         string
	goArch       string
	goPath       string
	urlDownloads string
	url          string
	goVersion    map[int][]string
	html         HtmlCollector
}

func NewManager(goos string, goarch string) (t *Manager) {
	t = &Manager{
		goOs:         goos,
		goArch:       goarch,
		urlDownloads: "https://dl.google.com/go/",
		url:          "https://golang.org/dl/",
		goPath:       build.Default.GOPATH,
		html: HtmlCollector{
			class:   "toggle",
			archive: "archive",
		},
		goVersion: make(map[int][]string),
	}
	t.getVersion()
	return
}
