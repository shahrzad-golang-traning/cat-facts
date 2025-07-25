package service

import (
	"cat-facts/internal/models"
	"sync"
)

func InitCalls() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go DoCall(&wg)
	}
	wg.Wait()
}

func DoCall(wg *sync.WaitGroup) {
	defer wg.Done()

	url := "https://catfact.ninja/fact"
	body := models.CatFactBody{}

	Call(url, &body)
	PrintBody(body)
}
