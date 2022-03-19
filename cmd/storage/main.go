package main

import (
	"fmt"
	"log"

	"github.com/gadzhimagomedov81/storage/internal/storage"
)

func main() {

	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It work", file)

	restoredFile, err := st.GetById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It is restored", restoredFile)
}
