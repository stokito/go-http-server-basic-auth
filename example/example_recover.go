package main

import (
	basicauth "github.com/stokito/go-http-server-basic-auth"
	"log"
	"net/http"
	"os"
)

var emptiness *string

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test", buggyHandler())
	logger := log.New(os.Stdout, "", 0)
	server := &http.Server{
		Addr: ":8080",
		Handler: &basicauth.RecoveryHandlerWrapper{
			Handler:  serveMux,
			ErrorLog: logger,
		},
	}
	server.ListenAndServe()
}

func buggyHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		println("%s", *emptiness)
	}
}
