package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("file"))))
	fmt.Println("starting donkey")
	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
}
