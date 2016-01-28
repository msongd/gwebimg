package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	"os"
	//"io"
	//"io/ioutil"
	//"log"
	"strconv"
	//"regexp"
	//"strings"
	"encoding/json"
	"fmt"
	"path/filepath"
)
/*
func Index(w http.ResponseWriter, r *http.Request, cfg *TOMLConfig) {
	endpoints := GetClientsList(cfg)
	IndexTemplate(w, endpoints, cfg)
}
*/
func ChapterList(w http.ResponseWriter, r *http.Request, cfg *Config) {
	chapterList := ListDir(cfg.ImgPath)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	err:= json.NewEncoder(w).Encode(chapterList)
	if err != nil {
		GlobalLog.Println("[VAR] unable to encode json")
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
}

func PageList(w http.ResponseWriter, r *http.Request, cfg *Config) {
	chapter := mux.Vars(r)["chapter"]
	GlobalLog.Println("[VAR] chapter:", chapter)
	if chapter == "" {
		GlobalLog.Println("[VAR] unable to get chapter")
		w.WriteHeader(404)
		fmt.Fprint(w, "chapter not found")
		return
	}
	pageList := ListFile(cfg.ImgPath+"/"+chapter, ".jpg .png .jpeg", cfg.ImgPath, "/static/images/"+chapter+"/")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	err:= json.NewEncoder(w).Encode(pageList)
	if err != nil {
		GlobalLog.Println("[VAR] unable to encode json")
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request, cfg *Config) {
	page := PageIndexTpl{}
	chapterList := ListDir(cfg.ImgPath)
	page.ChapterList = chapterList
	/*
	chapterListJson, err := json.Marshal(chapterList)
	if err != nil {
		GlobalLog.Println("Marshal chapter list to json err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "marshal json chapter list error")
		return
	}
	page.ChapterListJSON = string(chapterListJson)
	*/
	page.Title = filepath.Base(cfg.ImgPath)
	
	err := cfg.Template.ExecuteTemplate(w, "index", page)
	if err != nil {
		GlobalLog.Println("Template execute err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "template error")
		return
	}
}

func Server(cfg *Config) {
	listenString := cfg.Address + ":" + strconv.Itoa(cfg.Port)
	mainRoute := mux.NewRouter()

	
	mainRoute.HandleFunc("/chapter-list", 
		func( w http.ResponseWriter, r *http.Request) {
			ChapterList(w,r,cfg)
			})

	mainRoute.HandleFunc("/page-list/{chapter}", 
		func( w http.ResponseWriter, r *http.Request) {
			PageList(w,r,cfg)
			})

	mainRoute.HandleFunc("/", func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html; charset: UTF-8")
			//w.Write(Index_html)
			Index(w, r, cfg)
			})
	
	mainRoute.HandleFunc("/static/include/pure-min.css", func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css; charset: UTF-8")
			w.Write(Pure_min_css)
			})
	mainRoute.HandleFunc("/static/include/scooch-style.min.css",func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css; charset: UTF-8")
			w.Write(Scooch_style_min_css)
			})
	mainRoute.HandleFunc("/static/include/scooch.min.css",func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css; charset: UTF-8")
			w.Write(Scooch_min_css)
			})
	mainRoute.HandleFunc("/static/include/scooch.min.js",func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/javascript; charset: UTF-8")
			w.Write(Scooch_min_js)
			})
	mainRoute.HandleFunc("/static/include/zepto.min.js",func( w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/javascript; charset: UTF-8")
			w.Write(Zepto_min_js)
			})
	
	mainRoute.PathPrefix("/static/images/").Handler(http.StripPrefix("/static/images/",http.FileServer(http.Dir(cfg.ImgPath))))
	
	loggedRouter := handlers.LoggingHandler(os.Stderr, mainRoute)

	err := http.ListenAndServe(listenString,loggedRouter)
	if err != nil {
		GlobalLog.Fatal("ListenAndServe: ", err)
	}
}
