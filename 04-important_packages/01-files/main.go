package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Writing to file
	newFile, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	}

	size, err := newFile.WriteString("Hello, World!\n")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created! size: %v bytes\n", size)

	size, err = newFile.Write([]byte("Updating content of file.txt"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("File updated! new size: %v bytes\n", size)

	newFile.Close()

	// Reading file
	fileContent, err := os.ReadFile("./file.txt")

	if err != nil {
		panic(err)
	}

	println("Reading file:")
	println(string(fileContent))

	// Reading file using Reader
	file, err := os.Open("./file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Print(string(buffer[:n]))
	}

	fmt.Println()
}
