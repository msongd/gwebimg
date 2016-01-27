package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ListDir(p string) []string {
	lst,err := ioutil.ReadDir(p)
	if err != nil {
		GlobalLog.Println("ListDir err:", err)
		return nil
	}
	result := make([]string, len(lst))
	i:=0
	for _, entry := range lst {
		if entry.IsDir() {
			result[i] = entry.Name()
			i = i+1
		}
	}
	return result[:i]
}

func ListFile(p string, extList string, parent string, uriPrefix string) []string {
	fullPath, err := filepath.Abs(p)
	
	if err != nil {
		GlobalLog.Println("p: ", err)
		return nil
	}
	
	parentFullPath, err := filepath.Abs(parent)
	if err != nil {
		GlobalLog.Println("parent: ", err)
		return nil
	}
	GlobalLog.Println("p: ", fullPath, ",parent:", parentFullPath)
	if strings.HasPrefix(fullPath, parentFullPath) {
		lst,err := ioutil.ReadDir(fullPath)
		if err != nil {
			GlobalLog.Println("ListDir err:", err)
			return nil
		}
		result := make([]string, len(lst))
		i:=0
		for _, entry := range lst {
			if !entry.IsDir() {
				n := entry.Name()
				ext := filepath.Ext(n)
				if strings.Contains(extList, ext) {
					result[i] = uriPrefix + n
					i = i+1
				}
			}
		}
		return result[:i]
	} else {
		return nil
	}
}
