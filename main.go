package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/routes"
)

func main() {

	SetEnvVars()

	if os.Getenv("POST_MS_PORT") == "" {
		log.Println("Not PORT ENV_VAR found, 8080 default port used.")
		os.Setenv("POST_MS_PORT", "8080")
	}
	var ADDR = fmt.Sprintf(":%s", os.Getenv("POST_MS_PORT"))

	db.DBConnection()

	r := mux.NewRouter()
	server := http.Server{
		Addr:    ADDR,
		Handler: r,
	}

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/categories", routes.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id}", routes.GetCategoryHandler).Methods("GET")
	
	r.HandleFunc("/post", routes.GetPostsHandler).Methods("GET")
	r.HandleFunc("/post/{id}", routes.GetPostHandler).Methods("GET")
	r.HandleFunc("/count-post", routes.GetCountPostsHandler).Methods("GET")
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
