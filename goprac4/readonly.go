package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()

	buf := make([]byte, 100)
	n,_ := file.Read(buf)

	fmt.Println(string(buf[:n]))
}
