package main

import (
	"log"

	"github.com/shikidy/golang-rest/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
