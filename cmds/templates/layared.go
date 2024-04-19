package templates

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO: A LOT OF MESS IN HEREEEEEEEE TO CLEAN UP
func Layared() {
	currentUser, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentUser, "headstart")

	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/main.go", dir), []byte("package main\n\nfunc main() {\n\n}"), 0660)

	dir = filepath.Join(currentUser, "headstart/web")
	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/handler.go", dir), []byte("package web\n\nfunc hanlder() {\n}"), 0660)

	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(currentUser, "headstart/web/static")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/web/templates")
	err = os.MkdirAll(dir, 0750)

	dir = filepath.Join(currentUser, "headstart/api")
	err = os.MkdirAll(dir, 0750)
	err = os.WriteFile(fmt.Sprintf("%s/routes.go", dir), []byte("package api\n\nfunc routes(){\n}"), 0660)
	dir = filepath.Join(currentUser, "headstart/api/middleware")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/data")
	err = os.MkdirAll(dir, 0750)
	err = os.WriteFile(fmt.Sprintf("%s/database.go", dir), []byte("package data\n\nfunc database(){\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/repository.go", dir), []byte("package data\n\nfunc repository(){\n}"), 0660)
	dir = filepath.Join(currentUser, "headstart/configs")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/tests")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/docs")
	err = os.MkdirAll(dir, 0750)
}
