package service

import (
	"cat-facts/internal/models"
	"fmt"
)

func PrintBody(body models.CatFactBody) {
	fmt.Println("============= API Response =============")
	fmt.Println("Body:")
	fmt.Printf("Fact: %s\n", body.Fact)
	fmt.Printf("Length: %d\n", body.Length)
	fmt.Println("========================================")
}
