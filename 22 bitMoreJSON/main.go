package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}
func main() {
	fmt.Println("JSON example")

	// EncodeJson()

	DecodeJson()

}


func EncodeJson() {
	lcoCourses := []course{
		{"React Js Bootcamp", 299, "lco.in", "abcd123", []string{"react", "js"}},
		{"Angular Bootcamp", 299, "lco.in", "ijkl123", []string{"angular", "js", "ts"}},
		{"MERN Bootcamp", 199, "lco.in", "efgh123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React Js Bootcamp",
		"Price": 299,
		"website": "lco.in",
		"tags": ["react","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Data is Valid")

		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Data is invalid")
	}

	//cases were you just want to add data to key value pairs

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v := range myOnlineData {
			fmt.Printf("key is %v and value is %v and type: %T\n", k, v, v)
	}
}