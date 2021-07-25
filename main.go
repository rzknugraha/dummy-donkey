package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// func main() {
// 	fs := http.FileServer(http.Dir("file/zorro-file/"))
// 	http.Handle("/file/", http.StripPrefix("/file/", fs))

// 	http.HandleFunc("/", serve)

// 	fmt.Println("starting donkey")
// 	if err := http.ListenAndServe(":8070", nil); err != nil {
// 		log.Fatal("ListenAndServe: ", err)

// 	}
// }

func main() {
	router := NewRouter()

	handler := cors.Default().Handler(router)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug:          true,
		AllowedHeaders: []string{"*"},
	})
	handler = c.Handler(handler)

	fmt.Printf("donkey start")
	if err := http.ListenAndServe(":8070", handler); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

//NewRouter si a
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/file/"

	// Create the route
	router.
		PathPrefix("/").
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("file/zorro-file/"))))

	return router
}
