package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Student struct {
	ID int `xml:"ID"`
	Name string `xml:"Name"`
	Course string `xml:"Course"`
}

func main() {
	data,_:=os.ReadFile("student.xml")
	var s Student
	xml.Unmarshal(data,&s)
	fmt.Println(s)
}
