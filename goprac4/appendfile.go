package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("sample.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString("\nAppended line")
	fmt.Println("Done")
}
