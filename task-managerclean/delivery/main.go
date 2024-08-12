package main

import (
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := InitializeRouter()
	router.Route()
}
