package main

import (
	"flag"
	"fmt"
	c "go-gs/configs"
	m "go-gs/models"
	"os"
)

var (
	WORKERS     int
	THREADS     int
	FOLDER_PATH string
	cfg         m.Configs
)

func init() {
	cfg = c.LoadConfig()
}

func parseFlags() {
	flagSet := flag.NewFlagSet("greet", flag.PanicOnError)
	flagSet.StringVar(&cfg.Gs, "gs", "gs", "flag for ghostscript path or binary/exe variable name")
	flagSet.StringVar(&cfg.Gs, "config", "config", "config file path")
	flagSet.Parse(os.Args[1:])
}

func main() {
	parseFlags()
	fmt.Println(cfg)
}
