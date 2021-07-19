package main

import (
	"go/build"
	"log"
	"os/user"
	"runtime"
	"strings"
)

// HTMLCollector is here
type HTMLCollector struct {
	class   string
	archive string
}

type opt struct {
	command string
}

type version struct {
	Feature bool   `json:download`
	Archived bool   `json:"archived"`
	Unstable bool   `json:"unstable"`
	Version  string `json:"version"`
	Kind     string `json:"kind"`
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Types    string `json:"types"`
	Filename string `json:"filename"`
	GOARCH   string `json:"goarch"`
	GOOS     string `kson:"goos"`
}

// Manager et the base system manager
type Manager struct {
	feat            []string
	home            string
	goOs            string
	goArch          string
	goPath          string
	goRoot          string
	urlDownloads    string
	url             string
	unstable	bool
	goVersion       []version
	downloadVersion []version
	cachedVersion   []version
	opt             *opt
	html            HTMLCollector
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

// NewManager retrun manager
func NewManager(opt *opt) (t *Manager) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	t = &Manager{
		opt:          opt,
		unstable:	false,
		home:         usr.HomeDir,
		goOs:         runtime.GOOS,
		goArch:       runtime.GOARCH,
		urlDownloads: "https://dl.google.com/go/",
		url:          "https://golang.org/dl/",
		goPath:       build.Default.GOPATH,
		goRoot:       build.Default.GOROOT,
		html: HTMLCollector{
			class:   "toggle",
			archive: "archive",
		},
	}
	t.getHome()
	t.getVersion()
	t.getGo()
	t.cachedVersion = t.goVersion
	t.start()
	return
}
