package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Marks int
}

func main() {
	s := []Student{{"Rahul",75},{"Anita",90}}

	sort.Slice(s, func(i,j int) bool {
		return s[i].Marks < s[j].Marks
	})

	fmt.Println(s)
}
