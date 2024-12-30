package main

import (
	"net/http"

	"github.com/yogeshparihar2105Github/Video-Sharing-Website/Backend/handlers"
)

func main() {
	http.HandleFunc("/upload", handlers.UploadHandler)
	http.HandleFunc("/stream", handlers.StreamHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve HTML templates
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.ListenAndServe(":8080", nil)
}
