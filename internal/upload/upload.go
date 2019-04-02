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

// Upload Flags
var (
	bucket    string
	directory string
	all       bool
	image     bool
)

func init() {
	flag.StringVarP(&bucket, "bucket", "b", "", "Gives a bucket name for aws upload")
	flag.StringVarP(&directory, "directory", "d", "", "Creates a 'directory' within s3 bucket")
	flag.BoolVarP(&all, "all", "a", false, "Should upload all files in current directory")
	flag.BoolVarP(&image, "image", "i", false, "Is the file an image")
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return true
	}

	return fileInfo.IsDir()
}

// All Uploads site
func All(fpath string) {
	files, err := filepath.Glob(filepath.Join(fpath, "*"))
	if err != nil {
		log.Fatalln(err)
	}

	var infos []S3UploadInfo

	for _, f := range files {
		filename := filepath.Base(f)

		if directory != "" {
			filename = directory + "/" + filename
		}

		if isDirectory(f) {
			continue
		}

		file, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
		}

		var contentType string
		switch filepath.Ext(f) {
		case ".xml":
			contentType = "application/xml"
			break
		case ".svg":
			contentType = "image/svg+xml"
			break
		case "":
			contentType = "text/html"
			break
		default:
			contentType = ""
		}

		s3Info := S3UploadInfo{
			Bucket:   bucket,
			Filename: filename,
			File:     file,
		}

		if contentType != "" {
			s3Info.ContentType = contentType
		}

		infos = append(infos, s3Info)
	}

	Batch(infos)

}

// Image Upload an image and returns new hashed file name
func Image(f, fpath string) {

	if isDirectory(fpath) {
		log.Fatalln("File is a directory")
	}

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

	if directory != "" {
		filename = directory + "/" + filename
	}

	s3Info := S3UploadInfo{
		Bucket:   bucket,
		Filename: filename,
		File:     bytes.NewReader(fReader),
	}

	s3Info.Upload()

	fmt.Println("Uploaded:", filename, "to", bucket)
}

// File Uploads a file to aws
func File(f, fpath string) {
	if isDirectory(fpath) {
		log.Fatalln("File is a directory")
	}

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatalln(err)
	}

	var contentType string
	switch filepath.Ext(f) {
	case ".xml":
		contentType = "application/xml"
		break
	case ".svg":
		contentType = "image/svg+xml"
		break
	case "":
		contentType = "text/html"
		break
	default:
		contentType = ""
	}

	if directory != "" {
		f = directory + "/" + f
	}

	s3Info := S3UploadInfo{
		Bucket:       bucket,
		Filename:     f,
		File:         file,
		CacheControl: "max-age=300",
	}

	if contentType != "" {
		s3Info.ContentType = contentType
	}

	s3Info.Upload()

	fmt.Println("Uploaded", f, "to", bucket)
}

// Upload Handles calls to upload command
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
