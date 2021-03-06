package main

func (t *Manager) find() (s []version) {
	for _, v := range t.goVersion {
		if v.Arch == t.goArch && v.Os == t.goOs {
			s = append(s, v)
		}
	}

	return
}
