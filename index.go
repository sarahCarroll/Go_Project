//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/

package main

//To use the net/http package, it must be imported
import (
	"html/template"
	"net/http"
)

type Data struct {
	Message string
}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Guessing Game! ")
	m := Data{Message: "Guess a number between 1 and 20: "}
	//g.Execute(w, Data{Message: "Guess a number between 1 and 20"})

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, m)
}

func main() {
	//call handler function
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", templateHandler)

	http.HandleFunc("/guess", templateHandler)
	http.ListenAndServe(":8080", nil)
}