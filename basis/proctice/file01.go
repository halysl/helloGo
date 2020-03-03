package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readFromFile() {
	file, err := os.Open("/tmp/test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
	 err := file.Close()
	 if err != nil {
	 	log.Fatal(err)
	 }
	}()

	content := make([]byte, 128)
	tmp := make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}