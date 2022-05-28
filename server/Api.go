package main

import(
	"fmt"
	"github.com/gorilla/mux"
)
type homework struct {
	ID int `json:"ID"`
	Name string `json:Name`
	Content string`json:Content`
}

type Allhomeworks []homework 

var h = Allhomeworks{
	{
		ID: 1,
		Name: "First task",
		Content: "Cont", 
	},
}
func main() {
	fmt.Println("Hello world Champ")
}