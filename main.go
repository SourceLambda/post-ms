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

	/* Deuda tecnica xd
	- recibir n en getPost (y Reviews)
	- validación crud ops post y review //https://github.com/go-gorm/gorm/blob/v1.24.6/finisher_api.go#L161
	- crear func para los métodos usados en las rutas (no repetir código)
	- verificar cambio en NamingStrategy y Save
	- put/patch?
	- comentar todo y mejorar respuestas
	*/

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
