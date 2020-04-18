package main

import (
	"net/http"
)

func (t *Manager) upload(s version) (err error) {
	resp, err := http.Get(t.urlDownloads + s.filename)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	return
}
