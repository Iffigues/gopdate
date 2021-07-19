package main

import  (
	"net/http"
	"fmt"
	"os"
	"io"
)

func (c *Manager)Download(a string, b string) {
	resp, err := http.Get(a)
	println(c.url+c.goVersion[1].Filename)
	if err != nil {
		fmt.Println(err, resp)
	}
	defer resp.Body.Close()
	out, err := os.Create(b)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	fmt.Println(err)
}
