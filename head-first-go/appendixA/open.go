package main

import (
	"log"
	"os"
)

func main() {

	// file, err := os.OpenFile("data.txt", os.O_RDONLY, os.FileMode(0600))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if scanner.Err() != nil {
	// 	log.Fatal(scanner.Err())
	// }

	file, err := os.OpenFile("data.txt", os.O_WRONLY, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte("End of file"))
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
