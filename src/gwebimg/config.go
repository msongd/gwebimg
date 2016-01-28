package main

import (
	"html/template"
)

type Config struct {
	Address string
	Port   int
	//TemplatePath string
	//WebResourcePath string
	ImgPath string
	Template *template.Template
}

type PageIndexTpl struct {
	Title string
	//ChapterListJSON string
	ChapterList []string
}

