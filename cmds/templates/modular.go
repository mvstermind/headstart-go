package templates

import (
	"log"
	"os"
	"path/filepath"
)

// TODO: yep don't forget to clean this one 2
func Modular() {

	currentUser, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentUser, "headstart/module1")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/cmd")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/internal")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/pkg")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/api")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/web")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/scripts")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/configs")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/tests")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module1/docs")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/module2")
	err = os.MkdirAll(dir, 0750)
}
