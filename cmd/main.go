package main

import (
	"chatbot/cmd/ioc"
	"log"
)

func main() {
	app := ioc.InitApp()
	if err := app.Run(); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
