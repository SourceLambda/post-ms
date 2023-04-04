package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SourceLambda/sourcelambda_post_ms/controllers"
	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
	"github.com/gorilla/mux"
)

func GetReviewsHandler(w http.ResponseWriter, r *http.Request) {

	// return only 20 reviews for paginate
	numReviews := 20

	q := r.URL.Query()

	tx, reviews, err := controllers.GetReviewsQueryParams(q, numReviews)
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
	json.NewEncoder(w).Encode(&reviews)

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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&review)
	}

}

func CreateReviewHandler(w http.ResponseWriter, r *http.Request) {

	var review models.Review
	json.NewDecoder(r.Body).Decode(&review)

	txCreateRev := db.DB.Create(&review)
	if txCreateRev.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(txCreateRev.Error.Error()))
		return
	}

	// updating post's rating
	_, txErr := controllers.ChangeRatingPostCreate(review.PostID, review.Rating)
	if txErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(txErr.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Review successfully created, Rating of post %d updated.", review.PostID)))

}

func PutReviewHandler(w http.ResponseWriter, r *http.Request) {

	var oldReviewInfo models.OldReview
	json.NewDecoder(r.Body).Decode(&oldReviewInfo)

	var review = models.Review{
		PostID:      oldReviewInfo.PostID,
		User_name:   oldReviewInfo.User_name,
		User_email:  oldReviewInfo.User_email,
		Rating:      oldReviewInfo.Rating,
		Review_text: oldReviewInfo.Review_text,
	}

	vars := mux.Vars(r)

	reviewID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to convert string reviewID to integer."))
		return
	}
	review.ID = uint32(reviewID)
	fmt.Print(reviewID)

	tx := db.DB.Save(&review)
	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(tx.Error.Error()))
		return

	}

	if oldReviewInfo.OldRating == review.Rating {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Review successfully edited."))
		return
	}

	// update the post's rating if there is a new rating value
	rowsAffected, txErr := controllers.ChangeRatingPostUpdate(review.PostID, review.Rating, oldReviewInfo.OldRating)
	if txErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(txErr.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Review successfully edited. %d Posts affected.", rowsAffected)))

}

func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	reviewID := vars["id"]

	var body models.RequestBody
	json.NewDecoder(r.Body).Decode(&body)

	tx := db.DB.Delete(&models.Review{}, reviewID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
		return
	}

	rowsAffected, txErr := controllers.ChangeRatingPostDelete(body.PostID, body.OldRating)
	if txErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(txErr.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Review successfully deleted. %d Rows affected.", rowsAffected)))

}
