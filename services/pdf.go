package services

import (
	"bytes"
	"fmt"
	m "go-gs/models"
	"go-gs/utils"
	"log"
	"os/exec"
)

func PdfToJpg(cfg m.Configs) {
	files := utils.ListFilesInDirectory(m.FOLDER_PATH)
	for k, val := range files {
		fmt.Println(k, val)
		filepath := m.FOLDER_PATH + "/" + val.Name()
		if fileType, err := utils.GetFileContentType(filepath); err == nil {
			if fileType == "application/pdf" {
				convert(cfg, filepath)
			}
		}

	}

}

func convert(cfg m.Configs, filepath string) {
	//gswin64c.exe -sDEVICE=jpeg -sOutputFile=page-%03d.jpg -r100x100 -f file1.pdf
	cmd := exec.Command(cfg.Gs, "-sDEVICE=jpeg", "-sOutputFile=page-%03d.jpg", "-r100x100", "-f", filepath)
	fmt.Println(cfg.Gs, "-sDEVICE=jpeg", "-sOutputFile=page-%03d.jpg", "-r100x100", "-f", filepath)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Processed files: %q\n", out.String())
}
