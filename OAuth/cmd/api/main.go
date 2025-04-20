package main

import (
	// "context"
	// "fmt"
	// "log"
	// "net/http"
	// "os/signal"
	// "syscall"
	// "time"

	"OAuth/internal/server"
	"OAuth/internal/auth"
)

func main() {

	//Instantiate Gothic
	auth.NewAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}

