package handlers

import (
	"net/http"
	"path/filepath"
)

const videoPath = "./static/uploads/"

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	video := r.URL.Query().Get("video")
	if video == "" {
		http.Error(w, "Video not found", http.StatusNotFound)
		return
	}

	videoFile := filepath.Join(videoPath, video)
	http.ServeFile(w, r, videoFile)
}
