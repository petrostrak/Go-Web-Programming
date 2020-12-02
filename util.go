package main

import (
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (s data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		s = data.Session{Uuid: cookie.Value}
		if ok, _ := s.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
