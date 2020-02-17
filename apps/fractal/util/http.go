package util

import (
	"log"
	"net/http"
)

func HandleError(w http.ResponseWriter, msg string) {
	log.Printf("%s", msg)

	SetContentType(w, PlainText{})
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Internal Server Error\n"))

}

func HandleResourceNotFoundError(w http.ResponseWriter, msg string) {
	log.Printf("%s", msg)
	http.Error(w, "404 resource not found", http.StatusNotFound)
}

func SetContentType(w http.ResponseWriter, ct ContentType) {
	w.Header().Set("Content-Type", ct.String()+"; charset=utf-8")
}
