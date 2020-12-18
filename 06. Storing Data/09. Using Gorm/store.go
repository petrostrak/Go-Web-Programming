package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var Db gorm.DB

func int() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AuthoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Petros Trak"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "Petros Trak").First(&readPost)
	var comments []Comment

	Db.Model(&readPost).Related(&comments)
	fmt.Println(comment[0])
}
