package services

import (
	"bytes"
	"fmt"
	m "go-gs/models"
	"go-gs/utils"
	"os/exec"
	"path/filepath"
	"strings"
)

func PdfToJpg(cfg m.Configs) {
	files := utils.ListFilesInDirectory(m.FOLDER_PATH)
	wp := m.NewWorkerPool(m.WORKERS)
	wp.Run()
	defer wp.Done()
	for _, val := range files {
		filePath := m.FOLDER_PATH + "/" + val.Name()
		if fileType, err := utils.GetFileContentType(filePath); err == nil {
			if fileType == "application/pdf" {
				wp.AddTask(func() {
					convert(cfg, filePath)
				},
				)
			}
		}
	}
}

func convert(cfg m.Configs, filePath string) error {
	folderName := filepath.Dir(filePath) + "/" + strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	err := utils.CreateNewDir(folderName)
	if err != nil {
		return err
	}

	//gswin64c.exe -sDEVICE=jpeg -sOutputFile=page-%03d.jpg -r100x100 -f file1.pdf
	cmd := exec.Command(cfg.Gs, "-sDEVICE=jpeg", "-sOutputFile="+folderName+"/page-%03d.jpg", "-r100x100", "-f", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		return err
	}

	fmt.Println("Processed file: ", filePath)
	return nil
}
