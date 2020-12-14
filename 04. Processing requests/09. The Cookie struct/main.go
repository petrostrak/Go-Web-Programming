package main

import (
	"net/http"
	"time"
)

// The Cookie struct in Go
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Exprires   time.Time
	RawExpires string
	MaxAge     int
	Secure     bool
	HttpOnly   bool
	Raw        string
	Unparsed   []string
}

/*
If the Expires field isn’t set, then the cookie is a session or temporary cookie. Session
cookies are removed from the browser when the browser is closed. Otherwise, the
cookie is a persistent cookie that’ll last as long as it isn’t expired or removed.

There are two ways of specifying the expiry time: the Expires field and the MaxAge
field. Expires tells us exactly when the cookie will expire, and MaxAge tells us how
long the cookie should last (in seconds), starting from the time it’s created in the
browser. This isn’t a design issue with Go, but rather results from the inconsistent
implementation differences of cookies in various browsers. Expires was deprecated in
favor of MaxAge in HTTP 1.1, but almost all browsers still support it. MaxAge isn’t supported
by Microsoft Internet Explorer 6, 7, and 8. The pragmatic solution is to use
only Expires or to use both in order to support all browsers
*/

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Promming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Golang rocks",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Set("Set-cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}

/*
Go provides a simpler shortcut to setting the cookie: using the SetCookie function
in the net/http library.
It doesn’t make too much of a difference, though you should take note that you need
to pass in the cookies by reference instead.
*/
