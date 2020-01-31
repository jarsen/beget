package core

import (
	"log"
	"os"
	"path"
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
