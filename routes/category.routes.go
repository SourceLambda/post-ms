package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

	var categories []models.Category
	
	db.DB.Find(&categories)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&categories)
}
