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
	fmt.Fprintln(w, r.Form["hello"])
	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(3)", r.PostForm)
	fmt.Fprintln(w, "(4)", r.MultipartForm)
	//Of course, if you only want to get the value to the key post, you can use
	//r.Form["post"], which will give you a map with one element: [456]

	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

/*
What if you need just the form key-value pairs and want to totally ignore the URL
key-value pairs? For that you can use the PostForm field, which provides key-value
pairs only for the form and not the URL. If you change from using r.Form to using
r.PostForm in the code, this is what you get:

map[post:[456] hello:[petros trak]]
*/
