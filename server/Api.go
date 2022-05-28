package main

import(
	"fmt"
)
type homework struct {
	ID int `json:"ID"`
	name string `json:name`
	content string`json:content`
}

type Allhomeworks[]homework 


func main() {
	fmt.Println("Hello world Champ")
}