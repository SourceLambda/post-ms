package routes

import (
	"encoding/json"
	"fmt"
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
	
	err := db.DB.Limit(10).Offset(10*n).Find(&posts)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(&posts)
	}

}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	vars := mux.Vars(r)

	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	ErrRecordNotFound := db.DB.First(&post, postID)
	if ErrRecordNotFound != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "ErrRecordNotFound")
	} else {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(&post)
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) { // this is the POST method for Post entity

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	result := db.DB.Create(&post)

	if result.Error != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Failed request:", result.Error)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Post successfully created")
	}
}

func PutPostHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	var post models.Post
	post.ID = uint32(postID)
	json.NewDecoder(r.Body).Decode(&post)
	
	db.DB.Save(&post)
	// validation?
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post successfully edited"))
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	db.DB.Delete(&models.Post{}, postID)
	// validation?
	w.WriteHeader(http.StatusContinue)
	w.Write([]byte("Post deleted succesfully"))
	
}
