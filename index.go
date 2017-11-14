//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html
//

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Message, Chat string
	x             []string
	//Chat    string
}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func templateHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //needed to parse message to print out in console

	x := r.Form["usermsg"]

	//r.Form["chatbox"] = x

	//fmt.Fprintf(w, "Hello, %s!", r.Form["usermsg"])
	fmt.Println("jbh", x)
	//fmt.fprintf(w, &x)

	m := Data{Message: "Eliza robot chat", Chat: "user:"}

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, &m)

}

func main() {

	//call handler function
	http.HandleFunc("/", templateHandler)

	//http.HandleFunc("/guess", templateHandler)
	http.ListenAndServe(":8080", nil)
}
