package main

import (
	"io"
	"net/http"
	"os"
)

func (t *Manager) upload(s *version) (err error) {
	resp, err := http.Get(t.urlDownloads + s.Filename)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(t.home + "/.gopdate/version/" + s.Filename)

	if err != nil {
		return err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return
}
