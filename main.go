package main

import (
	"flag"
	"fmt"
	c "go-gs/configs"
	m "go-gs/models"
	"go-gs/services"
	"os"
)

var (
	cfg     m.Configs
	cfgPath string
)

func init() {
	parseFlags()
	cfg = c.LoadConfig(cfgPath)
}

func parseFlags() {
	flagSet := flag.NewFlagSet("flags", flag.PanicOnError)
	flagSet.StringVar(&cfgPath, "config", "tmp/config.json", "config file path")
	flagSet.StringVar(&m.FOLDER_PATH, "path", "./tmp", "folder path containing pdf files")
	flagSet.IntVar(&m.WORKERS, "workers", 4, "number of workers to run")
	flagSet.IntVar(&m.THREADS, "threads", 4, "number of threads to run")
	flagSet.Parse(os.Args[1:])
}

func main() {
	fmt.Println(cfg)
	services.PdfToJpg(cfg)
}
