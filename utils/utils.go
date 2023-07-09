package utils

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ListFilesInDirectory(folderPath string) []fs.FileInfo {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetFileContentType(filepath string) (string, error) {
	// Open the file whose type you
	// want to check
	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	// Get the file content
	// to sniff the content type only the first
	// 512 bytes are used.

	buf := make([]byte, 512)

	_, err = file.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)
	if err != nil {
		return "", err
	}

	return contentType, nil
}

func CreateNewDir(folderPath string) error {
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}
	return nil
}
