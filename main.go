package main

import (
	"log"
	"os"
)

func encryptFile(filename string, key int8) {
	text, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	ciphertext := caesar(text, key)

	err = os.WriteFile("ciphertext.txt", ciphertext, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func decryptFile(filename string, key int8) {
	ciphertext, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	text := decoder(ciphertext, key)

	err = os.WriteFile("decoded-text.txt", text, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	encryptFile("text.txt", 26)
	decryptFile("ciphertext.txt", 26)
}
