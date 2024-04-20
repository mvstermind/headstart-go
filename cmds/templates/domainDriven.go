package templates

import (
	"log"
	"os"
	"path/filepath"
)

func DomainDriven(rootDir string) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	basedDir := filepath.Join(currentDir, rootDir)
	err = os.MkdirAll(basedDir, 0750)
	if err != nil {
		log.Fatal(err)
	}
	directories := []string{
		basedDir,
		filepath.Join(basedDir, "cmd"),
		filepath.Join(basedDir, "cmd/app1"),
		filepath.Join(basedDir, "cmd/app2"),
		filepath.Join(basedDir, "internal"),
		filepath.Join(basedDir, "internal/auth"),
		filepath.Join(basedDir, "internal/orders"),
		filepath.Join(basedDir, "pkg"),
		filepath.Join(basedDir, "pkg/utility"),
		filepath.Join(basedDir, "api"),
		filepath.Join(basedDir, "api/app1"),
		filepath.Join(basedDir, "api/app2"),
		filepath.Join(basedDir, "web"),
		filepath.Join(basedDir, "web/app1"),
		filepath.Join(basedDir, "web/app2"),
		filepath.Join(basedDir, "scripts"),
		filepath.Join(basedDir, "configs"),
		filepath.Join(basedDir, "tests"),
		filepath.Join(basedDir, "docs"),
	}

	files := map[string]string{
		filepath.Join(basedDir, "internal/auth/handler.go"):   "package auth\n\nfunc Hanlder() {\n}",
		filepath.Join(basedDir, "internal/auth/service.go"):   "package auth\n\nfunc Service() {\n}",
		filepath.Join(basedDir, "internal/orders/handler.go"): "package orders\n\nfunc Hanlder() {\n}",
		filepath.Join(basedDir, "internal/orders/service.go"): "package orders\n\nfunc Service() {\n}",
	}

	// make directories
	for _, path := range directories {
		err := os.MkdirAll(path, 0750)
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
