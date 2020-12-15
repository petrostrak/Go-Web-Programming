package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}

/*
A web app using a different template engine that doesn’t scrub the input, and displays
user input directly on a web page, will get an alert message, or potentially any other
malicious code that the attacker writes. As you probably realize, the Go template
engine protects you from such mistakes because even if you don’t scrub the input,
when the input is displayed on the screen it’ll be converted into escaped HTML.

Comment: <script>alert('Pwnd!');</script>
*/
