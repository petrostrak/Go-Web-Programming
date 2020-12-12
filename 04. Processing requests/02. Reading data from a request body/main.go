package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

/*
Notice that you first need to determine how much to read; then you create a byte
array of the content length, and call the Read method to read into the byte array.

If you want to test this, you’ll need to send a POST request to the server with the
appropriate message body, because GET requests don’t have message bodies. Remember
that you can’t normally send POST requests through a browser—you need an
HTTP client. There are plenty of choices. You can use a desktop graphical HTTP client,
a browser plug-in or extension, or even cURL from the command line.
Type this on your console:

$ curl -id "first_name=petros&last_name=trakadas" 127.0.0.1:8080/body

Normally you wouldn’t need to read the raw form of the body, though, because Go provides
methods such as FormValue and FormFile to extract the values from a POST form.
*/
