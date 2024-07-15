package utils

import (
	"fmt"
	"log"
	"os"
)

func createFile(path string, fileName string) (string, error) {
	home, _ := os.UserHomeDir()
	configPath := home + "/.adoctl"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}

	file, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	fmt.Printf("adoctl auth file created successfully at %s\n", configPath)
	fileCreated := configPath + "/" + fileName
	return fileCreated, nil
}
