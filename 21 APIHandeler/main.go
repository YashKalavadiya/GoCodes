package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	performGetRequest()
	performPostJsonRequest()
	performPostFormRequest()

}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data

	data := url.Values{}
	data.Add("fname", "Yash")
	data.Add("lname", "Kalavadiya")
	data.Add("age", "20")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))


}

func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with the golang",
			"price": 0,
			"platform": "lco"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}

func performGetRequest(){
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("status code: ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)
	// fmt.Println("response: ", res.Body)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count: ", byteCount)

	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}