/*
In-memory storage refers not to storing data in in-memory databases but in the running
application itself, to be used while the application is running. In-memory data is
usually stored in data structures, and for Go, this primarily means with arrays, slices,
maps, and most importantly, structs
*/

package main

import "fmt"

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(p Post) {
	PostById[p.Id] = &p
	PostsByAuthor[p.Author] = append(PostsByAuthor[p.Author], &p)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World", Author: "Petros"}
	post2 := Post{Id: 2, Content: "こんにちは、世界！", Author: "エレニ"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Giannis"}
	post4 := Post{Id: 4, Content: "Bonjour Monde!", Author: "Greg"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Petros"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Greg"] {
		fmt.Println(post)
	}
}
