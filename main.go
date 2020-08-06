package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func startHTTPServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{Addr: ":5000"}

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Hello world")
		},
	)

	go func() {
		defer wg.Done()

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return server
}

func main() {
	serverExit := &sync.WaitGroup{}
	serverExit.Add(1)
	startHTTPServer(serverExit)
	serverExit.Wait()
}
