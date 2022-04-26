package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		filename := strings.TrimLeft(req.RequestURI, "/")
		if filepath.Ext(filename) == ".br" {
			resp.Header().Set("content-encoding", "br")
			resp.Header().Set("Content-Type", "application/wasm")
		}
		content, _ := ioutil.ReadFile(filename)
		resp.Write(content)
	})
	fmt.Println("server is running at :9091...")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println(err)
	}
}
