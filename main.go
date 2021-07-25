package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("file/zorro-file/"))
	http.Handle("/file/", http.StripPrefix("/file/", fs))
	http.HandleFunc("/", serve)

	fmt.Println("starting donkey")
	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)

	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")
}
