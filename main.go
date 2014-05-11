package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://quora.com" + r.URL.Path + "?share=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		panic(err)
	}
}
