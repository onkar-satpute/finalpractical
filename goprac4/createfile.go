package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content := "Hello Students!\nThis file is created using Go."
	file.WriteString(content)

	fmt.Println("File created!")
}
