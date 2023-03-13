package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/routes"
)

// split it in addr:port
var PORT string = ":8080"

func main() {

	db.DBConnection()

	r := mux.NewRouter()
	server := http.Server{
		Addr:    PORT,
		Handler: r,
	}

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/categories", routes.GetCategoriesHandler).Methods("GET")

	r.HandleFunc("/post", routes.GetPostsHandler).Methods("GET")
	r.HandleFunc("/post/{id}", routes.GetPostHandler).Methods("GET")
	r.HandleFunc("/post", routes.CreatePostHandler).Methods("POST")
	r.HandleFunc("/post/{id}", routes.PutPostHandler).Methods("PUT")
	r.HandleFunc("/post/{id}", routes.DeletePostHandler).Methods("DELETE")

	r.HandleFunc("/review", routes.GetReviewsHandler).Methods("GET")
	r.HandleFunc("/review/{id}", routes.GetReviewHandler).Methods("GET")
	r.HandleFunc("/review", routes.CreateReviewHandler).Methods("POST")
	r.HandleFunc("/review/{id}", routes.PutReviewHandler).Methods("PUT")
	r.HandleFunc("/review/{id}", routes.DeleteReviewHandler).Methods("DELETE")

	log.Fatal(server.ListenAndServe())
}
