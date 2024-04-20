package templates

import (
	"log"
	"os"
	"path/filepath"
)

// this is default new template name
var RootDir = "headstart"

func Clean() {
	currentUser, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	basedDir := filepath.Join(currentUser, RootDir)

	dirs := []string{
		basedDir,
		filepath.Join(basedDir, "cmd"),
		filepath.Join(basedDir, "cmd/project"),
		filepath.Join(basedDir, "internal"),
		filepath.Join(basedDir, "internal/app"),
		filepath.Join(basedDir, "internal/domain"),
		filepath.Join(basedDir, "pkg"),
		filepath.Join(basedDir, "pkg/utility"),
		filepath.Join(basedDir, "api"),
		filepath.Join(basedDir, "web"),
		filepath.Join(basedDir, "scripts"),
		filepath.Join(basedDir, "configs"),
		filepath.Join(basedDir, "tests"),
		filepath.Join(basedDir, "docs"),
	}

	filez := map[string]string{
		filepath.Join(basedDir, "cmd/project/main.go"):           "package main\n\nfunc main(){\n}",
		filepath.Join(basedDir, "internal/app/handler.go"):       "package app\n\nfunc Handler(){\n}",
		filepath.Join(basedDir, "internal/app/service.go"):       "package app\n\nfunc Service(){\n}",
		filepath.Join(basedDir, "internal/domain/model.go"):      "package domain\n\nfunc Model(){\n}",
		filepath.Join(basedDir, "internal/domain/repository.go"): "package domain\n\nfunc Repository(){\n}",
	}

	for _, directory := range dirs {
		os.MkdirAll(directory, 0750)
	}

	for file, content := range filez {
		err := os.WriteFile(file, []byte(content), 0660)
		if err != nil {
			log.Fatal(err)
		}
	}
}
