package main

import (
	"yofio/intermedio/app"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app.StartApp()
}
