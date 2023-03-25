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

// write some documentation :v
func MultipartHandler(mr *multipart.Reader, postMap map[string]string) (imgName string, err error) {

	var imgFilename string

	for {
		// each part is a form-data field, files too
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Print(err)
			return "", err
		}

		// if p is reading a file, we call the db.CreateFile()
		if p.FileName() != "" {

			imgFilename = "images/" + uuid.New().String() + path.Ext(p.FileName())

			err := db.CreateFile(imgFilename, p)
			if err != nil {
				fmt.Print(err)
				return "", err
			}
		} else {
			slurp, err := io.ReadAll(p)
			if err != nil {
				log.Print(err)
				return "", err
			}
			postMap[p.FormName()] = string(slurp)
		}
	}

	return imgFilename, nil

}
