/*Tthe include actions allow us to include a template in another
template. As you might realize, these actions let us nest templates. The format of the
include action is {{ template "name" }}, where name is the name of the template to
be included.
*/

package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
