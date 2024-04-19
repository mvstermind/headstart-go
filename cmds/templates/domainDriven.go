package templates

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// TODO: organize this thing 2
func DomainDriven() {
	currentUser, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(currentUser, "headstart/cmd/app1")
	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Fatal(err)
	}

	dir = filepath.Join(currentUser, "headstart/cmd/app2")
	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(currentUser, "headstart/internal/auth")
	err = os.MkdirAll(dir, 0750)

	err = os.WriteFile(fmt.Sprintf("%s/handler.go", dir), []byte("package auth\n\nfunc hanlder() {\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/service.go", dir), []byte("package auth\n\nfunc service() {\n}"), 0660)

	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(currentUser, "headstart/internal/orders")
	err = os.MkdirAll(dir, 0750)

	err = os.WriteFile(fmt.Sprintf("%s/handler.go", dir), []byte("package orders\n\nfunc handler(){\n}"), 0660)
	err = os.WriteFile(fmt.Sprintf("%s/service.go", dir), []byte("package orders\n\nfunc service(){\n}"), 0660)
	dir = filepath.Join(currentUser, "headstart/pkg/utility")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/api/app1")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/api/app2")
	err = os.MkdirAll(dir, 0750)

	dir = filepath.Join(currentUser, "headstart/web/app1")
	err = os.MkdirAll(dir, 0750)
	dir = filepath.Join(currentUser, "headstart/web/app2")
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
