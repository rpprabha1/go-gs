package services

import (
	"fmt"
	models "go-gs/models"
	"os"
	"testing"
)

func Test_PdfToJpg(t *testing.T) {
	models.WORKERS = 5
	models.FOLDER_PATH = "../tmp"
	_, err := os.Stat("../tmp/test")
	if err == nil {
		// Directory exists, so delete it
		err = os.RemoveAll("../tmp/test")
		if err != nil {
			t.Error("Error deleting directory:", err)
			return
		}
		fmt.Println("Directory deleted successfully.")
	}
	cfg := models.Configs{
		Gs: "C:/Program Files/gs/gs10.01.2/bin/gswin64c.exe",
	}
	err = convert(cfg, "../tmp/test.pdf")
	if err != nil {
		t.Error("Error in convert function: ", err)
	}
}
