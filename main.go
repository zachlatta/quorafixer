package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/quorafixer/view"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/{path:.*}", quoraHandler)
	http.Handle("/", r)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	view.RenderFile(w, r, "index.html")
}

func quoraHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]

	if view.Exists(path) {
		view.RenderFile(w, r, path)
		return
	}

	resp, err := http.Get("https://quora.com/" + path + "?share=1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		panic(err)
	}
}
