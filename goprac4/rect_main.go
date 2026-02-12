package main

import (
	"fmt"
	"practicals/rectangle"
)

func main() {
	var l, w float64
	fmt.Scanln(&l, &w)

	fmt.Println("Area:", rectangle.Area(l, w))
}
