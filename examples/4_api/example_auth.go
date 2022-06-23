package main

import (
	basicauth "github.com/stokito/go-http-server-basic-auth"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	credentials := map[string]string{"admin": "secret"}
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test", handleReq)
	corsHandler := &basicauth.CorsHandlerWrapper{
		Handler: serveMux,
	}
	authHandler := basicauth.NewAuthHandlerWrapper(
		corsHandler,
		credentials,
		"Admin Dashboard",
		[]string{"/robots.txt", "/favicon.ico", "/.well-known/"},
	)
	recoverHandler := &basicauth.RecoveryHandlerWrapper{
		Handler:  authHandler,
		ErrorLog: logger.Printf,
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: recoverHandler,
	}
	server.ListenAndServe()
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := r.Header.Get(basicauth.HeaderSubscriberId)
		body := "Hello " + username
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(body))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
