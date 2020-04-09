package main

import (
	"go/build"
)

type Manager struct {
	goOs string
	goArch string
	goPath string
	urlDownloads string
	url string
	goVersion []string
}

func NewManager(goos string, goarch string) (t *Manager) {
	t =  &Manager {
		goOs: goos,
		goArch: goarch,
		urlDownloads: "https://dl.google.com/go/",
		url: "https://golang.org/dl/",
		goPath: build.Default.GOPATH,
	}
	t.goVersion = t.getVersion()
	return
}
