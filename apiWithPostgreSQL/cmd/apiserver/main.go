package main

import (
	"log"
	// "fmt"
	"github.com/LevOrlov5404/goFirstRestApi/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nill {
		log.Fatal(err)
	}
}