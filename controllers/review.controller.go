package controllers

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"gorm.io/gorm"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

var post models.Post
var txUpdatePost *gorm.DB

func ChangeRatingPostCreate(postID uint32, rating uint) (rowsAffected uint, err error) {

	post.ID = postID

	txUpdatePost = db.DB.Model(&post).Updates(map[string]interface{}{
		"Sum_ratings": gorm.Expr("Sum_ratings + ?", rating),
		"Num_ratings": gorm.Expr("Num_ratings + ?", 1)})

	if txUpdatePost.Error != nil {
		return 0, txUpdatePost.Error
	}
	return uint(txUpdatePost.RowsAffected), nil

}

func ChangeRatingPostUpdate(postID uint32, newRating, oldRating uint) (rowsAffected uint, err error) {

	post.ID = postID

	txUpdatePost = db.DB.Model(&post).Update("Sum_ratings", gorm.Expr("Sum_ratings + ? - ?", newRating, oldRating))

	if txUpdatePost.Error != nil {
		return 0, txUpdatePost.Error
	}
	return uint(txUpdatePost.RowsAffected), nil

}

func ChangeRatingPostDelete(postID uint32, rating uint) (rowsAffected uint, err error) {

	post.ID = postID

	txUpdatePost = db.DB.Model(&post).Updates(map[string]interface{}{
		"Sum_ratings": gorm.Expr("Sum_ratings - ?", rating),
		"Num_ratings": gorm.Expr("Num_ratings - ?", 1)})

	if txUpdatePost.Error != nil {
		return 0, txUpdatePost.Error
	}
	return uint(txUpdatePost.RowsAffected), nil
}

// A successful GetReviewsQueryParams() returns the
// *gorm.DB transaction result and the []reviews
// received. Otherwise, it returns nil for both
// transaction and reviews, and an error text.
func GetReviewsQueryParams(q url.Values, numReviews int) (tx *gorm.DB, reviews []models.Review, err string) {

	var txResult *gorm.DB
	var revs []models.Review

	// query params parsing
	pagNumber, pagErr := strconv.Atoi(q.Get("page"))
	if pagErr != nil && q.Get("page") != "" || pagNumber < 0 {
		return nil, nil, "Error with query param 'page' in 'GET /review' request."
	}
	postIDNumber, postErr := strconv.Atoi(q.Get("postID"))
	if postErr != nil && q.Get("postID") != "" || postIDNumber < 0 {
		return nil, nil, "Error with query param 'postID' in 'GET /review' request."
	}

	// any query param given?
	if q.Get("page") == "" && q.Get("postID") == "" {
		// no given params, get first numReviews reviews
		txResult = db.DB.Limit(numReviews).Find(&revs)
	} else
	// page query param given diff to 0?
	if pagNumber != 0 && q.Get("postID") == "" {
		// page > 0, no postID query param, get n paginated reviews
		txResult = db.DB.Limit(numReviews).Offset(numReviews * (pagNumber - 1)).Find(&revs)
	} else
	// postID query param given!, is different to 0?
	if q.Get("page") == "" && postIDNumber != 0 {
		// postID > 0, no page query param, get reviews from post
		txResult = db.DB.Where("post_id = ?", q.Get("postID")).Find(&revs)
	} else
	// any other case must throw error
	{
		return nil, nil, "Error sending 'GET /review' request."
	}
	return txResult, revs, ""
}

// ValidateReviewData take a review and validate
// each of his attributes. If all are correct
// then returns true, otherwise returns false
// and an error text.
func ValidateReviewData(review models.Review) (valid bool, textErr string) {

	// PostID      uint32
	if review.PostID < 1 {
		return false, fmt.Sprintf("Error validating PostID field, '%d' invalid value", review.PostID)
	}
	// User_name   string
	usernameRegexp, err := regexp.Compile(`^[[:ascii:]]+$`) // CAMBIAR EL VALOR MINIMO DEL STRING
	if err != nil || !usernameRegexp.MatchString(review.User_name) {
		return false, fmt.Sprintf("Error validating User_name field, '%s' invalid value", review.User_name)
	}
	// User_email  string
	emailRegexp, err := regexp.Compile(`^[[:ascii:]]+@[a-z]+.[a-z]+$`) // VER CONVENCIONES PARA EMAILS
	if err != nil || !emailRegexp.MatchString(review.User_email) {
		return false, fmt.Sprintf("Error validating User_email field, '%s' invalid value", review.User_email)
	}
	// Rating      uint
	if review.Rating <= 0 || review.Rating > 5 {
		return false, fmt.Sprintf("Error validating Rating field, '%d' invalid value", review.Rating)
	}
	// Review_text string
	reviewTextRegexp, err := regexp.Compile(`^[[:ascii:]]{1,1000}$`)
	if err != nil || !reviewTextRegexp.MatchString(review.Review_text) {
		return false, fmt.Sprintf("Error validating Review_text field, '%s' invalid value", review.Review_text)
	}

	return true, ""
}

// Same as ValidateReviewData() but including
// oldRating in the validations of an OldReview
// given.
func ValidateOldReviewData(review models.OldReview) (valid bool, textErr string) {

	// normal review fields validation
	ValidateReviewData(
		models.Review{
			PostID:      review.PostID,
			User_name:   review.User_name,
			User_email:  review.User_email,
			Rating:      review.Rating,
			Review_text: review.Review_text,
		})
	// OldRating   uint
	if review.OldRating <= 0 || review.OldRating > 5 {
		return false, fmt.Sprintf("Error validating OldRating field, '%d' invalid value", review.OldRating)
	}
	return true, ""
}

// ValidateRequestBodyData take a RequestBody
// and validate its PostID and OldRating
// attributes. If all are correct then returns
// true, otherwise returns false and an error text.
func ValidateRequestBodyData(body models.RequestBody) (valid bool, textErr string) {

	// PostID    	uint32
	if body.PostID < 1 {
		return false, fmt.Sprintf("Error validating PostID field, '%d' invalid value", body.PostID)
	}
	// OldRating 	uint
	if body.OldRating <= 0 || body.OldRating > 5 {
		return false, fmt.Sprintf("Error validating OldRating field, '%d' invalid value", body.OldRating)
	}

	return true, ""
}
