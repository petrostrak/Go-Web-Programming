package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	// Writes to file and reads from file using WriteFile and ReadFIle
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// Writes to file and reads from file using the File struct
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}

/*
To write a file, you first create it using the Create function in the os
package, passing it the name of the file you want to create. It’s good practice to use
defer to close the file so that you won’t forget.

Reading a file with the File struct is similar. You need to use the Open function in
the os package, and then use the Read method on the File struct, or any of the other
methods to read the data. Reading data using the File struct is much more flexible
because File has several other methods you can use to locate the correct part of the
file you want to read from.
*/
