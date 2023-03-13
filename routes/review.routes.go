package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
	"github.com/gorilla/mux"
)

func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {

	var reviews []models.Review

	// number of reviews to return
	n := 0

	db.DB.Limit(10).Offset(10*n).Find(&reviews)

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&reviews)
}

func GetReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	vars := mux.Vars(r)
	reviewID := vars["id"]

	// number of reviews to return
	n := 0

	db.DB.Limit(10).Offset(10*n).Find(&review, reviewID)
	
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&review)
}

func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	json.NewDecoder(r.Body).Decode(&review)

	db.DB.Create(&review)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post successfully created"))
}

func PutReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	json.NewDecoder(r.Body).Decode(&review)

	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to convert string reviewID to integer."))
	} else {
		review.ID = uint32(reviewID)
		fmt.Print(reviewID)

		db.DB.Save(&review)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Review successfully edited."))
	}

}

func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	reviewID := vars["id"]

	db.DB.Delete(&models.Review{}, reviewID)
	// tx :=, tx.Error y RowsAffected

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Review successfully deleted."))
}
