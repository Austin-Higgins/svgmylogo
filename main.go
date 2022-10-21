package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	f, h, err := r.FormFile("myUpload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Printf("Uploaded File: %+v\n", h.Filename)
	fmt.Printf("File Size: %+v\n", h.Size)
	fmt.Printf("MIME Header: %+v\n", h.Header)

	tempFile, err := os.CreateTemp(".\\temp", "upload*.png")
	if err != nil {
		fmt.Println(err)
	} else {
		//If the photo exists, pass it to the API
		fmt.Println(tempFile.Name())

	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":4040", nil)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	fmt.Println("Hello World")
	setupRoutes()

}
