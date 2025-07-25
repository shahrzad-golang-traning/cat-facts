package models

type CatFactBody struct {
	Fact   string `json:"fact"`   // Fact determines the given fact about a cat
	Length int    `json:"length"` // Length determines the length of the body
}
