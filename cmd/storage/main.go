package main

import (
	"fmt"
	"log"

	"github.com/AzamatIshmuratov/storage/internal/storage"
)

func main() {
	storage := storage.NewStorage()

	file, err := storage.Upload("test.txt", []byte("Hello"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := storage.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Its restored file", restoredFile)

}
