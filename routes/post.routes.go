package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"

	"github.com/SourceLambda/sourcelambda_post_ms/controllers"
	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	// return only 20 posts for paginate
	numPosts := 20

	q := r.URL.Query()

	tx, posts, err := controllers.GetPostsByQueryParams(q, numPosts)
	if err != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err))
		return
	}
	// tx can return an error even if err == nil.
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&posts)

}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	vars := mux.Vars(r)
	postID := vars["id"]
	tx := db.DB.First(&post, postID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&post)
	}

}

func GetCountPostsHandler(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Count   int
	}
	  
	var result Result

	tx := db.DB.Raw("SELECT COUNT(*) FROM post").Scan(&result)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(tx.Error.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", result.Count)))
	//json.NewEncoder(w).Encode(&result)

}

// this is the POST method for Post entity
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	// post data validation
	isValid, err := controllers.ValidatePostData(post)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err))
		return
	}

	tx := db.DB.Create(&post)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Post successfully created."))
	}
}

func PutPostHandler(w http.ResponseWriter, r *http.Request) {

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	// post data validation
	isValid, validationErr := controllers.ValidatePostData(post)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validationErr))
		return
	}

	// Get the id path variable in /post/{id}
	// to set the postID to updating its values.
	vars := mux.Vars(r)

	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to convert string postID to integer."))
	} else {
		post.ID = uint32(postID)

		tx := db.DB.Save(&post)
		if tx.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(tx.Error.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Post successfully edited. %d Rows affected", tx.RowsAffected)))
		}

	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	postID := vars["id"]

	var reviews []models.Review
	deleteRevsTX := db.DB.Clauses(clause.Returning{}).Where("post_id = ?", postID).Delete(&reviews)
	if deleteRevsTX.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(deleteRevsTX.Error.Error()))
		return
	}
	w.Write([]byte(fmt.Sprintf("%d Post's reviews successfylly deleted.\n", len(reviews))))

	tx := db.DB.Delete(&models.Post{}, postID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Post successfully deleted. %d Rows affected", tx.RowsAffected)))
	}

}
