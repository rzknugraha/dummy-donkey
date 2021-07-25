package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("file"))))
	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
