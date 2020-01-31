package core

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func ProjectPath() string {
	projectPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return projectPath
}

func ObjectsPath() string {
	return path.Join(ProjectPath(), ".git/objects")
}

func RefsPath() string {
	return path.Join(ProjectPath(), ".git/refs")
}

func WorkspaceFiles() []string {
	paths := make([]string, 0)
	err := filepath.Walk(ProjectPath(),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == ".git" {
				return filepath.SkipDir
			} else if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return paths
}
