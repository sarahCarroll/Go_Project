//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Message, Chat string
	//Chat    string
}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func templateHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	x := r.Form["usermsg"]
	fmt.Println(x)
	r.Form["chatbox"] = x

	m := Data{Message: "Eliza robot chat", Chat: "user:"}

	t, _ := template.ParseFiles("guess.html")

	//print to console the input value from the user input

	t.Execute(w, &m)

	fmt.Println(r.Form)
	fmt.Println(r.Form["usermsg"])

}

func main() {

	//call handler function
	http.HandleFunc("/", templateHandler)

	//http.HandleFunc("/guess", templateHandler)
	http.ListenAndServe(":8080", nil)
}
