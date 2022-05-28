package main

import(
	"net/http"
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
	x := mux.NewRouter()
	x.HandleFunc("/",)
}