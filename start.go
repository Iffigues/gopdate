package main

func (c *Manager) help() {
}

func (c *Manager) start() {
	f := map[string]func(){
		"help": c.help,
	}
	f[c.opt.command]()
}
