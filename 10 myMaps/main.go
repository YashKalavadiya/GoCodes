package main

import "fmt"

func main() {

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println("JS: ",languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)


	//loops are interesting in golang

	for key, value := range languages {
		fmt.Println(key, " : ", value)
	}
}

