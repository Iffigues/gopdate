package main

import (
	"net/http"
)

func (t *Manager) upload(s version) (err error) {
	resp, err := http.Get(t.urlDownloads + s.Filename)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return
}
