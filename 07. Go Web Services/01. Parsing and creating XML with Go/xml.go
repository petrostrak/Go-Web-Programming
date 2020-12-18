package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Defines structs to represent the data
// `xml:"post"` called struct tag and Go determines the mapping
// between the struct and the XML elements using them.
// Struct tags are strings after each field that are a key-value pair.
// The key is a string and must not have a space, a quote ("), or a colon (:).
// The value must be a string between double quotes ("").
// For XML , the key must always be xml .
type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  string   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML Data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}

/*
Here are some rules for the XML struct tags:
To store the name of the XML element itself (normally the name of the struct is
the name of the element), add a field named XMLName with the type xml.Name .
The name of the element will be stored in that field.

To store the attribute of an XML element, define a field with the same name as
that attribute and use the struct tag `xml:"<name>,attr"` , where <name> is the
name of the XML attribute.

To store the character data value of the XML element, define a field with the
same name as the XML element tag, and use the struct tag `xml:",chardata"`.

To get the raw XML from the XML element, define a field (using any name) and
use the struct tag `xml:",innerxml"` .

If there are no mode flags (like ,attr or ,chardata or ,innerxml ) the struct
field will be matched with an XML element that has the same name as the
struct’s name.

To get to an XML element directly without specifying the tree structure to get to
that element, use the struct tag `xml:"a>b>c"` , where a and b are the interme-
diate elements and c is the node that you want to get to.
*/
