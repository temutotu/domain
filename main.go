package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"hello/domain/server/session"
	"hello/interface/handler"
)

func main() {
	ctx := context.Background()
	session.Init(&ctx)

	handler.HanlerInit()
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Ignore()
		signal.Notify(sigChan, syscall.SIGINT)

		s := <-sigChan

		switch s {
		case syscall.SIGINT:
			fmt.Println("SIGINT is recieved")
			os.Exit(0)

		default:
			panic("Other signal is recieved")
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
