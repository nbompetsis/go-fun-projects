package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ScanDir(filename string) error {
	fmt.Println(filename)
	files, err := ioutil.ReadDir(filename)

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			filePath := filepath.Join(filename, file.Name())
			err = ScanDir(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("File:", file.Name())
		}
	}
	return nil
}

func main() {
	ScanDir(os.Args[1])
}
