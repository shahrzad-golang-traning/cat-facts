package service

import (
	"cat-facts/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Call(url string, body *models.CatFactBody) error {

	req := MakeRequest(url)
	byteBody := makeCall(req)

	err := json.Unmarshal(byteBody, &body)
	if err != nil {
		return err
	}

	return nil
}

func MakeRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error while creating request:", err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "GoLang-Client/1.0")

	return req
}

func makeCall(req *http.Request) []byte {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	byteBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return byteBody
}
