package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func (t *Manager) getHome() {
	e := t.home + "/.gopdate"
	if _, err := os.Stat(e); os.IsNotExist(err) {
		os.Mkdir(e, 0777)
	}
	ee := e + "/version"
	if _, err := os.Stat(ee); os.IsNotExist(err) {
		os.Mkdir(ee, 0777)
	}
	g := e + "/package.json"
	if _, err := os.Stat(g); os.IsNotExist(err) {
		return
	}
	jsonFile, err := os.Open(g)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, &t.downloadVersion)
}

func (t *Manager) WriteJSON() {
	file, _ := json.Marshal(t.downloadVersion)
	_ = ioutil.WriteFile(t.home+"/.gopdate/package.json", file, 0644)

}
