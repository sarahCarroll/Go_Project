//Author- Sarah Carroll
// https://golang.org/doc/articles/wiki/
//https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html
//

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	//	"html"
	"html/template"
	"net/http"
)

var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

var chatter string

type Data struct {
	Message, Chat, Flag string
}

func ElizaResponse(input string) string {

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		//match string

		return "why dont you tell me more about your father?"

	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}

	return responses[rand.Intn(len(responses))]

}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler.

func templateHandler(w http.ResponseWriter, r *http.Request) {

	var z, name, flagit, resp string

	r.ParseForm() //needed to parse message to print out in console
	x := r.Form["usermsg"]
	flag := r.Form["flag"]

	if len(x) > 0 {
		fmt.Println("User Input", x)

		// build logic here to interpret question and set z appropriately
		// flag = 1 on initial query
		if flag[0] == "1" {
			name = x[0]
			z = chatter + "Hello " + name + "! What is your query?"
			flagit = "2"
		}

		if flag[0] == "2" {
			rand.Seed(time.Now().UTC().UnixNano())
			//fmt.Println("User Input:", x[0])
			//fmt.Println(x)
			//difficulty tryint to get array to be matched with elizas "father"

			resp = ElizaResponse(x[0])
			fmt.Println(resp)
			z = chatter + "Eliza:" + resp
			fmt.Println()
			flagit = "2"
		}
	} else {
		z = "Welcome, what is your name?      "
		flagit = "1"
	}

	chatter = z // save conversation to date

	m := Data{Message: "Eliza Robot Chat", Chat: "   " + chatter, Flag: "" + flagit}

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, &m)

}

func main() {

	//call handler function
	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8080", nil)
}
