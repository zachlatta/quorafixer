package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/{path}", quoraHandler)
	http.Handle("/", r)
	http.ListenAndServe(":4000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(homepage))
}

func quoraHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]

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

const homepage = `
<!doctype html>
<html>
<head>
<title>QuoraFix - View Quora Without An Account</title>
<link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.4.2/pure-min.css">
</head>
<body>
<div class="content">
test
</div>
</body>
</html>
`
