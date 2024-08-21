package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
)

const sockfile = "/var/tmp/go-domain-sock/main.sock"

type Handler struct{}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	fmt.Fprintln(w, "<h1>Hello, from Golang!</h1>")
}

func main() {
	// about graceful shutdown
	interrupts := make(chan os.Signal, 1)
	signal.Notify(interrupts, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// handler
	handler := Handler{}

	// http server
	server := &http.Server{
		Handler: &handler,
	}

	// create temp directory
	err := os.Mkdir(path.Dir(sockfile), 0700)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		if err := os.RemoveAll(path.Dir(sockfile)); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	// listen
	go func() error {
		listen, err := net.Listen("unix", sockfile)
		if err != nil {
			return err
		}
		defer func() {
			if err := listen.Close(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()
		if err := os.Chmod(sockfile, 0700); err != nil {
			return err
		}
		fmt.Println("listen on \"" + sockfile + "\"")
		return server.Serve(listen)
	}()
	defer func() {
		if err := server.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	// shutdown
	<-interrupts
	fmt.Println(" Interrupt.")
}
