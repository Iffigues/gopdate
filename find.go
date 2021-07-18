package main

import "regexp"
func (t *Manager) find() (s []version) {
	for _, v := range t.goVersion {
		e,_:= regexp.MatchString(t.goArch, v.Filename)
		if e && v.Os == t.goOs {
			s = append(s, v)
		}
	}

	return
}
