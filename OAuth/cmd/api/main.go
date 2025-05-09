// main.go
package main

import (
	"OAuth/internal/auth"
	"OAuth/internal/server"
)

func main() {
	// Initialize Authentication
	auth.NewAuth()

	// Start the server
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
