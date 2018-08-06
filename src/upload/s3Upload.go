package upload

import (
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess *session.Session
var uploader *s3manager.Uploader

func init() {
	// Force enable Shared Config support
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	uploader = s3manager.NewUploader(sess)
}

/*S3UploadInfo : Info needed to upload file to AWS S3*/
type S3UploadInfo struct {
	File         io.Reader
	Bucket       string
	ContentType  string
	Filename     string
	CacheControl string
}

/*Upload : Uploads file to s3*/
func (s *S3UploadInfo) Upload() {
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:       aws.String(s.Bucket),
		ContentType:  aws.String(s.ContentType),
		CacheControl: aws.String(s.CacheControl),
		Key:          aws.String(s.Filename),
		Body:         s.File,
	})
	if err != nil {
		// Print the error and exit.
		log.Fatalf("Unable to upload %q to %q, %v", s.Filename, s.Bucket, err)
	}
}

func UploadBatch(uploads []S3UploadInfo) {
	var objects []s3manager.BatchUploadObject
	for _, s := range uploads {
		var obj s3manager.BatchUploadObject

		input := &s3manager.UploadInput{
			Bucket:       aws.String(s.Bucket),
			ContentType:  aws.String(s.ContentType),
			CacheControl: aws.String(s.CacheControl),
			Key:          aws.String(s.Filename),
			Body:         s.File,
		}
		obj.Object = input
		objects = append(objects, obj)
		fmt.Println("Uploading:", s.Filename)
	}

	iter := &s3manager.UploadObjectsIterator{Objects: objects}
	if err := uploader.UploadWithIterator(aws.BackgroundContext(), iter); err != nil {
		log.Fatalln(err)
	}
}
