package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cat struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	response, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Not able to connect", err)
		panic(err)
	}
	resBody, _ := io.ReadAll(response.Body)
	var cat Cat
	// response.Body.Close()
	err = json.Unmarshal(resBody, &cat)
	if err != nil {
		fmt.Println("unable to parse JSON", err)
		panic(err)
	}
	fmt.Printf("%s", cat.Fact)

}
