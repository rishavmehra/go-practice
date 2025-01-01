package main

import (
	"net/http"

	fileupload "github.com/rishavmehra/golangPractice/fileupload/uploadfile"
)

func main() {

	http.HandleFunc("/uploaded", fileupload.UploadFile)
	http.ListenAndServe(":8000", nil)

}
