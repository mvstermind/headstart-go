package templates

import (
	"log"
	"os"
	"path/filepath"
)

func Modular() {
	directories := []string{
		"cmd",
		"internal",
		"pkg",
		"api",
		"web",
		"scripts",
		"configs",
		"tests",
		"docs",
	}

	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	basePath := filepath.Join(currentDir, RootDir, "/module1")
	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(basePath, dir), 0750)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.MkdirAll(filepath.Join(currentDir, RootDir, "/module2"), 0750)
	if err != nil {
		log.Fatal(err)
	}
}
