package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/routes"
)

var PORT string = ":8080"

func main() {

	db.DBConnection()

	/* deuda tecnica xd
	- ESTRUCTURA src/ DE PROYECTOS EN GO, CONFIG GOPATH
	- .env
	- recibir n en getPost (y Reviews)
	- crear func para los métodos usados en las rutas (no repetir código)
	- editing/deleting validation?
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

	log.Fatal(server.ListenAndServe())
}

// http.HandleFunc("/", HomeHandler)
// http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		routes.GetPostsHandler(w, r)
// 	case http.MethodPost:
// 		routes.CreatePostHandler(w, r)
// 	case http.MethodPut:
// 		routes.PutPostHandler(w, r)
// 	case http.MethodDelete:
// 		routes.DeletePostHandler(w, r)
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		w.Write([]byte("Metodo no aceptado"))
// 	}
// })
// http.HandleFunc("/post/{id}", func(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		routes.GetPostHandler(w, r)
// 	}
// })

// curl -H "Content-Type:application/json" -X POST -d '{"cityName":"NewYork", "countryCode":"IT", "postalCode": "00166"}' http://localhost:8080/ComuneUtenti/api/city
