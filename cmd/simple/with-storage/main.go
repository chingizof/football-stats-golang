package main

import (
	"context"
	"log"

	"github.com/chingizof/football-stats-golang/internal/http"
	"github.com/chingizof/football-stats-golang/internal/store/inmemory"
)

func main() {
	store := inmemory.NewDB()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}
