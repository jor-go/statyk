package upload

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

/* Upload Flags */
var (
	bucket string
	region string
	all    bool
	image  bool
)

func init() {
	flag.StringVarP(&bucket, "bucket", "b", "", "Gives a bucket name for aws upload")
	flag.StringVarP(&region, "region", "r", "", "Gives a region name for aws upload")
	flag.BoolVarP(&all, "all", "a", false, "Should upload all files in current directory")
	flag.BoolVarP(&image, "image", "i", false, "Is the file an image")
}

/*All : Uploads site*/
func All(fpath string) {
	files, err := filepath.Glob(filepath.Join(fpath, "*"))
	if err != nil {
		log.Fatalln(err)
	}

	var infos []S3UploadInfo

	for _, f := range files {
		filename := filepath.Base(f)
		file, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
		}

		var contentType string
		if filepath.Ext(f) == "" {
			contentType = "text/html"
		} else if filepath.Ext(f) == ".xml" {
			contentType = "application/xml"
		}

		s3Info := S3UploadInfo{
			Bucket:       bucket,
			Filename:     filename,
			File:         file,
			ContentType:  contentType,
			CacheControl: "max-age=300",
		}

		infos = append(infos, s3Info)
	}

	UploadBatch(infos)

}

/*Image : Upload an image and returns new hashed file name*/
func Image(f, fpath string) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalln(err)
	}

	hash := md5.New()

	fReader, err := ioutil.ReadAll(io.TeeReader(file, hash))
	if err != nil {
		log.Fatalln(err)
	}

	filename := hex.EncodeToString(hash.Sum(nil)) + filepath.Ext(f)

	s3Info := S3UploadInfo{
		Bucket:   bucket,
		Filename: filename,
		File:     bytes.NewReader(fReader),
	}

	s3Info.Upload()

	fmt.Println("Uploaded:", filename, "to", bucket)
}

func File(f, fpath string) {
	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalln(err)
	}

	var contentType string
	if filepath.Ext(f) == "" {
		contentType = "text/html"
	} else if filepath.Ext(f) == ".xml" {
		contentType = "application/xml"
	}

	s3Info := S3UploadInfo{
		Bucket:       bucket,
		Filename:     f,
		File:         file,
		ContentType:  contentType,
		CacheControl: "max-age=300",
	}

	s3Info.Upload()

	fmt.Println("Uploaded", f, "to", bucket)
}

func Upload() {

	flag.Parse()

	fileArg := flag.Arg(1)
	fileLocation, err := filepath.Abs(fileArg)
	if err != nil {
		log.Fatalln(err)
	}
	filename := filepath.Base(fileLocation)

	if bucket == "" {
		log.Fatalln("No bucket given")
	}

	if !all && filename == "" {
		log.Fatalln("No file given")
	}

	if all {
		All(fileLocation)
	} else if image {
		Image(filename, fileLocation)
	} else {
		File(filename, fileLocation)
	}
}
