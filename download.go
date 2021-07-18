package main

import  (
	"net/http"
	"fmt"
	"os"
	"io"
)

func (c *Manager)Download() {
	resp, err := http.Get(c.url + c.goVersion[1].Filename)
	println(c.url+c.goVersion[1].Filename)
	if err != nil {
		fmt.Println(err, resp)
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("./ii")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println(err)
}
