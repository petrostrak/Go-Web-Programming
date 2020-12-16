/*
The CSV format is a file format in which tabular data (numbers and text) can be easily
written and read in a text editor. In Go, CSV is manipulated by the encoding/csv package.
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// Creating a csv file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Petros Trak"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Eleni Apost"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Giannis Liol"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Greg Bail"},
	}

	writer := csv.NewWriter(csvFile)
	for _, p := range allPosts {
		line := []string{strconv.Itoa(p.Id), p.Content, p.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
