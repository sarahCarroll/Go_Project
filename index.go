// Author     - Sarah Carroll
// Student ID - G00330821

// References
//-----------
// https://golang.org/doc/articles/wiki/
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html
//https://stackoverflow.com/questions/16841320/an-html-tag-other-than-a-textarea-where-n-is-correctly-interpreted
//https://www.w3schools.com/tags/tryit.asp?filename=tryhtml5_input_type_hidden
//

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
)

var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}

var chatter, name string
var firstime int = 1

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

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with this handler.

func templateHandler(w http.ResponseWriter, r *http.Request) {

	var z, flagit, resp string

	r.ParseForm() //needed to parse message to print out in console
	x := r.Form["usermsg"]
	flag := r.Form["flag"]

	if len(x) > 0 {
		fmt.Println("User Input: ", x)
		fmt.Println("Flag[0]   : ", flag[0])
		fmt.Println("Name      : ", name)
		fmt.Println("Chat      : ", chatter)

		// build logic here to interpret question and set z appropriately

		// flag = 1 on initial query
		if flag[0] == "1" {
			name = x[0]
			z = "Eliza: Hello " + name + "!\nWhat is your query?\n\n"
			chatter += z // save conversation to date
			flagit = "2"
		} else {
			//	flag = 2 on subsequent queries
			resp = ElizaResponse(x[0])
			z = name + ": " + x[0] + "?\nEliza: " + resp + "\n\n"

			// z + = ElizaResponse(x[0]) + "\n\n"

			chatter += z // save conversation to date
			flagit = "2" // ensure flag remains at 2
		}
	} else {
		if firstime == 1 {
			z = "Eliza: Welcome, what is your name?\n\n"
			chatter = z // save conversation to date
			flagit = "1"
			firstime = 0
		}
	}

	//fmt.Println("Chat      : ", chatter)

	m := Data{Message: "Eliza Robot Chat", Chat: "" + chatter, Flag: "" + flagit}

	t, _ := template.ParseFiles("guess.html")

	t.Execute(w, &m)

}

func main() {

	//call handler function
	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8080", nil)
}
