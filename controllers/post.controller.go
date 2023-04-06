package controllers

import (
	"fmt"
	"regexp"

	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

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
	imageRegexp, err := regexp.Compile(`^https:\/\/firebasestorage\.googleapis\.com\/v0\/b\/([[:ascii:]])+\/o\/images%2F([a-z0-9\-]+)\.(webp|jpg|png)\?alt=media&token=([a-z0-9\-]+)$`)
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
