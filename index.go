// Author     - Sarah Carroll
// Student ID - G00330821

// References
//-----------
// https://golang.org/doc/articles/wiki/
// https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html
//https://stackoverflow.com/questions/16841320/an-html-tag-other-than-a-textarea-where-n-is-correctly-interpreted
//https://www.w3schools.com/tags/tryit.asp?filename=tryhtml5_input_type_hidden
//https://golang.org/pkg/regexp/syntax/
//https://stackoverflow.com/questions/7303948/how-to-auto-scroll-to-end-of-div-when-data-is-added
//https://www.icewarp.com/support/online_help/203030104.htm

package main

//To use the net/http package, it must be imported
import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
)

func Reflect(input string) string {
	// Split the input on word boundaries.
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)

	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
	}

	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}

	// Put the tokens back together.
	return strings.Join(tokens, ``)
}

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

	if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched {
		//match string

		return "Whats your mothers maiden name?"

	}

	if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, input); matched {
		//match string

		return "Does your brother annoy you too?"

	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}

	q := regexp.MustCompile(`(?i).*\bhi|hello\b.*([^.?!]*)[.?!]?`)
	if matched := q.MatchString(input); matched {
		return q.ReplaceAllString(input, "hello "+name)
	}

	qs := regexp.MustCompile(`(?i).*\bbye|goodbye\b.*([^.?!]*)[.?!]?`)
	if matched := qs.MatchString(input); matched {
		return qs.ReplaceAllString(input, "GoodBye "+name+" have a good day")
	}

	// List the reflections.

	// Put the tokens back together.

	//return the response
	return responses[rand.Intn(len(responses))]

}

//The main function begins with a call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with this handler.

func templateHandler(w http.ResponseWriter, r *http.Request) {

	var z, flagit, resp, ans string

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
			if len(x[0]) > 0 {
				ans = Reflect(x[0])
				fmt.Println(ans)
				resp = ElizaResponse(ans)
				//fmt.Println(Reflect("You are my friend."))
				z = name + ": " + x[0] + "\nEliza: " + resp + "\n\n"
			}

			// z + = ElizaResponse(x[0]) + "\n\n"

			chatter += z // save conversation to date
			flagit = "2" // ensure flag remains at 2
		}
	} else {
		//set global flag to 0 at to so only goes through the flag once
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

	fs := http.FileServer(http.Dir(""))
	http.Handle("/user-input", fs)

	http.ListenAndServe(":8080", nil)
}
