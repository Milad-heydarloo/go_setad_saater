package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ServeFileHandler serves a file for viewing or downloading based on file_id
func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request
	fmt.Println("Request received at /serve-file")

	// Parse query parameters
	fileID := r.URL.Query().Get("file_id")
	action := r.URL.Query().Get("action") // "view" or "download"

	// Validate input
	if fileID == "" || (action != "view" && action != "download") {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		fmt.Println("Invalid parameters received at /serve-file")
		return
	}

	// Construct the file metadata URL
	fileMetadataURL := fmt.Sprintf("http://80.210.37.118:9000/api/collections/file/records/%s", fileID)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(fileMetadataURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch file metadata: %v", err), http.StatusInternalServerError)
		fmt.Println("Error fetching file metadata:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to retrieve file metadata: HTTP %d", resp.StatusCode), http.StatusNotFound)
		fmt.Println("File metadata not found for file_id:", fileID)
		return
	}

	// Parse the metadata response to get the file name
	var metadata struct {
		Field []string `json:"field"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		http.Error(w, "Failed to parse file metadata", http.StatusInternalServerError)
		fmt.Println("Error parsing file metadata:", err)
		return
	}

	if len(metadata.Field) == 0 {
		http.Error(w, "File name not found in metadata", http.StatusNotFound)
		fmt.Println("File name not found in metadata for file_id:", fileID)
		return
	}

	fileName := metadata.Field[0]

	// Construct the file URL
	fileURL := fmt.Sprintf("http://80.210.37.118:9000/api/files/8676miv4ghschkb/%s/%s", fileID, fileName)

	// Fetch the actual file from the local server
	resp, err = client.Get(fileURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch file: %v", err), http.StatusInternalServerError)
		fmt.Println("Error fetching file:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Failed to retrieve file: HTTP %d", resp.StatusCode), http.StatusNotFound)
		fmt.Println("File not found for file_id:", fileID)
		return
	}

	// Set headers for viewing or downloading the file
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	if action == "download" {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	} else {
		w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", fileName))
	}

	// Stream file content to the client
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to stream file: %v", err), http.StatusInternalServerError)
		fmt.Println("Error streaming file content:", err)
		return
	}

	// Log success
	fmt.Println("File served successfully:", fileName)
}
