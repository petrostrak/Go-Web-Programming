/*
The format of the name-value pairs sent through a POST request is specified by the
content type of the HTML form. This is defined using the enctype attribute like

<form action="/process" method="post" enctype="application/x-www-form-urlencoded">
	<input type="text" name="first_name"/>
	<input type="text" name="last_name"/>
	<input type="submit"/>
</form>

HTML allows the method attribute to be either POST or GET, so this is
also a valid format:

<form action="/process" method="get">
	<input type="text" name="first_name"/>
	<input type="text" name="last_name"/>
	<input type="submit"/>
</form>

In this case, there’s no request body (GET requests have no request body), and all the
data is set in the URL as name-value pairs.
*/

package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form["hello"]) //Of course, if you only want to get the value to the key post, you can use
	//r.Form["post"], which will give you a map with one element: [456]
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
