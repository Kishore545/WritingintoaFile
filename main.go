package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("OutputFile.txt")
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer f.Close()
	s := []byte("Hello World ❤️")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
}
