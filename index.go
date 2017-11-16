//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html
//

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

var chatter string

type Data struct {
	Message, Chat string
	//Chat    []string
}

var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

func ElizaResponse(input string) string {

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		//match string

		return "why dont you tell me more about your father?"

	}

	return responses[rand.Intn(len(responses))]

}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.
func templateHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //needed to parse message to print out in console
	//var z string
	x := r.Form["usermsg"]

	m := Data{Message: "Eliza robot chat", Chat: ""}

	t, _ := template.ParseFiles("guess.html")

	/*x, err := r.URL.Query()["usermsg"]
	// if not found execute the template and exit
	if !err || len(x) < 1 {
		log.Println("Empty Input!!")
		// execute the template with the message
		t.Execute(w, m)
		return
	} // if*/
	//z = "" + x[0]
	//r.Form["chatbox"] = x

	//fmt.Fprintf(w, "Hello, %s!", r.Form["usermsg"])

	if len(x) > 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println("User Input", x)
		//fmt.Println(x)
		//difficulty tryint to get array to be matched with elizas "father"
		fmt.Println(ElizaResponse(""))
		fmt.Println()
	}
	//fmt.fprintf(w, &x)

	t.Execute(w, &m)

	//http.HandleFunc("/", inputHandler)

}

func inputHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() //needed to parse message to print out in console
	var z string
	x := r.Form["usermsg"]
	z = "" + x[0]

	m := Data{Message: "Eliza robot chat", Chat: "" + z}

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, m)

	//r.Form["chatbox"] = x

	//fmt.Fprintf(w, "Hello, %s!", r.Form["usermsg"])
	fmt.Println("User Input", x)
	//fmt.fprintf(w, &x)

	//t.Execute(w, &m)

}

func main() {

	//call handler function
	http.HandleFunc("/", templateHandler)
	//http.HandleFunc("/guess", inputHandler)

	//http.HandleFunc("/guess", templateHandler)
	http.ListenAndServe(":8080", nil)
}
