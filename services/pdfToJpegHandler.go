package services

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	g "go-gs/globals"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func PdfToJpegHandler(w http.ResponseWriter, r *http.Request) {
	//get single pdf file and process for single file...
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
		w.WriteHeader(400)
		msg := "invalid method"
		respBody, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
		}
		w.Write(respBody)
		return
	} else {
		saveFile(w, r)
	}
}

func saveFile(w http.ResponseWriter, r *http.Request) {

	// Parse the form data to retrieve the file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server to save the uploaded file
	newFile, err := os.Create(fileHeader.Filename)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the file content from the request body to the newly created file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")
	folder, err := processFile(fileHeader.Filename)
	if err != nil {
		http.Error(w, "Error pdf convertor engine", http.StatusInternalServerError)
		return
	}
	sendResponse(w, r, folder)
}

func processFile(filename string) (string, error) {
	//processPDf
	err := convert(g.Cfg, filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(filename, filepath.Ext(filename)), nil
}
func sendResponse(w http.ResponseWriter, r *http.Request, folder string) {
	// Open the folder and read the list of files
	fileInfos, err := os.ReadDir(folder)
	if err != nil {
		http.Error(w, "Error reading folder", http.StatusInternalServerError)
		return
	}

	// Set the response header as application/zip
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=files.zip")

	// Create a zip writer to write the files to the response
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// Iterate over each file in the folder and add them to the zip archive
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			filePath := filepath.Join(folder, fileInfo.Name())

			// Open the file
			file, err := os.Open(filePath)
			if err != nil {
				http.Error(w, "Error opening file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Create a new zip file entry
			zipFile, err := zipWriter.Create(fileInfo.Name())
			if err != nil {
				http.Error(w, "Error creating zip entry", http.StatusInternalServerError)
				return
			}

			// Copy the file content to the zip file entry
			_, err = io.Copy(zipFile, file)
			if err != nil {
				http.Error(w, "Error copying file content", http.StatusInternalServerError)
				return
			}
		}
	}

}
