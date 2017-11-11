//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/

package main

//To use the net/http package, it must be imported
import (
	"html/template"
	"net/http"
)

type Data struct {
	Message, Chat string
	//Chat    string
}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Guessing Game! ")
	m := Data{Message: "Eliza robot chat", Chat: "Hello how are you today"}
	//g.Execute(w, Data{Message: "Guess a number between 1 and 20"})
	//o := Data

	//g.Execute(w, Data{Message:"Hello how are you today"}

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, &m)
	//t.Execute(w, &o)
}

func main() {
	//call handler function
	http.HandleFunc("/", templateHandler)

	//http.HandleFunc("/guess", templateHandler)
	http.ListenAndServe(":8080", nil)
}
