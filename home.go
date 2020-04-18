package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func (t *Manager)getHome() {
	e := t.home + "/.gopdate"
	if _, err := os.Stat(e); os.IsNotExist(err) {
		os.Mkdir(e, 0777)
	}
	g := e+"/package.json"
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
	err = json.Unmarshal(byteValue, &t.downloadVersion)
	if err != nil {
    	log.Fatal(t.downloadVersion)
	}
}