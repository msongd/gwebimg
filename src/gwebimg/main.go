package main

import (
	"log"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var HOST     string
var PORT     int
var IMG_DIR string

var GlobalLog *log.Logger

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func initConfig() {
	flag.StringVar(&HOST, "h", "-", "listening interface")
	flag.IntVar(&PORT, "p", 8080, "Port to listen")
	flag.StringVar(&IMG_DIR, "i", ".", "image dir")
	
	flag.Parse()
	if HOST == "-" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateConfig(cfg *Config) bool {
	fullImgPath,err := filepath.Abs(cfg.ImgPath)
	if err != nil {
		fmt.Println("[ERR] imagepath err", err)
		return false
	}
	cfg.ImgPath = fullImgPath
	
	e, err := exists(cfg.ImgPath)
	if err != nil {
		fmt.Println("[ERR] imgpath check:", err)
		return false
	} else if !e {
		fmt.Println("[ERR] imgpath is not exists",)
		return false
	}

	return true
}

func main() {
	initConfig()
	
	GlobalLog = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)

	cfg := new(Config)
	cfg.Address = HOST
	cfg.Port = PORT
	cfg.ImgPath = IMG_DIR
	//cfg.TemplatePath = TEMPLATE_DIR
	
	ValidateConfig(cfg)
	GlobalLog.Println("[Config]:", cfg)
	Server(cfg)
	
}
