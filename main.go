package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CatFact struct {
	Fact string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error while fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var fact CatFact
	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		fmt.Println("Error while decoding response:", err)
		return
	}

	fmt.Println("🐱 Cat Fact of the Day:", fact.Fact)
}
