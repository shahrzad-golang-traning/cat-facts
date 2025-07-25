package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactBody struct {
	Fact   string `json:"fact"`   // Fact determines the given fact about a cat
	Length int    `json:"length"` // Length determines the length of the body
}

func main() {
	url := "https://catfact.ninja/fact"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error while creating request:", err.Error())
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "GoLang-Client/1.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while sending request:", err.Error())
		return
	}
	defer res.Body.Close()

	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err.Error())
		return
	}

	body := CatFactBody{}

	err = json.Unmarshal(byteBody, &body)
	if err != nil {
		fmt.Printf("Failed to parse the request, error: %+v", err)
	}

	fmt.Println("============= API Response =============")
	fmt.Println("Status:", res.Status)
	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Headers:", res.Header.Get("Content-Type"))
	fmt.Println("----------------------------------------")
	fmt.Println("Body:")
	fmt.Printf("Fact: %s\n", body.Fact)
	fmt.Printf("Length: %d\n", body.Length)
	fmt.Printf("raw: %s\n", byteBody)
	fmt.Println("========================================")
}
