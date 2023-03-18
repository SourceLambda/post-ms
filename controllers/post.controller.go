package controller

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"path"

	"github.com/SourceLambda/sourcelambda_post_ms/db"
	"github.com/google/uuid"
)

func MultipartHandler(mr *multipart.Reader, postMap map[string]string) (descName, imgName string) {

	// the filename of both image and html files, with its extensions
	// this values will be returned
	var desc string
	var img string

	for {
		// each part is a form-data field, files too
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			return
		}

		// if p is reading a file, we call the db.CreateFile()
		if p.FileName() != "" {

			// first case, description html file, CORRECT THE NESTED IFs
			var fileName string
			if path.Ext(p.FileName()) == ".html" {
				desc = "htmls/" + uuid.New().String() + path.Ext(p.FileName())
				fileName = desc

			} else {
				img = "images/" + uuid.New().String() + path.Ext(p.FileName())
				fileName = img
			}

			// call of db.CreateFile()
			err := db.CreateFile(fileName, p)
			if err != nil {
				fmt.Print(err)
				return
			}
		} else {

			// if p is a normal form value, we store it in postMap
			slurp, err := io.ReadAll(p)
			if err != nil {
				log.Print(err)
				return
			}
			postMap[p.FormName()] = string(slurp)
		}
	}

	return desc, img

}
