package routes

import (
	"encoding/json"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	controller "github.com/SourceLambda/sourcelambda_post_ms/controllers"
	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	var posts []models.Post

	// return only 20 posts for paginate
	numPosts := 20

	var tx *gorm.DB

	q := r.URL.Query()
	pagNumber, err := strconv.Atoi(q.Get("p"))
	/*
		case 1: pag == 1, ?p=1 or p not gived:
			// this can create an error in strconv.Atoi(),
			// that why it's first case
			db.Limit().Find(posts)
		case 2: error, ?p=invalid value:
			w.Write(error)
		case 3: pag == some num, ?p=number>1:
			db.Limit().Offset().Find(posts)
	*/
	if q.Get("p") == "1" || q.Get("p") == "" {
		tx = db.DB.Limit(numPosts).Find(&posts)
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		tx = db.DB.Limit(numPosts).Offset(numPosts * (pagNumber - 1)).Find(&posts)
	}

	if tx.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&posts)
	}

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
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(&post)
	}

}

// this is the POST method for Post entity
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	// we store the form-data text values in this map
	postValuesMap := make(map[string]string)

	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// mr return all the parts in the r.Body
	// params["boundary"] has 26 hyphen set and a big number
	// (in postman; that value isn't equal to Content-Type boundary value)
	mr := multipart.NewReader(r.Body, params["boundary"])

	// the filename of both image and html files, uuid names
	descName, imgName := controller.MultipartHandler(mr, postValuesMap)

	// quering the url/BUCKET_NAME/o/{FOLDER/FILENAME}?alg=media
	// get the image (with additional settings).
	imagePath := fmt.Sprintf("%s/o/%s", os.Getenv("BUCKET_NAME"), imgName)
	descPath := fmt.Sprintf("%s/o/%s", os.Getenv("BUCKET_NAME"), descName)

	unitsInt, _ := strconv.Atoi(postValuesMap["Units"])
	categoryIDInt, _ := strconv.Atoi(postValuesMap["CategoryID"])

	post := models.Post{
		Title:         postValuesMap["Title"],
		CategoryID:    uint32(categoryIDInt),
		Image:         imagePath,
		Description:   descPath,
		Creation_date: postValuesMap["Creation_date"],
		Units:         uint(unitsInt),
		Price:         postValuesMap["Price"],
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

	// we store the form-data text values in this map
	postValuesMap := make(map[string]string)

	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// same as in CreatePostHandler()

	mr := multipart.NewReader(r.Body, params["boundary"])

	descName, imgName := controller.MultipartHandler(mr, postValuesMap)

	// quering the url/BUCKET_NAME/o/{FOLDER/FILENAME}?alg=media
	// get the image (with additional settings).
	imagePath := fmt.Sprintf("%s/o/%s", os.Getenv("BUCKET_NAME"), imgName)
	descPath := fmt.Sprintf("%s/o/%s", os.Getenv("BUCKET_NAME"), descName)

	unitsInt, _ := strconv.Atoi(postValuesMap["Units"])
	categoryIDInt, _ := strconv.Atoi(postValuesMap["CategoryID"])

	post := models.Post{
		Title:         postValuesMap["Title"],
		CategoryID:    uint32(categoryIDInt),
		Image:         imagePath,
		Description:   descPath,
		Creation_date: postValuesMap["Creation_date"],
		Units:         uint(unitsInt),
		Price:         postValuesMap["Price"],
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

	tx := db.DB.Delete(&models.Post{}, postID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(tx.Error.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Post successfully deleted. %d Rows affected", tx.RowsAffected)))
	}

}
