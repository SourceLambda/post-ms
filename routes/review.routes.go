package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {

	var reviews []models.Review
	// return only 20 reviews for paginate
	numReviews := 20

	var tx *gorm.DB

	q := r.URL.Query()
	pagNumber, err := strconv.Atoi(q.Get("p"))
	/*
		case 1: pag == 1, ?p=1 or p not gived:
			db.Limit().Find(posts)
		case 2: error, ?p=invalid value:
			w.Write(error)
		case 3: pag == some num, ?p=number>1:
			db.Limit().Offset().Find(posts)
	*/
	if q.Get("p") == "1" || q.Get("p") == "" {
		tx = db.DB.Limit(numReviews).Find(&reviews)
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		tx = db.DB.Limit(numReviews).Offset(numReviews * (pagNumber - 1)).Find(&reviews)
	}

	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&reviews)
	}
}

func GetReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	vars := mux.Vars(r)
	reviewID := vars["id"]
	tx := db.DB.First(&review, reviewID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(&review)
	}

}

func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	json.NewDecoder(r.Body).Decode(&review)

	tx := db.DB.Create(&review)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Post successfully created"))
	}

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

		tx := db.DB.Save(&review)
		if tx.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(tx.Error.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Review successfully edited. %d Rows affected.", tx.RowsAffected)))
		}

	}

}

func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	reviewID := vars["id"]

	tx := db.DB.Delete(&models.Review{}, reviewID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Review successfully deleted. %d Rows affected.", tx.RowsAffected)))
	}

}
