package main

import (
	"archive/zip"
	//"archive/tar"
	//"compress/gzip"
)

// UnCompressor return interface connector
type UnCompressor interface {
	UnCompress(a, b string) error
}

// UnZapper is for untar
type UnZapper struct {

}

// UnZipper is for zip
type UnZipper struct {
	
}

// UnCompress is d
func (u *UnZapper) UnCompress(src, target string) error {
	
	return nil
}

// UnCompress is f
func (u *UnZipper) UnCompress(src, target string) error {
	zip.OpenReader("")
	return nil
}

func (s *version) getUnCompressor() (e UnCompressor) {
	return
}

// Uncompress isf
func (c *Manager) Uncompress(s *version) {

}
