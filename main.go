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
	cmd.InitServer()
	go func() {
		err := http.ListenAndServe(":8001", nil)
		if err != nil {
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("Shutting down server")
}
