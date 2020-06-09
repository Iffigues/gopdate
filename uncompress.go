package main

// UnCompressor return interface connector
type UnCompressor interface {
	UnCompress() error
}

// UnZapper is for untar
type UnZapper struct {
}

// UnZipper is for zip
type UnZipper struct {
}

// UnCompress is d
func (u *UnZapper) UnCompress() error {
	return nil
}

// UnCompress is f
func (u *UnZipper) UnCompress() error {
	return nil
}

func (s *version) getUnCompressor() (e UnCompressor) {
	return
}

// Uncompress isf
func (c *Manager) Uncompress(s *version) {

}
