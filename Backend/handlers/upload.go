package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const uploadPath = "./static/uploads/"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("video")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create upload directory if not exists
		os.MkdirAll(uploadPath, os.ModePerm)

		// Save file
		dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.ReadFrom(file)
		if err != nil {
			http.Error(w, "Error writing file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
