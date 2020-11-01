package main

import (
	"github.com/iluxaorlov/jetcracks/internal/app/jetcracks"
	"log"
	"os"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if err := jetcracks.Start(homeDir); err != nil {
		log.Fatal(err)
	}
}
