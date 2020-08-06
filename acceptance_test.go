package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

func TestAcceptance(t *testing.T) {
	serverExit := &sync.WaitGroup{}
	serverExit.Add(1)
	server := startHTTPServer(serverExit)

	defer func() {
		if err := server.Shutdown(context.TODO()); err != nil {
			panic(err)
		}

		serverExit.Wait()
	}()

	response, err := http.Get("http://localhost:5000/")
	if err != nil {
		t.Fatal(err)
	}

	defer response.Body.Close()

	bodyAsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	got := string(bodyAsBytes)
	want := "Hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
