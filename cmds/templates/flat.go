package templates

import (
	"log"
	"os"
	"path/filepath"
)

func FlatTemp() {
	currentUser, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	baseDir := filepath.Join(currentUser, "headstart")

	directories := []string{
		baseDir,
		filepath.Join(baseDir, "internal/domain"),
		filepath.Join(baseDir, "static"),
		filepath.Join(baseDir, "templates"),
		filepath.Join(baseDir, "scripts"),
		filepath.Join(baseDir, "configs"),
		filepath.Join(baseDir, "tests"),
		filepath.Join(baseDir, "docs"),
	}
	files := map[string]string{
		filepath.Join(baseDir, "main.go"):                       "package main\n\nfunc main() {\n\n}",
		filepath.Join(baseDir, "handler.go"):                    "package main\n\nfunc Hanlder() {\n}",
		filepath.Join(baseDir, "service.go"):                    "package main\n\nfunc Service() {\n}",
		filepath.Join(baseDir, "database.go"):                   "package main\n\nfunc Database() {\n}",
		filepath.Join(baseDir, "config.go"):                     "package main\n\nfunc Config() {\n}",
		filepath.Join(baseDir, "internal/domain/model.go"):      "package domain\n\nfunc Model(){\n}",
		filepath.Join(baseDir, "internal/domain/repository.go"): "package domain\n\nfunc Repository(){\n}",
	}

	// make directories
	for _, dir := range directories {
		err := os.MkdirAll(dir, 0750)
		if err != nil {
			log.Fatal(err)
		}
	}

	// write content to files in directories
	for file, content := range files {
		err := os.WriteFile(file, []byte(content), 0660)
		if err != nil {
			log.Fatal(err)
		}
	}
}
