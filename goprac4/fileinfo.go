package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var filename string
	fmt.Print("Enter file name: ")
	fmt.Scanln(&filename)

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n--- File Information ---")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Permission:", info.Mode())
	fmt.Println("Modified:", info.ModTime().Format(time.RFC1123))
	fmt.Println("Is Directory:", info.IsDir())
}
