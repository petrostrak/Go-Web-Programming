/*
The encoding/gob package manages streams of gobs, which are binary data,
exchanged between an encoder and a decoder. It’s designed for serialization and
transporting data but it can also be used for persisting data. Encoders and decoders
wrap around writers and readers, which conveniently allows you to use them to write
to and read from files.
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

// The Post struct
type Post struct {
	Id      int
	Content string
	Author  string
}

// Store data
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// Load data
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Petros Trak"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}