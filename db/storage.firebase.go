package db

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var ctx = context.Background()
var bucketHandle *storage.BucketHandle

func FirebaseStorageConn() error {

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, appError := firebase.NewApp(ctx, nil, opt)
	if appError != nil {
		return fmt.Errorf("firebase.NewApp: %v", appError)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}

	// bucket := "bucket-name"
	bucketHandle, err = client.Bucket(os.Getenv("BUCKET_NAME"))
	if err != nil {
		return fmt.Errorf("client.Bucket: %v", err)
	}
	return nil
}

func CreateFile(filename string, reader io.Reader) error {
	// filename := "object-name"

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	objectHandle := bucketHandle.Object(filename)

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	objectHandle = objectHandle.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	// 	return fmt.Errorf("object.Attrs: %v", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	writer := objectHandle.NewWriter(ctx)
	if _, err := io.Copy(writer, reader); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

func DeleteFile(filename string) (err error) { // completar

	objectHandle := bucketHandle.Object(filename)
	deleteError := objectHandle.Delete(context.Background())

	return deleteError
}
