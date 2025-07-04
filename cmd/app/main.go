package main

import (
	"log"

	"github.com/khaleelsyed/beLLMan/internal/api"
	"github.com/khaleelsyed/beLLMan/internal/storage"
)

func main() {
	storage, err := storage.NewMockStorage()
	if err != nil {
		log.Fatal(err)
	}

	if err = storage.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", storage)

	server.Run()
}
