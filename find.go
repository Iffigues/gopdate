package main

import "regexp"
import "os"
import "bytes"
func (t *Manager) find() (s []version) {
	for _, v := range t.goVersion {
		e,_:= regexp.MatchString(t.goArch, v.Filename)
		if e && v.Os == t.goOs {
			s = append(s, v)
		}
	}

	return
}

func (t *Manager) Show(a []version) {
	var buff bytes.Buffer

	for _, b := range a {
		if (b.Archived) {
			buff.Write([]byte("archived"))
		}
		if b.Unstable {
			buff.Write([]byte("unstable"))
		} else {
		}
		if b.Feature {
			buff.Write([]byte("edowdsqdsqdsqdsqdsqdsqdsqdsdsqdsqdsqdsqdsqdsqdsqdqssdsqds"))
		}
		buff.Write([]byte("\n"))
		buff.WriteTo(os.Stdout)	
		buff.Reset()
	}
	print(t.goOs, t.goArch)
}
