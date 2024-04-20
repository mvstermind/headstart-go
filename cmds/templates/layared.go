package templates

import (
	"log"
	"os"
	"path/filepath"
)

func Layared(rootDir string) {
	directories := []string{
		"web",
		"web/static",
		"web/template",
		"api",
		"api/middleware",
		"data",
		"config",
		"tests",
		"docs",
	}

	current, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Join(current, rootDir)

	err = os.MkdirAll(baseDir, 0750)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(baseDir, "main.go"), []byte("package main\n\nfunc main() {\n\n}"), 0660)
	if err != nil {
		log.Fatal(err)
	}

	//make direcotries and write content to files
	for _, dir := range directories {
		fullPath := filepath.Join(baseDir, dir)
		err = os.MkdirAll(fullPath, 0750)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case dir == "web":
			err = os.WriteFile(filepath.Join(fullPath, "handler.go"), []byte("package web\n\nfunc hanlder() {\n}"), 0660)
		case dir == "api":
			err = os.WriteFile(filepath.Join(fullPath, "routes.go"), []byte("package api\n\nfunc routes(){\n}"), 0660)
		case dir == "data":
			err = os.WriteFile(filepath.Join(fullPath, "database.go"), []byte("package data\n\nfunc database(){\n}"), 0660)
		}
		err = os.WriteFile(filepath.Join(fullPath, "repository.go"), []byte("package data\n\nfunc repository(){\n}"), 0660)

	}
}
