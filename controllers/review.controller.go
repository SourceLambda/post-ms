package controllers

import (
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

	txUpdatePost = db.DB.Model(&post).Update("Sum_ratings", gorm.Expr("Sum_ratings + ?", newRating-oldRating))

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
