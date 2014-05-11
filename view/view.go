package view

import "net/http"

// File directory
const fD = "view"

var files = []string{
	"index.html",
	"css/quorafixer.css",
	"favicon.ico",
}

func RenderFile(w http.ResponseWriter, r *http.Request, file string) {
	http.ServeFile(w, r, fD+"/"+file)
}

func Exists(file string) bool {
	for _, f := range files {
		if f == file {
			return true
		}
	}

	return false
}
