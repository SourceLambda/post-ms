package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	var posts []models.Post
	// futuro paginado (pagina n-esima) 
	n := 0
	
	// returns a tx *gorm.DB object
	db.DB.Limit(10).Offset(10*n).Find(&posts)

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&posts)

}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	vars := mux.Vars(r)
	postID := vars["id"]
	db.DB.First(&post, postID)

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&post)

}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) { // this is the POST method for Post entity

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	db.DB.Create(&post)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post successfully created."))
}

func PutPostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to convert string postID to integer."))
	} else {
		post.ID = uint32(postID)
	
		db.DB.Save(&post)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Post successfully edited"))
	}

}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	postID := vars["id"]

	db.DB.Delete(&models.Post{}, postID)
	
	w.WriteHeader(http.StatusContinue)
	w.Write([]byte("Post successfully deleted."))
	
}
