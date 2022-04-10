package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=cnijdsfier2xcesd8fx"

func main() {
	fmt.Println("Welcome to handeling urls in golang")

	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Hostname())

	params := result.Query()
	fmt.Println(params["coursename"])
	fmt.Println(params["paymentid"])

	for idx, val := range params {
		fmt.Println("Param is: ", val," Key is:", idx)
	}

	partsOfUrl := url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=yash",
	}

	fmt.Println(partsOfUrl.String())
}