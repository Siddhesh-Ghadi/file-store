package fileutil

import (
	"log"
	"os"
	"strings"
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

func GetWordCount(path string) (int, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	words := strings.Fields(string(raw)) 
	return len(words), nil
}

func GetWords(path string) ([]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	words := strings.Fields(string(raw)) 
	return words, nil
}

func GetWordFreq(words []string, limit int, order string) (map[string]int) {
	// use map to store freq
	wordFreq := make(map[string]int)
    for _, w := range words {
        wordFreq[w]++
    }
	// TODO: implement sort
	index := 0
	retVal := make(map[string]int)
	for k, v := range wordFreq {
		if index < limit {
			retVal[k] = v
		}
		index++
    }
	return retVal
}