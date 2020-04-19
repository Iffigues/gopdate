package main

import (
	"go/build"
	"log"
	"os/user"
	"runtime"
	"strings"
)

type HtmlCollector struct {
	class   string
	archive string
}

type opt struct {
}

type version struct {
	Archived bool   `json:"archived"`
	Version  string `json:"version"`
	Kind     string `json:"kind"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Types    string `json:"types"`
	Filename string `json:"filename"`
	GOARCH   string `json:"goarch"`
	GOOS     string `kson:"goos"`
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
	opt             *opt
	html            HtmlCollector
}

func (t *Manager) getGo() {
	for i, val := range t.goVersion {
		if val.Arch == "bootstrap" {
			t.goVersion[i].GOARCH = val.Arch
			t.goVersion[i].GOOS = val.Os
		} else if val.Kind != "Source" {
			z := strings.Split(val.Filename, ".")
			for _, vals := range z {
				if strings.Contains(vals, "-") {
					zz := strings.Split(vals, "-")
					t.goVersion[i].GOARCH = zz[1]
					t.goVersion[i].GOOS = zz[0]
				}
			}
		}
	}
}

func NewManager(opt *opt) (t *Manager) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	t = &Manager{
		opt:          opt,
		home:         usr.HomeDir,
		goOs:         runtime.GOOS,
		goArch:       runtime.GOARCH,
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
	t.getGo()
	return
}
