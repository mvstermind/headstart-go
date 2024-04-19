package templates

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO: This looks ugly too
func FlatTemp() {
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

	dir = filepath.Join(currentUser, "headstart")
	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/handler.go", dir), []byte("package main\n\nfunc hanlder() {\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/service.go", dir), []byte("package main\n\nfunc service() {\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/database.go", dir), []byte("package main\n\nfunc database() {\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/config.go", dir), []byte("package main\n\nfunc config() {\n}"), 0660)

	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(currentUser, "headstart/internal/domain")
	err = os.MkdirAll(dir, 0750)
	err = os.WriteFile(fmt.Sprintf("%s/model.go", dir), []byte("package domain\n\nfunc model(){\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/repository.go", dir), []byte("package domain\n\nfunc repository(){\n}"), 0660)
	dir = filepath.Join(currentUser, "headstart/static")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/templates")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/scripts")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/configs")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/tests")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/docs")
	err = os.MkdirAll(dir, 0750)

}
