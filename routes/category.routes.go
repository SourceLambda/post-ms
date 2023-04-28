package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {

	var category models.Category
	vars := mux.Vars(r)
	categoryID := vars["id"]

	tx := db.DB.First(&category, categoryID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&category)

}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {

	var categories []models.Category

	tx := db.DB.Find(&categories)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&categories)
	}

}
