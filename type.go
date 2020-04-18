package main

import (
	"go/build"
	"log"
	"os/user"
)

type HtmlCollector struct {
	class   string
	archive string
}

type version struct {
	archived bool   `json:"archived"`
	version  string `json:"version"`
	kind     string `json:"kind"`
	os       string `json:"os"`
	arch     string `json:"arch"`
	types    string `json:"types"`
	filename string `json:"filename"`
}

type Manager struct {
	home            string
	goOs            string
	goArch          string
	goPath          string
	urlDownloads    string
	url             string
	goVersion       []version
	downloadVersion []version
	html            HtmlCollector
}

func NewManager(goos string, goarch string) (t *Manager) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	t = &Manager{
		home:         usr.HomeDir,
		goOs:         goos,
		goArch:       goarch,
		urlDownloads: "https://dl.google.com/go/",
		url:          "https://golang.org/dl/",
		goPath:       build.Default.GOPATH,
		html: HtmlCollector{
			class:   "toggle",
			archive: "archive",
		},
	}
	t.getHome()
	t.getVersion()
	return
}
