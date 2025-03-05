package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "./test.txt"
	writeFile(filePath)

	readFile(filePath)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file, "Hello From Me!")
	fmt.Printf("Wrote a file with %v characters\n", length)
}

func readFile(filePath string) {
	file, err := os.ReadFile(filePath)
	checkError(err)
	fmt.Println("Text read from file", string(file))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
