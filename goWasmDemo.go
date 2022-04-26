package main

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func Fetch(name string, resolve js.Value, reject js.Value) {
	resp, err := http.Get("https://httpbin.org/get?name=" + name)
	if err != nil {
		reject.Invoke(err.Error())
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			reject.Invoke(err.Error())
		}
		resolve.Invoke(string(content))
	}

}

func FetchFunc(this js.Value, args []js.Value) any {
	reject := args[len(args)-1]  // 最后一个参数为 resolve
	resolve := args[len(args)-2] // 倒数第二个参数为 reject
	go Fetch(args[0].String(), resolve, reject)
	return nil
}

func main() {
	js.Global().Set("req", js.FuncOf(FetchFunc))
	select {}
}
