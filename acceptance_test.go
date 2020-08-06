package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

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
	assertNotError(t, err)

	defer response.Body.Close()
	bodyAsBytes, err := ioutil.ReadAll(response.Body)
	assertNotError(t, err)

	got := string(bodyAsBytes)
	want := "Hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
