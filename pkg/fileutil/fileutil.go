package fileutil

import (
	"log"
	"os"
)

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
        log.Fatal(err)
    }
}

func GetAllFileNames(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	returnVal := []string{}
	for _, file := range files {
		if !file.IsDir(){
			returnVal = append(returnVal, file.Name())
		}
	}
	return returnVal, nil
}