package server

import (
	"context"
	"go-todoapp/internal/db"
	"go-todoapp/internal/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	d := db.NewMemoryDB()
	s := http.NewServer(8080, d)
	errCh := make(chan error, 1)
	
	go func() {
		errCh <- s.Start()
	}()

	select {
	case <- termCh:
		return 0
	case <- errCh:
		return 1
	}
}
