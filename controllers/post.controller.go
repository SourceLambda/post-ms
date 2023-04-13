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

// A successful GetPostsByQueryParams() returns the
// *gorm.DB transaction result and the []reviews
// received. Otherwise, it returns nil for both
// transaction and reviews, and an error text.
func GetPostsByQueryParams(q url.Values, numPosts int) (tx *gorm.DB, posts []models.Post, err string) {

	var txResult *gorm.DB
	var postsToReturn []models.Post

	// query params parsing
	pagNumber, pagErr := strconv.Atoi(q.Get("page"))
	
	// strconv.Atoi() return 0 to the first value if
	// q.Get() param is not given, so if ?page=0 must
	// not be valid but if page isn't given (a valid)
	// case, that can also return an error. That's 
	// why there are many conditions below.
	if q.Get("page") != "" && (pagErr != nil || pagNumber == 0) || (pagNumber < 0) {
		return nil, nil, "Error with query param 'page' in 'GET /post' request."
	}

	categoryNumber, catErr := strconv.Atoi(q.Get("category"))
	if q.Get("category") != "" && (catErr != nil || categoryNumber == 0) || (categoryNumber < 0) {
		return nil, nil, "Error with query param 'category' in 'GET /post' request."
	}

	// both query params given?
	if (q.Get("page") != "" && q.Get("category") != "") {
		// both params given! get numPosts posts with paginate and where category = some
		txResult = db.DB.Where("category_id = ?", q.Get("category")).Limit(numPosts).Offset(numPosts * (pagNumber - 1)).Find(&postsToReturn)
	} else 
	// is the category param the only one given?
	if (q.Get("page") == "" && q.Get("category") != "") {
		// category param given, get first numPosts posts where category = some
		txResult = db.DB.Where("category_id = ?", q.Get("category")).Limit(numPosts).Find(&postsToReturn)
	} else 
	// no category param given, page param given?
	if (q.Get("page") != "" && q.Get("category") == "") {
		// get numPosts with paginate
		txResult = db.DB.Limit(numPosts).Offset(numPosts * (pagNumber - 1)).Find(&postsToReturn)
	} else 
	// no given params
	{
		// get first numPosts
		txResult = db.DB.Limit(numPosts).Find(&postsToReturn)
	}
	return txResult, postsToReturn, ""
}

// ValidatePostData take a post and validate
// each of his attributes. If all are correct
// then returns true, otherwise returns false
// and an error text.
func ValidatePostData(post models.Post) (valid bool, textErr string) {

	// Title         string
	titleRegexp, err := regexp.Compile(`^[[:ascii:]]+$`)
	if err != nil || !titleRegexp.MatchString(post.Title) {
		return false, fmt.Sprintf("Error validating Title field, '%s' invalid value", post.Title)
	}
	// CategoryID    uint32
	if post.CategoryID < 1 {
		return false, fmt.Sprintf("Error validating CategoryID field, '%d' invalid value", post.CategoryID)
	}
	// Image         string
	imageRegexp, err := regexp.Compile(`^https:\/\/firebasestorage\.googleapis\.com\/v0\/b\/([[:ascii:]])+\/o\/images%2F([a-z0-9\-]+)\.(webp|jpeg|png)\?alt=media&token=([a-z0-9\-]+)$`)
	if err != nil || !imageRegexp.MatchString(post.Image) {
		return false, fmt.Sprintf("Error validating Image field, '%s' invalid value", post.Image)
	}
	// Description   string
	//esunjsoncreanme
	// Creation_date string
	dateRegexp, err := regexp.Compile(`^(20[0-9]{2})-((1[0-2])|(0[1-9]))-((3[01])|([12][0-9])|(0[1-9]))$`)
	if err != nil || !dateRegexp.MatchString(post.Creation_date) {
		return false, fmt.Sprintf("Error validating Creation_date field, '%s' invalid value", post.Creation_date)
	}
	// Units         uint
	if post.Units < 1 {
		return false, fmt.Sprintf("Error validating Units field, '%d' invalid value", post.Units)
	}
	// Price         float32
	if post.Price < 0.0 {
		return false, fmt.Sprintf("Error validating Price field, '%f' invalid value", post.Price)
	}

	return true, ""
}
