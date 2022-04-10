package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files section")

	content := "This need to go inside file - yash"

	file, err := os.Create("./myFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)

	readFile("./myFile.txt")
	file.Close()
}

func readFile(path string) {

	databyte, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text Data inside the file is: \n", string(databyte))
}

