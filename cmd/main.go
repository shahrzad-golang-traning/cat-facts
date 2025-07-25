package main

import (
	"fmt"
	"io"
	"net/http"
)

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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err.Error())
		return
	}

	fmt.Println("============= API Response =============")
	fmt.Println("Status:", res.Status)
	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Headers:", res.Header.Get("Content-Type"))
	fmt.Println("----------------------------------------")
	fmt.Println("Body:")
	fmt.Println(string(body))
	fmt.Println("========================================")
}
