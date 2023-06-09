package main

import (
	"log"
	"os"
	"path"
)

const rootDir = ".nanovault"

var rootPath string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("An error occurred while retrieving user's home folder")
	}

	rootPath = path.Join(home, rootDir)

	if err := os.Mkdir(rootPath, os.ModeDir); err != nil {
		if !os.IsExist(err) {
			log.Fatal("An error occurred while creating " + rootDir + " folder")
		}
	}
}

func main() {

}
