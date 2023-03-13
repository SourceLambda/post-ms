package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/SourceLambda/sourcelambda_post_ms/models"
)

// var storageHost = "http://localhost:8080"

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	
	var posts []models.Post
	// futuro paginado (pagina n-esima) 
	n := 0
	
	// returns a tx *gorm.DB object
	db.DB.Limit(10).Offset(10*n).Find(&posts)
	
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&posts)
	
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	
	var post models.Post
	vars := mux.Vars(r)
	postID := vars["id"]
	db.DB.First(&post, postID)
	
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(&post)
	
}

// this is the POST method for Post entity
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	// we store the form-data text values in this map
	postValuesMap := make(map[string]string)

	// the filename of both image and html files, with its extensions
	var dscpName string
	var imgName string

	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	// mr return all the parts in the r.Body
	// params["boundary"] has 26 hyphen set and a big number 
	// (in postman; that value isn't equal to Content-Type boundary value)
	mr := multipart.NewReader(r.Body, params["boundary"])
	
	for {
		// each part is a form-data field, files too
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		slurp, err := io.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}

		if p.FileName() != "" {
			// first case, description html file, CORRECT THE NESTED IFs
			var fileName string
			if path.Ext(p.FileName()) == ".html" {
				dscpName = uuid.New().String() + path.Ext(p.FileName())
				fileName = dscpName
				
			} else {
				imgName = uuid.New().String() + path.Ext(p.FileName())
				fileName = imgName
			}
			file, _ := os.Create(fmt.Sprintf("files/%s", fileName))
			file.Write(slurp)
		} else {
			postValuesMap[p.FormName()] = string(slurp)
			fmt.Printf("\nPart %q: %q", p.FormName(), slurp)
		}
	}

	// later it will be a url, for now it's only a local path
	imagePath := fmt.Sprintf("files/%s", imgName)
	descPath := fmt.Sprintf("files/%s", dscpName)

	unitsInt, _ := strconv.Atoi(postValuesMap["Units"])
	categoryIDInt, _ := strconv.Atoi(postValuesMap["CategoryID"])

	post := models.Post{
		Title: postValuesMap["Title"],
		CategoryID: uint32(categoryIDInt),
		Image: imagePath,
		Description: descPath,
		Creation_date: postValuesMap["Creation_date"],
		Units: uint(unitsInt),
		Price: postValuesMap["Price"],
	}

	db.DB.Create(&post)
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Post successfully created."))
}

func PutPostHandler(w http.ResponseWriter, r *http.Request) {
	
	// we store the form-data text values in this map
	postValuesMap := make(map[string]string)

	// the filename of both image and html files, with its extensions
	var dscpName string
	var imgName string

	_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	// mr return all the parts in the r.Body
	// params["boundary"] has 26 hyphen set and a big number 
	// (in postman; that value isn't equal to Content-Type boundary value)
	mr := multipart.NewReader(r.Body, params["boundary"])
	
	for {
		// each part is a form-data field, files too
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		slurp, err := io.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}

		if p.FileName() != "" {
			// first case, description html file, CORRECT THE NESTED IFs
			var fileName string
			if path.Ext(p.FileName()) == ".html" {
				dscpName = uuid.New().String() + path.Ext(p.FileName())
				fileName = dscpName
				
			} else {
				imgName = uuid.New().String() + path.Ext(p.FileName())
				fileName = imgName
			}
			// in this case we use os.OpenFile() to overwrite a file
			file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.FileMode(0666))
			file.Write(slurp)
		} else {
			postValuesMap[p.FormName()] = string(slurp)
			fmt.Printf("\nPart %q: %q", p.FormName(), slurp)
		}
	}

	// later it will be a url, for now it's only a local path
	imagePath := fmt.Sprintf("files/%s", imgName)
	descPath := fmt.Sprintf("files/%s", dscpName)

	unitsInt, _ := strconv.Atoi(postValuesMap["Units"])
	categoryIDInt, _ := strconv.Atoi(postValuesMap["CategoryID"])

	post := models.Post{
		Title: postValuesMap["Title"],
		CategoryID: uint32(categoryIDInt),
		Image: imagePath,
		Description: descPath,
		Creation_date: postValuesMap["Creation_date"],
		Units: uint(unitsInt),
		Price: postValuesMap["Price"],
	}

	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error to convert string postID to integer."))
	} else {
		post.ID = uint32(postID)
		
		db.DB.Save(&post)
		
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Post successfully edited"))
	}
}
	
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	
	postID := vars["id"]
	
	db.DB.Delete(&models.Post{}, postID)
	
	w.WriteHeader(http.StatusContinue)
	w.Write([]byte("Post successfully deleted."))
	
}
