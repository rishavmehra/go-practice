package fileupload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error file retrieving the error", err)
		return
	}
	defer file.Close()

	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)

	tempfile, err := os.CreateTemp("images", "upload-*.png")
	if err != nil {
		fmt.Println("Error while", err)
	}
	defer tempfile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error file read the file: ", err)
	}

	tempfile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File \n")

}
