package main

import (
	"go-video-service/cmd"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := cmd.InitServer()
	go func() {
		err := http.ListenAndServe(":8001", handler)
		if err != nil {
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("Shutting down server")
}
